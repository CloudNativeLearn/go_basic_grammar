package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	macGoroutines   = 25 // 要使用地goroutine地数量
	pooledResources = 2  // 池中地资源数量
)

// dbConnection模拟要共享地资源
type dbConnection struct {
	ID int32
}


//Close实现leio.Close接口,以便dbConnection可以被池管理.
// Close用来完成任意资源地释放管理
func (dbConn *dbConnection)Close()error  {
	log.Println("Close:Connection",dbConn.ID)
	return nil
}

// idCounter用来给每个连接分配独一无二地id
var idCounter int32

// cerateConnection是一个工厂函数
// 当需要一个新链接时,资源池会调用这个函数
func createConnection()(io.Closer,error)  {
	id:=atomic.AddInt32(&idCounter,1)
	log.Println("Create:New Connection",id)
	return &dbConnection{id},nil

}


//  performQueries用来测试链接的资源池
func performQueries(query int,p *Pool)  {
	// 从池里请求一个链接
	conn,err := p.Acquire()
	if err != nil{
		log.Println(err)
		return
	}
	// 将该链接释放回池里
	defer p.Release(conn)
	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)

}

// main是所有Go程序的入口
func main()  {
	var wg sync.WaitGroup
	wg.Add(macGoroutines)

	// 创建用来管理链接的池
	p,err:=New(createConnection,pooledResources)
	if err!=nil{
		log.Println(err)
	}

	// 使用池里的链接来完成查询
	for query :=0;query<macGoroutines;query++{
		// 每个goroutine需要自己复制一份要查询值的副本 不然所有的查询会共享同一个查询
		go func(q int) {
			performQueries(q,p)
		}(query)

	}

	// 等待 goroutine结束
	wg.Wait()

	// 关闭池
	log.Println("Shutdown Program .")
	p.Close()
}

