package main

import "server/internal/bootstrap"

func main() {
	r := bootstrap.InitGin()
	r.Run(":8080")
}
