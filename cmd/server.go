package main

import (
	"fmt"
	"github.com/shevchenkobn/blog-backend/internal/services/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
}
