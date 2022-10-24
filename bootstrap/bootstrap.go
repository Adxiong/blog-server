/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-16 23:18:13
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 22:48:42
 */
package bootstrap

import (
	"blogserver/model/dao/db"
	"io"
	"log"
	"os"
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
	f, err := os.OpenFile("log/blog.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
