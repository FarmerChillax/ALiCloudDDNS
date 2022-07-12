package log

import (
	"io"
	"log"
	"os"
	"path"
)

type FileLog string

// 写日志到文件
func (fl FileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func SetLoggerWithFile(filename string) {
	dir, _ := path.Split(filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("日志初始化失败，请检查路径: %s", dir)
		return
	}

	log.SetOutput(FileLog(filename))
}

func SetLogger(out io.Writer) {
	log.SetOutput(out)
}
