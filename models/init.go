/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-06-07 11:13:17
 * @LastEditors: Y95201
 * @LastEditTime: 2022-06-07 14:24:31
 */
package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/think?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(false)      //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(10)    //最大空闲连接
	db.DB().SetMaxOpenConns(100)   //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)
	return db
}
