package tlogs

type Logs interface {
	Info(args... interface{})
}

var _tLogger Logs

func SetLogger(l Logs) {
	_tLogger = l
}

func GetLogs() Logs {
	if _tLogger == nil {
		//SetLogger(GetJsonDLog())
	}
	return _tLogger
}

