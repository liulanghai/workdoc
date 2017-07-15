/**
 * @brief   : 打印文件日志接口
 * @author  : mxl
 * @date    : 2017-7-11
 */

package logc

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logc struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errLog   *log.Logger
	file     *os.File
	console  bool
}

/**
 * @brief  : 初始化日志文件
 * @param  : [in] basePath 日志保存目录
 * @param  : [in] moduleName 模块名称
 * @param  : [in] console 是否打屏
 * @return : 返回日志句柄
 */
func InitLogc(basePath string, moduleName string, console bool) (*Logc, error) {

	fileName := fmt.Sprintf("%s/%s_%s.log", basePath, moduleName,
		time.Now().Format("20060102150405"))

	logFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error:", err.Error())
		return nil, err
	}

	debugLog := log.New(logFile, "["+moduleName+"][Debug]", log.LstdFlags)
	infoLog := log.New(logFile, "["+moduleName+"][Info ]", log.LstdFlags)
	warnLog := log.New(logFile, "["+moduleName+"][Warn ]", log.LstdFlags)
	errLog := log.New(logFile, "["+moduleName+"][Error]", log.LstdFlags)

	logc := Logc{
		file:     logFile,
		debugLog: debugLog,
		infoLog:  infoLog,
		warnLog:  warnLog,
		errLog:   errLog,
		console:  console,
	}
	return &logc, nil
}

/**
 * @brief  : 关闭文件句柄
 * @param  : -
 * @return : void
 */
func (l *Logc) CloseLogc() {
	l.file.Close()
}

/**
 * @brief  : 打印debug日志
 * @param  : -
 * @return : void
 */
func (l *Logc) Debug(format string, v ...interface{}) {
	if l.console {
		log.Printf(fmt.Sprintf(format, v...))
	} else {
		l.debugLog.Printf(fmt.Sprintf(format, v...))
	}
}

/**
 * @brief  : 打印info日志
 * @param  : -
 * @return : void
 */
func (l *Logc) Info(format string, v ...interface{}) {
	if l.console {
		log.Printf(fmt.Sprintf(format, v...))
	} else {
		l.infoLog.Printf(fmt.Sprintf(format, v...))
	}
}

/**
 * @brief  : 打印warn日志
 * @param  : -
 * @return : void
 */
func (l *Logc) Warn(format string, v ...interface{}) {
	if l.console {
		log.Printf(fmt.Sprintf(format, v...))
	} else {
		l.warnLog.Printf(fmt.Sprintf(format, v...))
	}
}

/**
 * @brief  : 打印error日志
 * @param  : -
 * @return : void
 */
func (l *Logc) Error(format string, v ...interface{}) {
	if l.console {
		log.Printf(fmt.Sprintf(format, v...))
	} else {
		l.errLog.Printf(fmt.Sprintf(format, v...))
	}
}
