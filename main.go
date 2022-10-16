/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:16:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-16 23:26:02
 */
package main

import (
	"blogserver/bootstrap"
	"fmt"
)

func main() {
	ip := "127.0.0.1:8080"
	server, errServer := bootstrap.Server(ip)
	if errServer != nil {

	}

	if errRun := server.Run(ip); errRun != nil {
		fmt.Println("start listen server failed")
	}

}
