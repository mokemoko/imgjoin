package main

import (
	"github.com/gin-gonic/gin"
	"imgjoin/controllers"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}

func main() {
	r := gin.Default()

	controllers.RegisterJoin(r.Group("/join"))

	r.Run()
}
