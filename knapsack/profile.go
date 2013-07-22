package main

import (
	"runtime/pprof"
	"os"
)

func StartCPUProfile(filename string) {
    f, _ := os.Create(filename)
    pprof.StartCPUProfile(f)
}

func StopCPUProfile() {
    pprof.StopCPUProfile()
}