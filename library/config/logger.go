package config

type Logger struct {
	Type       string `json:"type"`
	Level      string `json:"level"`
	Path       string `json:"path"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}
