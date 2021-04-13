package logger

//钩子接口

type ibothook interface {
	BotCall(l *Clog)
}

type itophook interface {
	TopCall(l *Clog)
}
