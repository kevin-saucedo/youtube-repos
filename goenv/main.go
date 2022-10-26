package main

import (
	"fmt"
	"goenv/envs"
)

func main() {
	v := envs.Get("DB_PASSWORDDDDDDD", "default-password")
	fmt.Println(v)
}
