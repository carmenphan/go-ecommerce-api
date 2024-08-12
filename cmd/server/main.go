package main

import (
	"github.com/carmenphan/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
