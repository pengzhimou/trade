package spidernet

import (
	"fmt"
	"time"
	"trade/config"
	q "trade/quant"

	// . "trade/utils"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

var (
	huobisystem = q.HuobiSystem{
		Client: new(client.CommonClient).Init(config.Host),
	}
	huobimarket = q.HuobiMarket{
		Client: new(client.MarketClient).Init(config.Host),
	}
	huobialgo = q.HuobiAlgo{
		Client: new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host),
	}
)

// 定义一个任务类型task
type Task struct {
	f func() error
}

// task 执行业务的方法
func (t *Task) Execute() {
	t.f()
}

// 创建一个task任务
func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

// 定义协程池
type Pool struct {
	// 对外的task入口
	EntryChannel chan *Task

	// 内部的Task队列
	JobsChannel chan *Task

	// 协程池最大的worker数量
	worker_num int
}

// 让协程池开始工作
func (p *Pool) run() {
	// 1根据work num来创建worker来工作
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i)
	}
	// 2Jobs是从Entry入口取的任务
	for task := range p.EntryChannel {
		// 一旦有task读到
		p.JobsChannel <- task
	}
}

// 协程池创建一个worker，并且让这个worker去工作
func (p *Pool) worker(worker_ID int) {
	// 一个worker具体的工作

	// 1 永久从Jobschannel取任务
	for task := range p.JobsChannel {
		//task就是当前woerk从jobschannel拿到的任务
		fmt.Print("worker id:", worker_ID, "开始执行。")
		task.Execute()
		fmt.Println("worker id:", worker_ID, "执行完了")
	}

	// 2 一旦取到任务就执行任务
}

// 创建Pool的函数
func NewPool(cap int) *Pool {
	// 创建pool
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		worker_num:   cap,
	}
	// 返回这个pool
	return &p
}

// 主函数 来测试协程池的工作
func Run() {
	// 1创建一些任务
	// t := NewTask(
	// 	// func() error {
	// 	// 	fmt.Println(time.Now())
	// 	// 	time.Sleep(time.Second * time.Duration(rand.Intn(4)))
	// 	// 	return nil
	// 	// },
	// 	func(sn SpiderNet, stk q.Stock) error {
	// 		sn.GetStockRiseFall(stk, &huobimarket, market.DAY1)
	// 		return nil
	// 	}, sn, stk,
	// )

	//2创建一个协程池
	p := NewPool(3)

	go func() {
		sn := SpiderNet{}
		for {
			stocks := sn.GetAllAvaliableStock(&huobisystem, &huobimarket)
			for stock := range stocks {
				stk := q.Stock{
					Name: stock,
				}
				t := NewTask(func() error {
					fmt.Println(sn.GetStockRiseFall(stk, &huobimarket, market.DAY1))
					time.Sleep(time.Second * 1)
					return nil
				})
				p.EntryChannel <- t
			}
		}
	}()

	//3将这些任务交给协程池
	// go func() {
	// 	for {
	// 		p.EntryChannel <- t
	// 		// time.Sleep(time.Second * 5)
	// 	}
	// }()

	//4启动pool
	p.run()
}
