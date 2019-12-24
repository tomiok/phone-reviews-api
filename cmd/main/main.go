package main

import (
	"github.com/tomiok/course-phones-review/internal/database"
	"github.com/tomiok/course-phones-review/internal/logs"
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSqlClient("root:root@/phones-review")

}
