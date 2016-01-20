package logger

type Logger interface {
	Debug()
	Debugf()
	Warn()
	Warnf()
	Error()
	Errorf()
}
