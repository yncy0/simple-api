package main

import (
	"github.com/yncy0/simple-api/internal/message"
)

func main() {
	r := message.SetupRouter()	
	r.Run(":8080")
}
