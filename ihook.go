package logger

type ibothook interface {
	Call(l *Clog)
}

type itophook interface {
	Call(l *Clog)
}
