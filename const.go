package tlogs

import "time"

const (
	default_topic  = "default_topic" 	// 默认主题
	default_dir = "/tmp/"				// 默认路径
	default_duration = 24 * time.Hour	// 默认分割间隔
	default_format = "%Y-%m-%d"			// 默认日志显示格式
)

