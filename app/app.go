package app

import (
	"context"
	"fmt"
	mogo "github.com/Han-Ya-Jun/gin_scaffold/models/mongo"
	redis "github.com/Han-Ya-Jun/gin_scaffold/pkg/gredis"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/logging"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/setting"
	"github.com/Han-Ya-Jun/gin_scaffold/routers"
	"log"
	"net/http"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/06/06 15:13
* @Name:app
* @Function: app初始化
 */

type ApiServer struct {
	server *http.Server
}

func Init() {
	var err error
	err = setting.Setup()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	mogo.SetUp()
	err = redis.Setup()
	if err != nil {
		panic(err)
	}
	logging.Setup()
}

func NewApiServer() *ApiServer {
	Init()
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	log.Printf(fmt.Sprintf("start http server listening %s", endPoint))
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	return &ApiServer{
		server: server,
	}

}

func (s *ApiServer) Start() {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatal("server start failed")
		} else {
			log.Fatal(" server start success")
		}
	}()
}

func (s *ApiServer) Stop() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = s.server.Shutdown(ctx)
}
