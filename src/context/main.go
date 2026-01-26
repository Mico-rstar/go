package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个 goroutine 模拟工作
	go doWork(ctx)

	// 主 goroutine 等待 2 秒后发出取消信号
	time.Sleep(2 * time.Second)
	fmt.Println("主程序发出取消信号")
	cancel() // 发出取消信号

	// 再等一会儿，确保子 goroutine 收到取消信号
	time.Sleep(1 * time.Second)
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到取消信号，退出工作:", ctx.Err())
			return
		default:
			fmt.Println("正在工作...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}