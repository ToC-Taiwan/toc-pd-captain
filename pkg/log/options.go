// Package log package log
package log

type config struct {
	timeFormat string
	format     Format
	level      Level
	needCaller bool
	fileName   string
}

// Option -.
type Option func(*config)

func LogLevel(level Level) Option {
	return func(c *config) {
		if level == "" {
			return
		}

		if level != "panic" && level != "fatal" && level != "error" && level != "warn" && level != "info" && level != "debug" && level != "trace" {
			return
		}

		c.level = level
	}
}

func TimeFormat(format string) Option {
	return func(c *config) {
		c.timeFormat = format
	}
}

func LogFormat(format Format) Option {
	return func(c *config) {
		if format == "" {
			return
		}

		if format != "json" && format != "text" {
			return
		}

		c.format = format
	}
}

func NeedCaller(need bool) Option {
	return func(c *config) {
		c.needCaller = need
	}
}

func FileName(name string) Option {
	return func(c *config) {
		c.fileName = name
	}
}
