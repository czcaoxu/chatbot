package logger

type Config struct {
	Level      string // debug, info, warn, error
	Encoding   string // json or console
	OutputPath string // e.g., stdout or logs/app.log
	DevMode    bool   // true for human-readable logs
}
