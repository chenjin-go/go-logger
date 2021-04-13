package logger

type ibothook interface {
	Call(l *Clog)
}
