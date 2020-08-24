package main

import (
	"log"
	"os"
	"time"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// main 是程序的入口
func main() {
	log.Println("Starting work")

	// 为本次执行分配超时时间
	r := New(timeout)

	// 加入要执行分配超时时间
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Termination due timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terrination due to interrupt . ")
			os.Exit(2)
		}
	}

	log.Println("Process ended. ")

}

// createTask 返回一个根据id 休眠指定秒的实例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
