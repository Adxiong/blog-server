/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-22 16:00:32
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-22 16:00:37
 */
package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var datetimePrecision = 2

var GlobalDb *gorm.DB

func NewDb() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                          // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                         // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision,                                                           // default datetime precision
		DontSupportRenameIndex:    true,                                                                         // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                         // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                        // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("GET_DB_Conn_failed!")
	}
	return db
}
