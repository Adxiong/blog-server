/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:16:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 00:00:51
 */
package main

import (
	"blogserver/bootstrap"
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	ip := "127.0.0.1:8080"
	server, errServer := bootstrap.Server(ctx, ip)
	if errServer != nil {

	}
	if errRun := server.Run(ip); errRun != nil {
		fmt.Println("start listen server failed")
	}

}
