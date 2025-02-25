package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_fiber_crud?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Database Connected")
	DB = database
}

/*
dsn := "root:password@tcp(127.0.0.1:3306)/go_fiber_crud?... 	- এটি MySQL কানেকশন স্ট্রিং, যেখানে root:password পরিবর্তন করে আপনার MySQL ইউজারনেম ও পাসওয়ার্ড দিন।
gorm.Open(mysql.Open(dsn), &gorm.Config{}) 						- GORM দিয়ে MySQL-এর সাথে কানেক্ট করা হয়েছে।
DB = database													- ডাটাবেজ অবজেক্ট গ্লোবালি স্টোর করা হয়েছে।
*/
