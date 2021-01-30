package yzapi

//自定义日志接口
type MsgLoger interface {
	ToLogString() string
}
