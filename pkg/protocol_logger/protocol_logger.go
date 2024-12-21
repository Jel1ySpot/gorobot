package protocol_logger

import (
	"fmt"
	"github.com/LagrangeDev/LagrangeGo/utils"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

var (
	DumpsPath = "dump"
	LogLevel  = logrus.InfoLevel
)

// ProtocolLogger from https://github.com/Mrs4s/go-cqhttp/blob/a5923f179b360331786a6509eb33481e775a7bd1/cmd/gocq/main.go#L501
type ProtocolLogger struct {
	fromProtocol string
}

func NewProtocolLogger(fromProtocol string) ProtocolLogger {
	return ProtocolLogger{fromProtocol: fromProtocol}
}

func (p ProtocolLogger) Info(format string, arg ...any) {
	logger.Infof(p.fromProtocol+format, arg...)
}

func (p ProtocolLogger) Warning(format string, arg ...any) {
	logger.Warnf(p.fromProtocol+format, arg...)
}

func (p ProtocolLogger) Debug(format string, arg ...any) {
	logger.Debugf(p.fromProtocol+format, arg...)
}

func (p ProtocolLogger) Error(format string, arg ...any) {
	logger.Errorf(p.fromProtocol+format, arg...)
}

func (p ProtocolLogger) Dump(data []byte, format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	if _, err := os.Stat(DumpsPath); err != nil {
		err = os.MkdirAll(DumpsPath, 0o755)
		if err != nil {
			logger.Errorf("出现错误 %v. 详细信息转储失败", msg)
			return
		}
	}
	dumpFile := path.Join(DumpsPath, fmt.Sprintf("%v.dump", time.Now().Unix()))
	logger.Errorf("出现错误 %v. 详细信息已转储至文件 %v 请连同日志提交给开发者处理", msg, dumpFile)
	_ = os.WriteFile(dumpFile, data, 0o644)
}

const (
	// 定义颜色代码
	colorReset  = "\x1b[0m"
	colorRed    = "\x1b[31m"
	colorYellow = "\x1b[33m"
	colorGreen  = "\x1b[32m"
	colorBlue   = "\x1b[34m"
	colorWhite  = "\x1b[37m"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(LogLevel)
	logger.SetFormatter(&ColoredFormatter{})
	logger.SetOutput(colorable.NewColorableStdout())
}

type ColoredFormatter struct{}

func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取当前时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// 根据日志级别设置相应的颜色
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = colorBlue
	case logrus.InfoLevel:
		levelColor = colorGreen
	case logrus.WarnLevel:
		levelColor = colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = colorRed
	default:
		levelColor = colorWhite
	}

	return utils.S2B(fmt.Sprintf("[%s] [%s%s%s]: %s\n",
		timestamp, levelColor, strings.ToUpper(entry.Level.String()), colorReset, entry.Message)), nil
}
