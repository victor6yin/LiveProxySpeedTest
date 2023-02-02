package global

import (
	"LiveProxySpeedTest/pkg/setting"

	"go.uber.org/zap"
)

var (
	LogSetting *setting.LogSettingS
	Logger     *zap.Logger
)
