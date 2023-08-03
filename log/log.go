package log

import (
	"ISPS/config"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// consoleLogger 是终端结构体
type consoleLogger struct {
	level logLevel
}

// fileLogger 是文件结构体
type fileLogger struct {
	level       logLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *logMsg
}
type logMsg struct {
	level     logLevel
	msg       string
	funcName  string
	fileName  string
	timeStamp string
	line      int
}

// ConsoleFileLogger 是终端和文件的结构体
type consoleFileLogger struct {
	consolelogger consoleLogger
	filelogger    *fileLogger
}

/*功能：构造函数
 *参1：日志等级
 *参2：文件路径
 *参3：文件名称
 *参4：文件最大容量
 */
func NewLog() *consoleFileLogger {
	y := config.Yaml{}
	level, err := parseLogLevel(y.ReadYamlString("log.levelstr"))
	if err != nil {
		panic(err)
	}
	f1 := &consoleFileLogger{
		consolelogger: consoleLogger{
			level: level,
		},
		filelogger: &fileLogger{
			level:       level,
			filePath:    y.ReadYamlString("log.fp"),
			fileName:    y.ReadYamlString("log.fn"),
			maxFileSize: y.ReadYamlInt64("log.maxFileSize"),
			logChan:     make(chan *logMsg, 50000),
		},
	}
	err = f1.filelogger.initFile()
	if err != nil {
		panic(err)
	}

	return f1
}

/*功能：调用日志
 *参1：日志等级
 *参2：向端口输出还是向文件输出
 */
func (f *consoleFileLogger) Info(format, wher string) {
	if wher == "filepath" || wher == "" {
		f.filelogger.log(INFO, format)
	}
	if wher == "consolepath" {
		f.consolelogger.log(INFO, format)
	}
}
func (f *consoleFileLogger) Debug(format, wher string) {
	if wher == "filepath" || wher == "" {
		f.filelogger.log(DEBUG, format)
	}
	if wher == "consolepath" {
		f.consolelogger.log(DEBUG, format)
	}
}
func (f *consoleFileLogger) Warn(format, wher string) {
	if wher == "filepath" || wher == "" {
		f.filelogger.log(WARN, format)
	}
	if wher == "consolepath" {
		f.consolelogger.log(WARN, format)
	}
}
func (f *consoleFileLogger) Error(format, wher string) {
	if wher == "filepath" || wher == "" {
		f.filelogger.log(ERROR, format)
	}
	if wher == "consolepath" {
		f.consolelogger.log(ERROR, format)
	}
}
func (f *consoleFileLogger) Fatal(format, wher string) {
	if wher == "filepath" || wher == "" {
		f.filelogger.log(FATAL, format)
	}
	if wher == "consolepath" {
		f.consolelogger.log(FATAL, format)
	}
}

// 往终端写日志相关内容
func (c consoleLogger) enable(loglevel logLevel) bool {
	if loglevel > 5 {
		err := errors.New("invalid args")
		fmt.Println(err)
	}
	return c.level <= loglevel
}

func (c consoleLogger) log(lv logLevel, format string) {
	if lv > 5 {
		err := errors.New("invalid args")
		fmt.Println(err)
	}
	if c.enable(lv) {
		msg := format
		funcName, fileName, lineNum := getInfo(3)
		now := time.Now().Format("2006/01/02 15:04:05")
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), funcName, fileName, lineNum, msg)
	}
}

// 往文件里面写日志相关代码
func (f *fileLogger) enable(loglevel logLevel) bool {
	if loglevel > 5 {
		err := errors.New("invalid args")
		fmt.Println(err)
	}
	return f.level <= loglevel
}

// 记录日志方法
func (f *fileLogger) log(lv logLevel, format string) {
	if lv > 5 {
		err := errors.New("invalid args")
		fmt.Println(err)
	}
	if f.enable(lv) {
		msg := format
		funcName, fileName, lineNum := getInfo(3)
		now := time.Now().Format("2006/01/02 15:04:05")
		// 造一个logMsg对象
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timeStamp: now,
			line:      lineNum,
		}
		select {
		case f.logChan <- logTmp: //如果放不进去就走default
		default:
			// 如果收日志的goroutine服务断了，就丢掉日志不要再写了以免阻塞
		}
	}
}

func (f *fileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("1 open err log file failed, err:%v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	// 开启一个后台的goroutine写日志
	go f.writeLogBackground()
	return nil
}

// 关闭连接
// func (f *fileLogger) close() {
// 	f.fileObj.Close()
// 	f.errFileObj.Close()
// }

// 判断文件大小，若判断时打不开文件，就再打开一次
func (f *fileLogger) checkSize(file *os.File) bool {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("checkSize log file failed, err:%v\n", err)
	}
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *fileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割
	//1.备份一下，rename
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info file failed, err:%v", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())

	nowStr := time.Now().Format("20060102150405000")

	newlogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个备份文件名字

	//2.关闭当前文件
	file.Close()
	os.Rename(logName, newlogName)

	//3.打开一个新的日志文件
	retFileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("返回文件是", retFileObj.Name())
	//4.将打开的新日文件对象赋值给 f.fileObj
	return retFileObj, nil
}

// 写日志
func (f *fileLogger) writeLogBackground() error {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return err
			}
			f.fileObj = newFile
		}

		if f.checkSize(f.errFileObj) {
			newFile, err := f.splitFile(f.errFileObj)
			if err != nil {
				return err
			}
			f.errFileObj = newFile
		}

		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timeStamp, getLogString(logTmp.level), logTmp.funcName, logTmp.fileName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, "%s", logInfo)

			if logTmp.level >= ERROR {
				// 如果要记录的日志大于等于ERROR级别，我还要在err日志中再记录一遍
				fmt.Fprintf(f.errFileObj, "%s", logInfo)
			}
		default:
			// fmt.Println("--------取不到日志，缓5毫毛")
			// 取不出来日志就休息会
			time.Sleep(time.Millisecond * 500)
		}
	}
}

type logLevel uint16

const (
	//定义日志级别
	UNKNOWN logLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func parseLogLevel(s string) (logLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "unknown":
		return UNKNOWN, nil
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("日志级别错误")
		return UNKNOWN, err
	}
}

func getInfo(skip int) (funcName, fileName string, lineNum int) {
	pc, file, lineNum, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNum
}

func getLogString(lv logLevel) string {
	switch lv {
	case UNKNOWN:
		return "UNKNOWN"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "missing"
	}
}
