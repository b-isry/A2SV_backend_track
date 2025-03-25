package main

import (
	"github.com/Bisrat/task-manager/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
