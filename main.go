package main

import (
	_ "config-server/internal/pkg/etcd3"
	"config-server/internal/router"
)

func main() {
	r := router.New()
	r.InitRoute()

	r.Run("127.0.0.1:8080")
}
