/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 22:45:42
 */
package bootstrap

import (
	"blogserver/httpserver"
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

	httpserver.RegisterController(ctx, g)

	return g, nil
}
