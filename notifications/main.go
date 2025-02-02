package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/qaZar1/checkerNewGO/notifications/internal/api"
	"github.com/qaZar1/checkerNewGO/notifications/internal/bot"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})
}

func main() {
	const (
		token   = "TOKEN"
		address = "ADDRESS"
	)

	go bot.NewBot(os.Getenv(token), api.NewAPIUsers("http://localhost:8000/api"))
	// router := http.NewServeMux()
	// router.Handle("/telegram/", httpSwagger.Handler())

	// server := &http.Server{
	// 	Addr:    os.Getenv(address),
	// 	Handler: router,
	// }

	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		logrus.Panicf("Invalid service starting: %s", err)
	// 	}
	// 	logrus.Info("Service stopped")
	// }()
	// logrus.Info("Service started")

	// channel := make(chan os.Signal, 1)
	// signal.Notify(channel,
	// 	syscall.SIGABRT,
	// 	syscall.SIGHUP,
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT,
	// )
	// <-channel

	// if err := server.Shutdown(context.Background()); err != nil {
	// 	logrus.Panicf("Invalid service stopping: %s", err)
	// }

	// go bot.NewBot(os.Getenv(token), nil)

	select {}
}
