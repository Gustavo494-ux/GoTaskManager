package main

import (
	"GoTaskManager/internal/app/inicializar"
	"runtime"
)

func init() {
	inicializar.Inicializar()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

}
