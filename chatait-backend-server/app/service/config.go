// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Config = &configService{}

type configService struct {
}

// AllOption 全部选项
func (s *configService) AllOption(r *ghttp.Request) (re *response.ConfigOptionList, err error) {
	configList := &[]*entity.Config{}
	if err := dao.Config.Where("1=1").Order("type ASC, sort ASC").Scan(configList); err != nil {
		return nil, err
	}
	re = &response.ConfigOptionList{}
	for _, item := range *configList {
		tempItem := &response.ConfigOptionItem{
			ConfigName: item.ConfigName,
			Title:      item.Title,
			Unit:       item.Unit,
			InputType:  item.InputType,
			Options:    make([]*response.ConfigOptionItemOption, 0),
			Value:      item.Value,
			Type:       item.Type,
			Sort:       item.Sort,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		if item.Options != "" {
			err = gjson.DecodeTo(item.Options, &tempItem.Options)
			if err != nil {
				return nil, err
			}
		}
		*re = append(*re, tempItem)
	}
	return re, err
}

// OptionEdit 编辑选项
func (s *configService) OptionEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigOptionEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	configData, err := dao.Config.Where("config_name=?", requestModel.ConfigName).One()
	if err != nil {
		return err
	}
	if configData.IsEmpty() {
		return errors.New("选项不存在")
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		if _, err := dao.Config.Ctx(ctx).TX(tx).Data(g.Map{
			"value": requestModel.Value,
		}).Where("config_name=?", requestModel.ConfigName).Update(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *configService) LevelList(r *ghttp.Request) (re *response.ConfigLevelList, err error) {
	re = &response.ConfigLevelList{}
	err = dao.ConfigLevel.Where("1=1").Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) LevelEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigLevelEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if _, err := dao.ConfigLevel.Data(g.Map{
		requestModel.Field: requestModel.Value,
	}).Where("id", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

func (s *configService) WalletList(r *ghttp.Request) (re *response.ConfigWalletList, err error) {
	re = &response.ConfigWalletList{}
	err = dao.ConfigWallet.Where("1=1").Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) WalletEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigWalletEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if _, err := dao.ConfigWallet.Data(g.Map{
		"wallet_name": requestModel.WalletName,
	}).Where("field", requestModel.Field).Update(); err != nil {
		return err
	}
	return nil
}

func (s *configService) WalletOne(r *ghttp.Request) (re *response.ConfigWallet, err error) {
	requestModel := &request.ConfigWalletOne{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ConfigWallet{}
	err = dao.ConfigWallet.Where("field=?", requestModel.Field).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) PayList(r *ghttp.Request) (re *response.ConfigPayList, err error) {
	payListData := &[]*entity.ConfigPay{}
	err = dao.ConfigPay.Where("1=1").Scan(payListData)
	if err != nil {
		return nil, err
	}
	re = &response.ConfigPayList{}
	*re = make(response.ConfigPayList, 0)
	if len(*payListData) > 0 {
		for _, item := range *payListData {
			tempItem := &response.ConfigPay{
				Id:                  gconv.String(item.Id),
				ApiName:             item.ApiName,
				FrontendDescription: item.FrontendDescription,
				BackendDescription:  item.BackendDescription,
				Status:              item.Status,
				CreatedAt:           item.CreatedAt,
				UpdatedAt:           item.UpdatedAt,
			}
			if item.Params != "" {
				paramsJson, err := gjson.Decode(item.Params)
				if err != nil {
					return nil, err
				}
				err = gconv.Scan(paramsJson, &tempItem.Params)
				if err != nil {
					return nil, err
				}
			}
			if item.PayChannel != "" {
				payChannelJson, err := gjson.Decode(item.PayChannel)
				if err != nil {
					return nil, err
				}
				err = gconv.Scan(payChannelJson, &tempItem.PayChannel)
				if err != nil {
					return nil, err
				}
			}
			*re = append(*re, tempItem)
		}
	}
	return re, nil
}

func (s *configService) PayOne(r *ghttp.Request) (re *response.ConfigPay, err error) {
	requestModel := &request.ConfigId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	payData := &entity.ConfigPay{}
	err = dao.ConfigPay.Where("id=?", requestModel.Id).Scan(payData)
	if err != nil {
		return nil, err
	}
	re = &response.ConfigPay{
		Id:                  gconv.String(payData.Id),
		ApiName:             payData.ApiName,
		FrontendDescription: payData.FrontendDescription,
		BackendDescription:  payData.BackendDescription,
		Status:              payData.Status,
		CreatedAt:           payData.CreatedAt,
		UpdatedAt:           payData.UpdatedAt,
	}
	if payData.Params != "" {
		paramsJson, err := gjson.Decode(payData.Params)
		if err != nil {
			return nil, err
		}
		err = gconv.Scan(paramsJson, &re.Params)
		if err != nil {
			return nil, err
		}
	}
	if payData.PayChannel != "" {
		payChannelJson, err := gjson.Decode(payData.PayChannel)
		if err != nil {
			return nil, err
		}
		err = gconv.Scan(payChannelJson, &re.PayChannel)
		if err != nil {
			return nil, err
		}
	}
	return re, nil
}

func (s *configService) PayEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigPayEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	updatedData := gconv.Map(requestModel)
	delete(updatedData, "id")
	params, err := gjson.Encode(requestModel.Params)
	if err != nil {
		return err
	}
	updatedData["params"] = gconv.String(params)
	payChannel, err := gjson.Encode(requestModel.PayChannel)
	if err != nil {
		return err
	}
	updatedData["pay_channel"] = gconv.String(payChannel)
	if _, err = dao.ConfigPay.Data(updatedData).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

// MidjourneyList Midjourney配置列表
func (s *configService) MidjourneyList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Config{}
	listColumns := columnsModel.MidjourneyListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConfigMidjourneyList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ConfigMidjourney.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		OrderBy:     "id ASC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *configService) MidjourneyOne(r *ghttp.Request) (re *response.ConfigMidjourney, err error) {
	requestModel := &request.ConfigId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("id=?", g.Slice{requestModel.Id}).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) MidjourneyAdd(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigMidjourneyAdd{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigMidjourney.Where("title=?", requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	insertID := snowflake.GenerateID()
	nowTime := xtime.GetNowTime()
	insertData := gconv.Map(requestModel)
	insertData["id"] = insertID
	insertData["created_at"] = nowTime
	if _, err := dao.ConfigMidjourney.Data(insertData).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *configService) MidjourneyEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigMidjourneyEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigMidjourney.Where("id<>? AND title=?", requestModel.Id, requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	nowTime := xtime.GetNowTime()
	updateData := gconv.Map(requestModel)
	updateData["updated_at"] = nowTime
	if _, err := dao.ConfigMidjourney.Data(updateData).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

func (s *configService) MidjourneyDelete(r *ghttp.Request) error {
	requestModel := &request.ConfigIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		for _, id := range requestModel.Selected {
			data := &entity.ConfigMidjourney{}
			err = dao.ConfigMidjourney.Ctx(ctx).TX(tx).Where("id=?", id).Scan(data)
			if err != nil {
				return err
			}
			if data.Status == 1 {
				return errors.New("修改【" + data.Title + "】的状态为未启用后再删除")
			}
			if _, err := dao.ConfigMidjourney.Ctx(ctx).TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// OpenaiList Openai配置列表
func (s *configService) OpenaiList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Config{}
	listColumns := columnsModel.OpenaiListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConfigOpenaiList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ConfigOpenai.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		OrderBy:     "id ASC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *configService) OpenaiOne(r *ghttp.Request) (re *response.ConfigOpenai, err error) {
	requestModel := &request.ConfigId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ConfigOpenai{}
	err = dao.ConfigOpenai.Where("id=?", g.Slice{requestModel.Id}).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) OpenaiAdd(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigOpenaiAdd{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigOpenai.Where("title=?", requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	insertID := snowflake.GenerateID()
	nowTime := xtime.GetNowTime()
	insertData := gconv.Map(requestModel)
	insertData["id"] = insertID
	insertData["created_at"] = nowTime
	if _, err := dao.ConfigOpenai.Data(insertData).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *configService) OpenaiEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigOpenaiEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigOpenai.Where("id<>? AND title=?", requestModel.Id, requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	nowTime := xtime.GetNowTime()
	updateData := gconv.Map(requestModel)
	updateData["updated_at"] = nowTime
	if _, err := dao.ConfigOpenai.Data(updateData).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

func (s *configService) OpenaiDelete(r *ghttp.Request) error {
	requestModel := &request.ConfigIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		for _, id := range requestModel.Selected {
			data := &entity.ConfigOpenai{}
			err = dao.ConfigOpenai.Ctx(ctx).TX(tx).Where("id=?", id).Scan(data)
			if err != nil {
				return err
			}
			if data.Status == 1 {
				return errors.New("修改【" + data.Title + "】的状态为未启用后再删除")
			}
			if _, err := dao.ConfigOpenai.Ctx(ctx).TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// BaiduList Baidu配置列表
func (s *configService) BaiduList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Config{}
	listColumns := columnsModel.BaiduListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConfigBaiduList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ConfigBaidu.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		OrderBy:     "id ASC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *configService) BaiduOne(r *ghttp.Request) (re *response.ConfigBaidu, err error) {
	requestModel := &request.ConfigId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ConfigBaidu{}
	err = dao.ConfigBaidu.Where("id=?", g.Slice{requestModel.Id}).Scan(re)
	if err != nil {
		return nil, err
	}
	if len(re.Features) > 0 {
		featuresJson, err := gjson.Decode(re.Features[0])
		if err != nil {
			return nil, err
		}
		re.Features = gconv.SliceStr(featuresJson)
	}
	return re, nil
}

func (s *configService) BaiduAdd(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigBaiduAdd{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigBaidu.Where("title=?", requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	insertID := snowflake.GenerateID()
	nowTime := xtime.GetNowTime()
	insertData := gconv.Map(requestModel)
	insertData["id"] = insertID
	insertData["created_at"] = nowTime
	if _, err := dao.ConfigBaidu.Data(insertData).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *configService) BaiduEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigBaiduEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigBaidu.Where("id<>? AND title=?", requestModel.Id, requestModel.Title).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置标题重复")
	}
	nowTime := xtime.GetNowTime()
	updateData := gconv.Map(requestModel)
	updateData["updated_at"] = nowTime
	if _, err := dao.ConfigBaidu.Data(updateData).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

func (s *configService) BaiduDelete(r *ghttp.Request) error {
	requestModel := &request.ConfigIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		for _, id := range requestModel.Selected {
			data := &entity.ConfigBaidu{}
			err = dao.ConfigBaidu.Ctx(ctx).TX(tx).Where("id=?", id).Scan(data)
			if err != nil {
				return err
			}
			if data.Status == 1 {
				return errors.New("修改【" + data.Title + "】的状态为未启用后再删除")
			}
			if _, err := dao.ConfigBaidu.Ctx(ctx).TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *configService) SensitiveWordList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Config{}
	listColumns := columnsModel.SensitiveWordListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConfigSensitiveWordList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ConfigSensitiveWord.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		OrderBy:     "id Desc",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *configService) SensitiveWordAdd(r *ghttp.Request) (err error) {
	requestModel := &request.ConfigSensitiveWordAdd{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.ConfigSensitiveWord.Where("content=?", requestModel.Content).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("配置内容重复")
	}
	insertID := snowflake.GenerateID()
	nowTime := xtime.GetNowTime()
	insertData := gconv.Map(requestModel)
	insertData["id"] = insertID
	insertData["created_at"] = nowTime
	if _, err := dao.ConfigSensitiveWord.Data(insertData).Insert(); err != nil {
		return err
	}
	return nil
}

func (s *configService) SensitiveWordDelete(r *ghttp.Request) error {
	requestModel := &request.ConfigIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		for _, id := range requestModel.Selected {
			if _, err := dao.ConfigSensitiveWord.Ctx(ctx).TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
