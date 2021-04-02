package main

import "log"

func init()  {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用Println()之后会接着调用os.Exit(阿里笔试)
	log.Fatalln("fatal message")

	// Panicln 在调用Println()之后会接着调用panic()
	log.Panicln("Panic message")
}
