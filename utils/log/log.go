package log

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"runtime"
	"strconv"
)

var (
	ip     string
	logger *logrus.Logger
)

// 获取本地 IPv4 地址
func getLocalIP() (string, error) {
	// 获取环境变量clientIp
	clientIp, exists := os.LookupEnv("CLIENT_IP")
	if exists {
		return clientIp, nil
	}
	// 环境变量clientIp不存在则获取本机ipv4
	ipv4, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range ipv4 {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}
	return "", errors.New("not found ipv4 address")
}

func init() {
	// 获取本机ip
	ip, _ = getLocalIP()
	fmt.Print("init logger object\n")

	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})
}

func printStackTrace(skip int) (string, string, string, error) {
	/*

	 */
	// 获取调用栈信息
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		line := strconv.Itoa(line)
		return file, line, funcName, nil
	}
	return "1", "0", "1", errors.New("not found call stack information")
}

func HandleExtraFields() logrus.Fields {
	file, line, funcName, err := printStackTrace(3)
	fields := logrus.Fields{"ip": ip}
	if err == nil {
		fields["file"] = file + ":" + line
		fields["func"] = funcName
	}
	return fields
}

func Info(msg string) {
	// 获取调用栈信息
	fields := HandleExtraFields()
	logger.WithFields(fields).Info(msg)
}

func Warning(msg string) {
	// 获取调用栈信息
	fields := HandleExtraFields()
	logger.WithFields(fields).Warning(msg)
}

func Error(msg string) {
	// 获取调用栈信息
	fields := HandleExtraFields()
	logger.WithFields(fields).Error(msg)
}
