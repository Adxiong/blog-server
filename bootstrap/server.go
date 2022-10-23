/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 23:59:08
 */
package bootstrap

import (
	"blogserver/httpserver"

	"github.com/gin-gonic/gin"
)

func Server(ctx content.Content addr string) (*gin.Engine, error) {

	g := gin.Default()

	httpserver.RegisterController(ctx, g)

	return g, nil
}
