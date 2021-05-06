package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// runner 包用于展示如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以
//用 runner 包来终止程序。当开发需要调度后台处理任务的程序的时候，这种模式会很有用。这
//个程序可能会作为 cron 作业执行，或者在基于定时任务的云环境（如 iron.io）里执行。
//让我们来看一下 runner 包里的 runner.go 代码文件
// runner 包管理处理任务的运行和生命周期

// Runner在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统
	// 发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout 报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	task []func(int)
}

// ErrTimeout 会在任务执行超时的时候返回
var ErrTimeout = errors.New("received time out")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的准备使用Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将一个任务加到Runner上，这个任务是一个 接收int类型的id作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.task = append(r.task, tasks...)
}

// Start执行所有任务，并监视通道的任务
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的gorout执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时候发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrInterrupt
	}
}

// run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.task {
		// 检查操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已经注册的任务
		task(id)
	}
	return nil
}

// gotInterrupt验证是否接收中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被促发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信息
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false

	}

}
