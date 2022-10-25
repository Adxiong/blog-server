/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-25 23:38:47
 */
package bootstrap

import (
	"blogserver/httpserver"
	"blogserver/library/middlewares"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, addr string) (*gin.Engine, error) {

	Mustinit()

	server, errServer := Server(ctx, addr)

	if errServer != nil {
		log.Println(errServer)
		return nil, errServer
	}

	return server, nil
}

func Server(ctx context.Context, addr string) (*gin.Engine, error) {

	g := gin.Default()
	g.Use(middlewares.Cors())

	httpserver.RegisterController(ctx, g)

	return g, nil
}
