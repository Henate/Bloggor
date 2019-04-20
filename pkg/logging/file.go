package logging

import (
	"os"
	"time"
	"fmt"
	"log"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	dir, _ := os.Getwd()
	path := dir + "/" + LogSavePath
	log.Printf("Log file path : %s", path)
	return fmt.Sprintf("%s", path)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)	//返回文件信息结构描述文件。如果出现错误，会返回*PathError
	switch {
	case os.IsNotExist(err):	//能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
		mkDir()
	case os.IsPermission(err):	//能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
		log.Fatalf("Permission :%v", err)
	}

	//OpenFile调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于I/O。如果出现错误，则为*PathError
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()	//返回与当前目录对应的根路径名
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)	//创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error
	if err != nil {
		panic(err)
	}
}