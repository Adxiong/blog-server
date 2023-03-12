/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:13
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-02-27 00:19:24
 */
package bootstrap

import (
	"blogserver/model/dao/db"
	"io"
	"log"
	"os"

	"blogserver/library/resource"
)

func Mustinit() {
	initLog()
	initMysql()
}

func initConfig() {

}

func initMysql() {
	db.GlobalDb = db.NewDb()

	db.GlobalDb.AutoMigrate(db.NewUser())
	db.GlobalDb.AutoMigrate(db.NewArticle())
}

func initLog() {
	f, err := os.OpenFile("log/blog.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panic(err)
		return
	}
	defer f.Close()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	resource.ServiceLogger = log.New(multiWriter, "Info", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

}
