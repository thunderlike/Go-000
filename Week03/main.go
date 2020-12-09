package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, groupCTX := errgroup.WithContext(ctx)
	g.Go(func() error {
		return registerSign(groupCTX)
	})
	g.Go(func() error {
		addr := ":8888"
		return startNewServer(groupCTX, addr, &appHandler{}, 5)
	})
	g.Go(func() error {
		addr := ":9999"
		return startNewServer(groupCTX, addr, &debugHandler{}, 10)
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("子协程返回 : %v\n", err.Error())
	}
	time.Sleep(10 * time.Second)
}

func registerSign(ctx context.Context) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	fmt.Printf("registerSign 开始 ！\n")
	for {
		select {
		case si := <-sig:
			//发信号中止http server
			fmt.Printf("registerSign 获得信号 %v \n", si)
			return fmt.Errorf("获得信号 %v", si)
		case <-ctx.Done():
			fmt.Printf("registerSign被其他协程结束了 \n")
			return fmt.Errorf("registerSign被其他协程结束了")
		default:
			fmt.Printf("registerSign running !\n")
			time.Sleep(time.Second)
		}

	}
	//go func() {
	//	for {
	//		select {
	//		case si := <-sig:
	//			//发信号中止http server
	//			fmt.Printf("registerSign 获得信号 %v \n", si)
	//			return fmt.Errorf("获得信号 %v", si)
	//		case <-ctx.Done():
	//			return fmt.Errorf("registerSign被其他协程结束了")
	//		}
	//	}
	//}()
}

func startNewServer(ctx context.Context, addr string, h http.Handler, duration int) error {
	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func(ctx context.Context, duration int) {
		for i := 0; i < duration; i++ {
			select {
			default:
				fmt.Printf("http %v running 持续时间 %v \n", addr, duration)
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Printf("其他子协程关闭 %v 也要关闭了\n", addr)
				i = duration
			}
			if i == duration-1 {
				fmt.Printf("server %v 到时间了，自动停止 \n", addr)
			}
		}
		s.Shutdown(ctx)

	}(ctx, duration)

	fmt.Printf("http %v 开始 ！\n", addr)
	return s.ListenAndServe()
}

type appHandler struct {
}

func (h *appHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {

}

type debugHandler struct {
}

func (h *debugHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {

}
