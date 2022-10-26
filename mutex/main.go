package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.RWMutex
var usuarios = make(map[string]any)

func main() {
	go insertar("v1", "hola")
	go insertar("v2", 1)
	go insertar("v3", 2)
	go insertar("v4", 3)
	go leer()
	go insertar("v5", 4)
	go insertar("v6", 5)
	go insertar("v7", 6)
	time.Sleep(time.Second)
}
func insertar(key string, value any) {
	mutex.Lock()
	defer mutex.Unlock()
	usuarios[key] = value
}
func leer() {
	mutex.Lock()
	value, ok := usuarios["v1"]
	mutex.Unlock()
	if ok {
		fmt.Println(value)
	}
	fmt.Println(ok)
}
