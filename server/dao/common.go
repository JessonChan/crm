package dao

import (
	"context"
	"crm/global"
	"crm/models"
)

const (

	// 数据库表名
	USER        = "user"
	CUSTOMER    = "customer"
	CONTRACT    = "contract"
	PRODUCT     = "product"
	SUBSCRIBE   = "subscribe"
	NOTICE      = "notice"
	MAIL_CONFIG = "mail_config"

	// 空值
	NumberNull = 0
	StringNull = ""
)

var ctx = context.Background()

// RestPage 分页查询
// page  设置起始页、每页条数,
// name  查询目标表的名称
// query 查询条件,
// dest  查询结果绑定的结构体,
// bind  绑定表结构对应的结构体
func restPage(page models.Page, name string, query interface{}, dest interface{}, bind interface{}) (int64, error) {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		global.Db.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	}
	res := global.Db.Table(name).Where(query).Find(bind)
	return res.RowsAffected, res.Error
}
