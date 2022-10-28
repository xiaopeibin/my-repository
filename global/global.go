package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"my_go_project/config"
)

var (
	GVA_CONFIG config.Server
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
)
