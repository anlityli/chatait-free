// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"sync"
	"time"
)

type QueueEvent struct {
	EventType string
	Config    *entity.ConfigMidjourney
	QueueData *entity.QueueMidjourney
	Message   *QueueEventMessage
}

type QueueEventMessage = WsReceiveMessageDCommon

type QueueTask struct {
	Data          *entity.QueueMidjourney  // 队列数据，与数据库一致
	StatusChannel chan int                 // 状态通道，用来监听任务执行进度情况
	Config        *entity.ConfigMidjourney // 任务对应的配置信息 防止重复查询所以单提出来
}

// QueueClient Queue客户端
type QueueClient struct {
	Event          chan *QueueEvent
	channel        chan *QueueTask
	concurrentSize int
	queueSize      int
	workingPool    sync.Map // 正在消费过程中的池
}

var queueClient *QueueClient
var queueClientOnce sync.Once

// QueueInstance 单例
func QueueInstance() *QueueClient {
	queueClientOnce.Do(func() {
		queueClient = &QueueClient{}
		if err := queueClient.init(); err != nil {
			glog.Line().Println("queue init失败:" + err.Error())
		}
	})
	return queueClient
}

func (q *QueueClient) init() (err error) {
	// 拿到配置中的队消费者数和队列数
	midjourneyProgressSizeStr, err := helper.GetConfig("midjourneyProgressSize")
	if err != nil {
		return err
	}
	q.concurrentSize = gconv.Int(midjourneyProgressSizeStr)
	midjourneyQueueSizeStr, err := helper.GetConfig("midjourneyQueueSize")
	if err != nil {
		return err
	}
	q.queueSize = gconv.Int(midjourneyQueueSizeStr)
	q.channel = make(chan *QueueTask, q.queueSize)
	q.Event = make(chan *QueueEvent)
	glog.Line(true).Debug("队列实例初始化完成")
	return nil
}

func (q *QueueClient) Run() {
	glog.Line(true).Debug("队列开始运行")
	// 建立同时执行任务的消费者
	for i := 0; i < q.concurrentSize; i++ {
		go func(workerId int) {
			glog.Line(true).Debug("消费者开始运行", workerId)
			// 执行消费任务
			for {
				task := <-q.channel
				glog.Line(true).Debug("拿到队列任务", task.Data.Id)
				if task.Data.Id > 0 {
					q.execTask(task)
				}
			}
		}(i + 1)
	}
	glog.Line(true).Debug("队列结束运行")
}

func (q *QueueClient) InsertTask(queueData *entity.QueueMidjourney) (err error) {
	glog.Line(true).Debug("队列生产者插入任务", queueData)
	configData := &entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("id=?", queueData.ConfigId).Scan(configData)
	if err != nil {
		return err
	}
	// 在这里执行ws的监听
	if configData.ListenModel == constant.ConfigMidjourneyListenModelUserWss {
		waitConnChan := make(chan int)
		err = WsRun(configData, func() {
			waitConnChan <- 1
		})
		if err != nil {
			glog.Line(true).Debug("Ws监听发生错误", err)
			close(waitConnChan)
			return err
		}
		afterConn := <-waitConnChan
		if afterConn > 0 {
			glog.Line(true).Debug("Ws连接完成，开始插入队列任务")
		}
	}
	task := &QueueTask{
		Data:          queueData,
		StatusChannel: make(chan int),
		Config:        configData,
	}
	select {
	case q.channel <- task:
		// 把任务写入数据库
		if _, err = dao.QueueMidjourney.Data(queueData).Insert(); err != nil {
			return err
		}
		return nil
	default:
		glog.Line(true).Debug("队列已经满了")
		return errors.New("排队任务已满，请稍候再试")
	}
}

func (q *QueueClient) execTask(task *QueueTask) {
	glog.Line(true).Debug("开始执行队列任务", task)
	// 把任务加入到正在执行的任务池中
	q.workingPool.Store(task.Data.Id, task)
	// 建立队列的状态通道
	task.StatusChannel = make(chan int)
	// 开一个定时器，超时后，如果任务还没有变化就中断
	q.setTimeoutTimer(task.Data.Id)
	go func() {
		// 把任务设置为正在执行
		task.Data.Status = constant.QueueMidjourneyStatusProceeding
		task.Data.StartedAt = gconv.Int(xtime.GetNowTime())
		q.changeTaskData(&changeTaskDataParams{
			eventType: constant.QueueMidjourneyEventInsertQueue,
			queueData: task.Data,
		})
		glog.Line(true).Debug("开始请求", task.Data.RequestData, task.Data.RequestUrl)
		_, err := q.request(task.Config, task.Data.RequestData, task.Data.RequestUrl)
		if err != nil {
			glog.Line(true).Debug(err)
			return
		}
		glog.Line(true).Debug("开启定时器", task.Data.RequestData, task.Data.RequestUrl)
	}()
	// 等待任务的反馈情况，直到任务完成或者超时退出则完成任务
	for {
		status := <-task.StatusChannel
		glog.Line(true).Debug("队列状态通道有状态", status)
		if status == constant.QueueMidjourneyStatusEnded {
			close(task.StatusChannel)
			break
		} else if status == constant.QueueMidjourneyStatusError {
			close(task.StatusChannel)
			break
		}
	}
	glog.Line(true).Debug("删除掉任务池中的任务")
	// 任务结束，把任务池中的任务删除掉
	q.workingPool.Delete(task.Data.Id)
}

// changeTaskDataParams 队列任务发生变化的参数
type changeTaskDataParams struct {
	eventType string
	queueData *entity.QueueMidjourney
	message   interface{}
}

// changeTaskData 队列任务发生变化
func (q *QueueClient) changeTaskData(params *changeTaskDataParams) {
	glog.Line(true).Debug("任务开始发生变化的事件", params)
	// 从池中找到对应的任务
	taskObj, ok := q.workingPool.Load(params.queueData.Id)
	if !ok {
		glog.Line(true).Println("队列任务"+gconv.String(params.queueData.Id)+"不存在", params)
		return
	}
	glog.Line(true).Debug("任务更新变化", params)
	task := taskObj.(*QueueTask)
	task.Data = params.queueData
	if _, err := dao.QueueMidjourney.Data(task.Data).Where("id=?", task.Data.Id).Update(); err != nil {
		glog.Line(true).Println("队列任务变化同步到数据库失败", params)
	}
	// 根据事件回调事件函数
	var message *QueueEventMessage
	if params.message != nil {
		messageMap := gconv.Map(params.message)
		message = &QueueEventMessage{}
		if err := gconv.Scan(messageMap, message); err != nil {
			glog.Line(true).Println("队列任务变化转换消息内容失败", params, messageMap)
		}
	}
	eventData := &QueueEvent{
		EventType: params.eventType,
		Config:    task.Config,
		QueueData: task.Data,
		Message:   message,
	}
	glog.Line(true).Debug("任务发生变化的事件", eventData)
	q.Event <- eventData
	task.StatusChannel <- params.queueData.Status
}

// SetTimeoutTimer 设定任务超时定时器，到时间任务还没有完成，则直接标记为任务超时结束
func (q *QueueClient) setTimeoutTimer(queueId int64) {
	queueTimeoutConfig, err := helper.GetConfig("midjourneyQueueTimeout")
	if err != nil {
		glog.Line(true).Println("获取队列配置错误", err)
		return
	}
	queueTimeout := gconv.Int(queueTimeoutConfig)
	gtimer.AddOnce(time.Duration(queueTimeout)*time.Second, func() {
		taskObj, ok := q.workingPool.Load(queueId)
		if ok {
			task := taskObj.(*QueueTask)
			if task.Data.Status <= constant.QueueMidjourneyStatusProceeding {
				task.Data.Status = constant.QueueMidjourneyStatusError
				task.Data.ErrorAt = gconv.Int(xtime.GetNowTime())
				task.Data.ErrorData = "任务超时结束"
				q.changeTaskData(&changeTaskDataParams{
					eventType: constant.QueueMidjourneyEventError,
					queueData: task.Data,
				})
			}
		}
	})
}

func (q *QueueClient) request(config *entity.ConfigMidjourney, requestData string, requestUrl string) (response []byte, err error) {
	httpClient := ghttp.NewClient()
	if config.Proxy != "" {
		glog.Line(true).Debug("使用代理访问", config.Proxy)
		httpClient.SetProxy(config.Proxy)
	}
	httpClient.Timeout(30)
	httpClient.SetHeader("Content-Type", "application/json")
	httpClient.SetHeader("Authorization", config.UserToken)
	httpClient.SetHeader("User-Agent", config.UserAgent)
	glog.Line(true).Debug(gconv.String(requestData))
	glog.Line(true).Debug(requestUrl)
	resp, err := httpClient.Post(requestUrl, requestData)
	if err != nil {
		glog.Line(true).Println(err.Error())
		return nil, err
	}
	defer resp.Close()
	if resp.StatusCode < 300 {
		return resp.ReadAll(), nil
	} else {
		return nil, errors.New("请求错误" + resp.Status)
	}
}
