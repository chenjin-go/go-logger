package logger

type ibothook interface {
	BotCall(l *Clog)
}

type itophook interface {
	TopCall(l *Clog)
}
