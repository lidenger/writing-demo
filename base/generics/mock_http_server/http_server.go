package mock_http_server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	api(g)
	httpPort := 6666
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", httpPort),
		Handler: g,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err, "http服务异常")
		}
	}()
	log.Printf("Http服务已启动,端口:%d", httpPort)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func R(ctx *gin.Context, data any, err error) {
	code := 0
	msg := "success"
	if err != nil {
		code = -1
		msg = err.Error()
	}
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func api(g *gin.Engine) {
	g.GET("/", func(ctx *gin.Context) {
		R(ctx, "ok", nil)
	})

	v1 := g.Group("/api/v1")

	account := v1.Group("account")
	account.GET("", func(ctx *gin.Context) {
		R(ctx, GetMockAccountRespData(), nil)
	})

	order := v1.Group("order")
	order.GET("", func(ctx *gin.Context) {
		R(ctx, GetMockOrderRespData(), nil)
	})

	goods := v1.Group("goods")
	goods.GET("", func(ctx *gin.Context) {
		R(ctx, GetMockGoodsRespData(), nil)
	})

}
