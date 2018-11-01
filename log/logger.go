package log

import (
	//"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)

func init(){
	errFile,err:=os.OpenFile("../src/agenda/log/errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("cannot open log file",err)
	}

	infoFile,err:=os.OpenFile("../src/agenda/log/info.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("cannot open log file",err)
	}

	Info = log.New(infoFile,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(errFile,"Error:",log.Ldate | log.Ltime | log.Lshortfile)

}