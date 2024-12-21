package bot

import (
	"fmt"
	"github.com/Jel1ySpot/gorobot/pkg/protocol_logger"
	"github.com/LagrangeDev/LagrangeGo/utils/log"
)

type Logger struct {
	protocolLogger protocol_logger.ProtocolLogger
}

func NewLogger(fromProtocol string) Logger {
	return Logger{
		protocolLogger: protocol_logger.NewProtocolLogger(fromProtocol),
	}
}

func (l *Logger) Info(format string, args ...any) {
	l.protocolLogger.Info(log.Getcaller(format), args...)
}

func (l *Logger) Infoln(msgs ...any) {
	l.protocolLogger.Info(log.Getcaller(fmt.Sprint(msgs...)))
}

func (l *Logger) Warning(format string, args ...any) {
	l.protocolLogger.Warning(log.Getcaller(format), args...)
}

func (l *Logger) Warningln(msgs ...any) {
	l.protocolLogger.Warning(log.Getcaller(fmt.Sprint(msgs...)))
}

func (l *Logger) Error(format string, args ...any) {
	l.protocolLogger.Error(log.Getcaller(format), args...)
}

func (l *Logger) Errorln(msgs ...any) {
	l.protocolLogger.Error(log.Getcaller(fmt.Sprint(msgs...)))
}

func (l *Logger) Debug(format string, args ...any) {
	l.protocolLogger.Debug(log.Getcaller(format), args...)
}

func (l *Logger) Debugln(msgs ...any) {
	l.protocolLogger.Debug(log.Getcaller(fmt.Sprint(msgs...)))
}

func (l *Logger) Dump(format string, data []byte, args ...any) {
	l.protocolLogger.Dump(data, log.Getcaller(format), args...)
}
