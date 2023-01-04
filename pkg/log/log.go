// Package log package log
package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	_defaultLogFormat  = FormatText
	_defaultLogLevel   = LevelInfo
	_defaultNeedCaller = false
	_defaultTimeFormat = "2006-01-02 15:04:05"
	_defaultFilePath   = "./logs"
	_defaultFileName   = "log"
)

// Log -.
type Log struct {
	*logrus.Logger
	*config

	basePath string
}

var singleton *Log

// New -.
func New(opts ...Option) *Log {
	if singleton != nil {
		if len(opts) > 0 {
			singleton.Warn("logger already initialized, options will be ignored")
		}
		return singleton
	}

	l := &Log{
		Logger: logrus.New(),
		config: &config{
			timeFormat: _defaultTimeFormat,
			format:     _defaultLogFormat,
			level:      _defaultLogLevel,
			needCaller: _defaultNeedCaller,
			fileName:   _defaultFileName,
		},
		basePath: getExecPath(),
	}

	for _, opt := range opts {
		opt(l.config)
	}

	jsonFormatter := &logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   l.timeFormat,
		PrettyPrint:       false,
		CallerPrettyfier:  l.setCallerPrettyfier(true),
	}

	textFormatter := &logrus.TextFormatter{
		TimestampFormat:  l.timeFormat,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		PadLevelText:     false,
		ForceColors:      true,
		CallerPrettyfier: l.setCallerPrettyfier(false),
	}

	var formatter logrus.Formatter
	if l.format == FormatJSON {
		formatter = jsonFormatter
	} else {
		formatter = textFormatter
	}

	if l.needCaller {
		l.SetReportCaller(true)
	}

	l.Hooks.Add(l.fileHook(jsonFormatter))
	l.SetFormatter(formatter)
	l.SetLevel(l.level.Level())
	l.SetOutput(os.Stdout)

	singleton = l
	return singleton
}

func (l *Log) fileHook(formatter logrus.Formatter) *lfshook.LfsHook {
	return lfshook.NewHook(
		filepath.Join(l.basePath, fmt.Sprintf("%s/%s-%s.log", _defaultFilePath, l.fileName, time.Now().Format("20060102"))),
		formatter,
	)
}

func (l *Log) setCallerPrettyfier(isJSON bool) func(*runtime.Frame) (function string, file string) {
	return func(frame *runtime.Frame) (function string, file string) {
		fileName := strings.ReplaceAll(frame.File, fmt.Sprintf("%s/", l.basePath), "")
		if isJSON {
			return fmt.Sprintf("%s:%d", fileName, frame.Line), ""
		}
		return fmt.Sprintf("[%s:%d]", fileName, frame.Line), ""
	}
}
