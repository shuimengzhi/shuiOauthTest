package main

import (
	"shuiOauth/router"
)

func main() {

	r := router.NewRouter()

	r.Run(":80")
}
