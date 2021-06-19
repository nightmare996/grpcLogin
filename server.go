package main

import (
	"log"

	"github.com/nightmare996/grpcLogin/config"
	"github.com/nightmare996/grpcLogin/model"
)

func main() {
	initServer()

}

func initServer() {
	err := config.ConfigInit("./conf/datasourse.json")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(config.DataSourseConfig)
	err = model.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("connect database success!")

	defer model.Close()
}
