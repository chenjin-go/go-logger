package logger

type ihook interface {
	Call(l *Clog)
}
