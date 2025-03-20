package main

import (
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		cmd := exec.Command("go", "run", "service1/main.go")
		cmd.Run()
	}()

	go func() {
		defer wg.Done()
		cmd := exec.Command("go", "run", "service2/main.go")
		cmd.Run()
	}()

	wg.Wait()
}
