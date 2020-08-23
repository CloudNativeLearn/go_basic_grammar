package main

import (
	"fmt"
	"time"
)

// -----------------------------Task类型-----------------------
// 定义一个任务类型Task
type Task struct {
	f func() error // 一个任务类型 应该有一个具体的任务
}

// 创建一个任务
func NewTask(arg_f func() error) *Task {
	p := Task{
		f: arg_f,
	}
	return &p
}

// Task也需要一执行方法
func (t *Task) Excecute() {
	t.f()
}

// ----------------------------- 关于携程池 pool 角色的功能-----------------------
// 定义一个pool
type Pool struct {
	// 对外的Task入口
	EntryChannel chan *Task
	// 内部Task入口
	JobsChannel chan *Task
	// 协程池中的work数量
	worker_num int
}

// 创建一个poll的函数
func NewPoll(num int) *Pool {
		return &Pool{
			EntryChannel:make(chan *Task),
			JobsChannel:make(chan *Task),
			worker_num: num,
		}
}

// 协程池创建一个work，并且让Work工作
func (p *Pool) Worker(work_ID int) {
	// 永久的从JobsChannel 取数据
	for task := range p.JobsChannel{
		task.Excecute()
		fmt.Println("Word--",work_ID,"完成任务")
	}

}

// 让携程池，真正开始工作，协程池的一个启动方法
func (p *Pool) run() {
	// 1 根据 work_num 创建work去g工作
	for i:=0;i<p.worker_num;i++{
		go p.Worker(i)
	}

	// 读取从入口进入的队列  把队列的事件 传递给JobsChannel
	for task:=range p.EntryChannel{
		// 一但有task读到
		p.JobsChannel <- task
	}
}



// ----------------------------- 协程池真正开始工作-----------------------

func main() {
	// 1  创建一些任务
	t := NewTask(func() error {
		fmt.Println("现在的时间",time.Now())
		return nil
	})
	// 2  创建一个pool池子
	pool := NewPoll(4)

	// 3 把这些任务交给pool的入口
	go func(){
		task_num :=0
		for {
			task_num+=1
			pool.EntryChannel <- t
			fmt.Println("正在执行第",task_num,"任务")

		}
	}()

	// 4 开始运行
	pool.run()





}


