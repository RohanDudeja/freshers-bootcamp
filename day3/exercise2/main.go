package main

import (
	"fmt"
	Config2 "freshers-bootcamp/day3/exercise2/Config"
	Models2 "freshers-bootcamp/day3/exercise2/Models"
	Routes2 "freshers-bootcamp/day3/exercise2/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config2.DB, err = gorm.Open("mysql", Config2.DbURL(Config2.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config2.DB.Close()
	Config2.DB.AutoMigrate(&Models2.Student{})
	r := Routes2.SetupRouter()
	//running
	r.Run()
}