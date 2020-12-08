package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// err group for safe exit
	g, ctx := errgroup.WithContext(context.Background())

	// 监听linux signal
	c := make(chan os.Signal)

	// create a server
	server := &http.Server{
		Addr:    ":8001",
		Handler: nil,
	}

	// register a cleanup function
	server.RegisterOnShutdown(func() {
		log.Println("RegisterOnShutdown(): completed!")
	})

	// 启动服务
	g.Go(func() error {
		err := server.ListenAndServe()
		return err
	})
	log.Println("Server started.")

	// 监听linux signal
	g.Go(func() error {
		signal.Notify(c, os.Interrupt)
		return nil
	})

	// 等待监听到linux signal
	osCall := <-c
	// get signal
	log.Printf("signal %+v", osCall)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 关闭服务器
	server.Shutdown(ctx)

	if err := g.Wait(); err != nil && err != http.ErrServerClosed {
		log.Printf("server start failed or server shutdown failed: %+v", err)
	} else {
		log.Printf("server shutdown properly")
	}

}
