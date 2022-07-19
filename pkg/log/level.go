package log

type Level int

const (
	Debug Level = iota - 1
	Info
	Warn
	Error
	Fatal
	Panic
)

func getLevel(lvl string) Level {
	m := map[string]Level{
		"DEBUG": Debug,
		"INFO":  Info,
		"WARN":  Warn,
		"ERROR": Error,
		"FATAL": Fatal,
		"PANIC": Panic,
	}
	return m[lvl]
}
