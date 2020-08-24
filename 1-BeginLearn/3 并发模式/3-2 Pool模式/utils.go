package main

import (
	"errors"
	"io"
	"log"
	"sync"
)

// 本章会介绍pool包 ①
//① 本书是以 Go 1.5 版本为基础写作而成的。在 Go 1.6 及之后的版本中，标准库里自带了资源池的实现
//（sync.Pool）。推荐使用。——译者注
//。这个包用于展示如何使用有缓冲的通道实现资源池，来管理可以在
//任意数量的goroutine之间共享及独立使用的资源。这种模式在需要共享一组静态资源的情况（如
//共享数据库连接或者内存缓冲区）下非 常有用。如果goroutine需要从池里得到这些资源中的一个，
//它可以从池里申请，使用完后归还到资源池里。

// Pool 管理一组可以安全地在多个goroutine间
// 共享地资源。被管理的资源必须
// 实现io.close接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrpoolClosed表示请求（Acquire）了一个
// 已经关闭地池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New创建一个用来管理资源地池
// 这个池需要一个可以分配新资源地函数
// 并规定池地大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获得一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检测是否有空闲地资源
	case r, ok := <-p.resources:
		log.Println("Acquire：", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 因为没有空闲资源可用，所以提供一个新的资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}

}

// Release将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作地安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果池被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}
	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// 如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closeing")
		r.Close()
	}
}


// Close 会让资源池停止工作 ， 并且关闭所有现在地资源
func (p *Pool) Close() {
	// 保证本操作与Release操作地安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果pool已经关闭，什么也不做
	if p.closed{
		return
	}
	// 将池关闭
	p.closed = true
	// 在清空管道李地资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r:= range p.resources{
		r.Close()
	}
}

