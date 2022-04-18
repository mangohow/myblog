package main

/*
* @Author: mgh
* @Date: 2022/2/28 18:33
* @Desc:
 */

import (
	"blog_web/db/dao"
	"blog_web/router"
	"blog_web/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	engine := gin.New()
	//gin.SetMode(gin.ReleaseMode)

	engine.Use(gin.LoggerWithWriter(utils.Logger().Out), gin.Recovery())
	engine.Static("/images", viper.GetString("server.imagePath") + "/images")

	router.Register(engine)

	addr := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
	defer dao.CloseMysqlConn()

	srv := http.Server{
		Addr:    addr,
		Handler: engine,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("Server closed!")
				return
			} else {
				panic(err)
			}
		}
	}()

	dealSignal(&srv)
}

func dealSignal(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-sigChan:
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			utils.Logger().Fatal("Server is going to shutdown...")
			c, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFunc()

			if err := server.Shutdown(c); err != nil {
				utils.Logger().Fatal("Stop server error:%v", err)
			}
			return
		}
	}

}
