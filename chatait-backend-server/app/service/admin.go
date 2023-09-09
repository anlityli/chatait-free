// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

var Admin = &adminService{}

type adminService struct {
}

// List 管理员列表
func (s *adminService) List(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Admin{}
	listColumns := columnsModel.ListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	adminListModel := &response.AdminListItemList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.Admin.Table + " a",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.AdminRole.Table + " ar",
				On:        "a.role_id=ar.id",
			},
		},
		Field:   "a.*, ar.role_name",
		OrderBy: "a.id ASC",
	}, adminListModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

// One 一条数据
func (s *adminService) One(r *ghttp.Request) (re *response.AdminListItem, err error) {
	requestModel := &request.AdminId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	data, err := dao.Admin.As("a").LeftJoin(dao.User.Table+" u", "u.id=a.user_id").Where("a.id=?", g.Slice{requestModel.Id}).Fields("a.*,u.username").One()
	if err != nil {
		return nil, err
	}
	re = &response.AdminListItem{}
	if err := gconv.Struct(data, re); err != nil {
		return nil, err
	}
	return re, nil
}

// AllRole 所有角色
func (s *adminService) AllRole(r *ghttp.Request) (re *response.AdminRoleItemList, err error) {
	re = &response.AdminRoleItemList{}
	data, err := dao.AdminRole.Where("1=1").All()
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return re, err
	}
	if err := gconv.SliceStruct(data, re); err != nil {
		return nil, err
	}
	return re, nil
}

func (s *adminService) RoleList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Admin{}
	listColumns := columnsModel.RoleListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.AdminRoleItemList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.AdminRole.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *adminService) RoleOne(r *ghttp.Request) (re *response.AdminRoleItem, err error) {
	requestModel := &request.AdminId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.AdminRoleItem{}
	err = dao.AdminRole.Where("id=?", g.Slice{requestModel.Id}).Scan(re)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return re, nil
}

// RoleAdd 添加角色
func (s *adminService) RoleAdd(r *ghttp.Request) (re *response.AdminRoleItem, err error) {
	requestModel := &request.AdminRoleAdd{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	data, err := dao.AdminRole.Where("role_name=?", g.Slice{requestModel.RoleName}).One()
	if err != nil {
		return nil, err
	}
	if !data.IsEmpty() {
		return nil, errors.New("角色名重复")
	}
	// 添加角色
	id := snowflake.GenerateID()
	if _, err := dao.AdminRole.Data(g.Map{
		"id":           id,
		"role_name":    requestModel.RoleName,
		"remark":       requestModel.Remark,
		"create_admin": auth.GetAdminName(r),
		"created_at":   xtime.GetNowTime(),
	}).Insert(); err != nil {
		return nil, err
	}

	re = &response.AdminRoleItem{}
	re.ID = gconv.String(id)
	re.RoleName = requestModel.RoleName
	re.Remark = requestModel.Remark
	return re, nil
}

func (s *adminService) RoleEdit(r *ghttp.Request) (err error) {
	requestModel := &request.AdminRoleEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	data, err := dao.AdminRole.Where("role_name=? AND id<>?", g.Slice{requestModel.RoleName, requestModel.Id}).One()
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("角色名重复")
	}
	if _, err := dao.AdminRole.Data(g.Map{
		"role_name":     requestModel.RoleName,
		"remark":        requestModel.Remark,
		"updated_admin": auth.GetAdminName(r),
		"updated_at":    xtime.GetNowTime(),
	}).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}

	return nil
}

func (s *adminService) RoleDelete(r *ghttp.Request) error {
	requestModel := &request.AdminIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range requestModel.Selected {
			if _, err := dao.AdminRole.TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// ResetPassword 重置密码
func (s *adminService) ResetPassword(r *ghttp.Request) (err error) {
	reqModel := &request.AdminResetPassword{}
	if err := r.Parse(reqModel); err != nil {
		return err
	}
	adminID := auth.GetAdminID(r)
	if err := s.ChangeAdminPassword(adminID, reqModel.Password); err != nil {
		return err
	}
	return nil
}

// ResetOtherPassword 重置会员密码
func (s *adminService) ResetOtherPassword(r *ghttp.Request) (err error) {
	reqModel := &request.AdminResetOtherPassword{}
	if err := r.Parse(reqModel); err != nil {
		return err
	}
	if err := s.ChangeAdminPassword(gconv.Uint64(reqModel.AdminId), reqModel.Password); err != nil {
		return err
	}
	return nil
}

// ChangeAdminPassword 重置管理员密码
func (s *adminService) ChangeAdminPassword(adminID uint64, password string) (err error) {
	passwordHash, err := security.GeneratePassword(password)
	if err != nil {
		return err
	}
	if _, err := dao.Admin.Where("id=?", adminID).Data(g.Map{
		"password_hash": passwordHash,
	}).Update(); err != nil {
		return err
	}
	return nil
}

// Add 添加管理员
func (s *adminService) Add(r *ghttp.Request) (err error) {
	reqModel := &request.AdminAdd{}
	if err := r.Parse(reqModel); err != nil {
		return err
	}
	if reqModel.Password == "" {
		return errors.New("密码不能为空")
	}
	adminData, err := dao.Admin.Where("admin_name=?", reqModel.AdminName).One()
	if err != nil {
		return err
	}
	if !adminData.IsEmpty() {
		return errors.New("管理员名已经存在")
	}
	if reqModel.UserID != "" && reqModel.UserID != "0" {
		userData, err := dao.User.Where("id=?", reqModel.UserID).One()
		if err != nil {
			return err
		}
		if userData.IsEmpty() {
			return errors.New("前台会员不存在")
		}
	}
	if reqModel.RoleID != "" {
		roleData, err := dao.AdminRole.Where("id=?", reqModel.RoleID).One()
		if err != nil {
			return err
		}
		if roleData.IsEmpty() {
			return errors.New("角色不存在")
		}
	}
	bindIP := ""
	if len(reqModel.BindIP) > 0 {
		for _, item := range reqModel.BindIP {
			if err := gvalid.CheckValue(context.TODO(), item, "ip", "绑定的IP格式不正确"); err != nil {
				return err
			}
		}
		tempBindIP, err := gjson.Encode(reqModel.BindIP)
		if err != nil {
			return err
		}
		bindIP = gconv.String(tempBindIP)
	}
	insertID := snowflake.GenerateID()
	passwordHash, err := security.GeneratePassword(reqModel.Password)
	nowTime := xtime.GetNowTime()
	if _, err := dao.Admin.Data(g.Map{
		"id":            insertID,
		"user_id":       gconv.Int64(reqModel.UserID),
		"admin_name":    reqModel.AdminName,
		"real_name":     reqModel.RealName,
		"remark":        reqModel.Remark,
		"role_id":       reqModel.RoleID,
		"is_enable":     reqModel.IsEnable,
		"password_hash": passwordHash,
		"bind_ip":       bindIP,
		"create_admin":  auth.GetAdminName(r),
		"created_at":    nowTime,
	}).Insert(); err != nil {
		return err
	}
	return nil
}

// Edit 添加管理员
func (s *adminService) Edit(r *ghttp.Request) (err error) {
	reqModel := &request.AdminEdit{}
	if err := r.Parse(reqModel); err != nil {
		return err
	}
	adminData, err := dao.Admin.Where("admin_name=? AND id<>?", reqModel.AdminName, reqModel.ID).One()
	if err != nil {
		return err
	}
	if !adminData.IsEmpty() {
		return errors.New("管理员名已经存在")
	}
	if reqModel.UserID != "" && reqModel.UserID != "0" {
		userData, err := dao.User.Where("id=?", reqModel.UserID).One()
		if err != nil {
			return err
		}
		if userData.IsEmpty() {
			return errors.New("前台会员不存在")
		}
	}
	if reqModel.RoleID != "" {
		roleData, err := dao.AdminRole.Where("id=?", reqModel.RoleID).One()
		if err != nil {
			return err
		}
		if roleData.IsEmpty() {
			return errors.New("角色不存在")
		}
	}
	bindIP := ""
	if len(reqModel.BindIP) > 0 {
		for _, item := range reqModel.BindIP {
			if err := gvalid.CheckValue(context.TODO(), item, "ip", "绑定的IP格式不正确"); err != nil {
				return err
			}
		}
		tempBindIP, err := gjson.Encode(reqModel.BindIP)
		if err != nil {
			return err
		}
		bindIP = gconv.String(tempBindIP)
	}
	nowTime := xtime.GetNowTime()
	updateData := g.Map{
		"user_id":      gconv.Int64(reqModel.UserID),
		"admin_name":   reqModel.AdminName,
		"real_name":    reqModel.RealName,
		"remark":       reqModel.Remark,
		"role_id":      reqModel.RoleID,
		"is_enable":    reqModel.IsEnable,
		"bind_ip":      bindIP,
		"update_admin": auth.GetAdminName(r),
		"updated_at":   nowTime,
	}
	if reqModel.Password != "" {
		passwordHash, err := security.GeneratePassword(reqModel.Password)
		if err != nil {
			return err
		}
		updateData["password_hash"] = passwordHash
	}
	if _, err := dao.Admin.Data(updateData).Where("id=?", reqModel.ID).Update(); err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (s *adminService) Delete(r *ghttp.Request) error {
	requestModel := &request.AdminIds{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range requestModel.Selected {
			if _, err := dao.Admin.TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// Info 管理员基本信息
func (s *adminService) Info(r *ghttp.Request) (re *response.AdminInfo, err error) {
	adminId := auth.GetAdminID(r)
	re = &response.AdminInfo{}
	err = dao.Admin.As("a").LeftJoin(dao.User.Table+" u", "u.id=a.user_id").Where("a.id=?", adminId).Fields("a.*,u.username").Scan(re)
	if err != nil {
		return nil, err
	}
	roleID := auth.GetRoleID(r)
	menuList := auth.RoleMenu(r)
	rolePermission, err := auth.RolePermission(gconv.String(roleID))
	if err != nil {
		return nil, err
	}

	adminEntity := &entity.Admin{}
	err = dao.Admin.Where(dao.Admin.Columns.Id, adminId).Scan(adminEntity)
	if err != nil {
		return nil, err
	}
	re.Menu = menuList
	re.AdminPermission = rolePermission
	return re, nil
}
