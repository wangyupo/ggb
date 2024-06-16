package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	response2 "github.com/wangyupo/GGB/model/system/response"
	"github.com/wangyupo/GGB/utils"
)

// GetLoginLogList 列表
func GetLoginLogList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userId := c.Query("userId")

	// 声明 log.SysLogLogin 类型的变量以存储查询结果
	loginLogList := make([]response2.LoginLogResponse, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysLogLogin{})
	if userId != "" {
		db = db.Where("user_id = ?", "%"+userId+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Table("sys_log_login").
		Select("sys_log_login.*, sys_user.user_name").
		Joins("JOIN sys_user ON sys_log_login.user_id = sys_user.id").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&loginLogList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  loginLogList,
		Total: total,
	}, c)
}