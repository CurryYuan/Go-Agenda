package log

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	GOPATH := os.Getenv("GOPATH")
	dname := GOPATH + "/src/agenda/log"
	os.MkdirAll(dname, os.ModeDir|os.ModePerm)
	errFile, err := os.OpenFile(dname+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("cannot open log file", err)
	}

	infoFile, err := os.OpenFile(dname+"/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("cannot open log file", err)
	}

	Info = log.New(infoFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
