/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 00:01:59
 */
package bootstrap

import (
	"blogserver/httpserver"
	"context"

	"github.com/gin-gonic/gin"
)

func Server(ctx context.Context, addr string) (*gin.Engine, error) {

	g := gin.Default()

	httpserver.RegisterController(ctx, g)

	return g, nil
}
