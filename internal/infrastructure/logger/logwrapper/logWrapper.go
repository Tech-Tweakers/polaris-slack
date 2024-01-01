package logwrapper

import (
	"go.uber.org/zap"
)

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

// LoggerWrapper is a wrapper for zap.Logger + Span handling
type LoggerWrapper interface {
	SetTraceID(v string) LoggerWrapper
	SetVersion(v string) LoggerWrapper
	CreateSpan() LoggerWrapper
	RemoveSpan() LoggerWrapper
	Logger // interface extends logger
	TraceID() string
	Version() string
	Span() *Span
}

type logWrapper struct {
	logger  Logger
	traceID string
	span    *Span
	version string
}

// New returns a new logger
func New(logger Logger) LoggerWrapper {
	return &logWrapper{
		logger: logger,
	}
}

func (l *logWrapper) SetTraceID(v string) LoggerWrapper {
	l.traceID = v
	return l.clone()
}

func (l *logWrapper) TraceID() string {
	return l.traceID
}

func (l *logWrapper) SetVersion(v string) LoggerWrapper {
	l.version = v
	return l.clone()
}

func (l *logWrapper) Version() string {
	return l.version
}

func (l *logWrapper) Span() *Span {
	return l.span
}

func (l *logWrapper) CreateSpan() LoggerWrapper {
	l.span = createSpan(l.span)
	return l
}

func (l *logWrapper) RemoveSpan() LoggerWrapper {
	if l.span != nil {
		l.span = l.span.parent
	}
	return l
}

func (l *logWrapper) Info(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Info(msg, f...)
}

func (l *logWrapper) Warn(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Warn(msg, f...)
}

func (l *logWrapper) Error(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Error(msg, f...)
}

func (l *logWrapper) Fatal(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Fatal(msg, f...)
}

func (l *logWrapper) Debug(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Debug(msg, f...)
}

func (l *logWrapper) clone() LoggerWrapper {
	return &logWrapper{
		logger:  l.logger,
		traceID: l.traceID,
		span:    l.span,
		version: l.version,
	}
}

func (l *logWrapper) mergeField(fields ...zap.Field) []zap.Field {

	parentID := ""
	spanID := ""
	if l.span != nil {
		spanID = l.span.ID
		if l.span.parent != nil {
			parentID = l.span.parent.ID
		}
	}
	s := []zap.Field{
		zap.String("version", l.version),
		zap.String("trace_id", l.traceID),
		zap.String("span_parent_id", parentID),
		zap.String("span_id", spanID),
	}

	s = append(s, fields...)

	return s
}
