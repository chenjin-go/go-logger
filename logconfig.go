package logger

type LogConfig struct {
	AllLog   string
	InforLog string
	ErrorLog string
	WarnLog  string
}

var config = &LogConfig{}

func init() {

}
