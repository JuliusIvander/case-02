package main

import (
	"log"

	"github.com/JuliusIvander/case-02/route"
)

func main() {
	userRoute := route.NewUserRoute()
	log.Fatal(userRoute.Listen(":3306"))
}
