package main

import (
	"github.com/ykubernetes/GUUID/router"
)

func main() {
	r := router.Setup()
	r.Run("0.0.0.0:9000")
}
