// 主入口

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

// StartHttpServer 启动 HTTP server
func StartHttpServer(s *http.Server) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello laffey")
	})
	err := s.ListenAndServe()
	return err
}

// main 主函数
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)

	//http server
	s := &http.Server{Addr: ":8080"}

	group.Go(func() error {
		return StartHttpServer(s)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return s.Shutdown(errCtx) // 关闭 http server
	})

	// 注册信号量
	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // cancel、timeout、deadline 等
				return errCtx.Err()
			case <-chanel: // 终止信号
				cancel()
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("err group error: ", err)
	}
	fmt.Println("main exit")
}
