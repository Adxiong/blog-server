/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:16:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 22:56:25
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

	ip := ":8080"
	server, errServer := bootstrap.Start(ctx, ip)
	if errServer != nil {
		fmt.Println("bootstrap server start failed")
	}
	if errRun := server.Run(ip); errRun != nil {
		fmt.Println("start listen server failed")
	}

}
