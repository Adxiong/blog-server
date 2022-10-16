/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-16 23:26:14
 */
package bootstrap

import (
	"github.com/gin-gonic/gin"
)

func Server(addr string) (*gin.Engine, error) {

	g := gin.Default()

	g.Routes()

	return g, nil
}
