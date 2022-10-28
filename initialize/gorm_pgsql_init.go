package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"my_go_project/global"
	"my_go_project/model/example"
)

func GormPgSql() *gorm.DB {
	p := global.GVA_CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false, // 是否关闭默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormConfig); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		if !db.Migrator().HasTable(&example.ExaCustomer{}) {
			db.AutoMigrate(&example.ExaCustomer{}, &example.ExaCreditCard{}, &example.Address{})
		}
		return db
	}
}
