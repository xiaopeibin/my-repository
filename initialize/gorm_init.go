package initialize

import (
	"gorm.io/gorm"
	"my_go_project/global"
)

func Gorm() *gorm.DB {
	dbType := global.GVA_CONFIG.System.DbType
	switch dbType {
	case "pgsql":
		return GormPgSql()
	case "mysql":
		return GormMysql()
	default:
		return nil
	}
}
