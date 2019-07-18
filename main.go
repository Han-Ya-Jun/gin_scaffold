package main

import (
	"fmt"
	"github.com/Han-Ya-Jun/gin_scaffold/app"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/setting"
	mlog "github.com/e421083458/golang_common/log"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	apiServer := app.NewApiServer()
	apiServer.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		mlog.Info(fmt.Sprintf("gin_scaffold service get a signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			mlog.Info("gin_scaffold exit")
			apiServer.Stop()
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
