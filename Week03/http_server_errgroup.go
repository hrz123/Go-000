package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// err group for safe exit
	g, ctx := errgroup.WithContext(context.Background())

	// create a svr
	svr := &http.Server{
		Addr:    ":8001", // port
		Handler: nil,
	}

	// signal通信
	sig := make(chan os.Signal, 1)

	// 启动服务
	g.Go(func() error {
		go func() {
			// 等待ctx取消或者收到signal，就要关掉svr
			select {
			case <-ctx.Done():
				log.Println("http ctx done")
			case <-sig:
				log.Println("get shutdown signal")
			}
			svr.Shutdown(context.TODO())
		}()
		return svr.ListenAndServe()
	})

	// 监听linux signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		// 监听linux的signal
		signal.Notify(sig, exitSignals...)
		return nil
	})

	// 注入错误
	g.Go(func() error {
		log.Println("inject")
		time.Sleep(time.Second * 3)
		log.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	log.Println(err)

}
