/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"guyu/cmd"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleSignal(ctx context.Context, cancel context.CancelFunc) chan os.Signal {
	out := make(chan os.Signal, 1)

	notify := make(chan os.Signal, 10)

	//signal.Notify(notify, os.Interrupt, os.Kill, syscall.SIGTERM)
	//signal.Notify(notify, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		defer close(out)
		for {
			//sig := <-notify
			//switch sig {
			//case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
			//	cancel()
			//	out <- sig
			//	return
			//default:
			//	log.Println("unhandled signal: ", sig)
			//}
			select {
			case <-ctx.Done():
				return
			case sig := <-notify:
				switch sig {
				case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
					cancel()
					out <- sig
					return
				default:
					log.Println("unhandled signal: ", sig)
				}
			}
		}
	}()
	return out
}

const PETSOTRE_BOT_APP_TOKEN = "xapp-1-A04HYC52C59-4595987458807-acf4702412f0048feda0929ecefdd3d506603c063ab6e30b3396a33a8468db43"
const SLACK_AUTH_TOKEN = "xoxb-4610391462931-4634227803968-Vlydb3gZxpWmIWjQKcCrZL0o"

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger = logger.Named("guyu-fw-test-app")
	logger.Info(
		"starting",
		zap.String("program", "guyu-fw-test-app"),
		zap.String("version", "0.0.1"),
		zap.Int("pid", os.Getpid()))

	var sigCh chan os.Signal
	//go func() {
	sigCh = handleSignal(ctx, cancel)
	//}()

	cmd.Execute(ctx)
	//cancel()

	//if sig := <-sigCh; sig == syscall.SIGQUIT || sig == syscall.SIGTERM {
	//	fmt.Println("cancelled")
	//}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigCh
		fmt.Println("cancelled")
		os.Exit(0)
	}()
	wg.Wait()
}
