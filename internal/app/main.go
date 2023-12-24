package main

import (
	"GoTaskManager/internal/app/inicializar"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	inicializar.Inicializar()
}
