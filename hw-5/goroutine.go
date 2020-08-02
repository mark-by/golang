package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func GoRun(tasks []func() error, maxGoroutines int, maxErrors int) {
	var start = make(chan int)
	var chError = make(chan error, len(tasks))
	lastCounterValue := 0
	for i := 0; i < len(tasks); i++ {
		if i < maxGoroutines {
			go func(i int) {
				<- start
				chError <- tasks[i]()
			}(i)
		} else {
			lastCounterValue = i
			break
		}
	}
	if lastCounterValue < len(tasks) {
		go func() {
			<- start
			for i := lastCounterValue; i < len(tasks); i++ {
				chError <- tasks[i]()
			}
		}()
	}
	close(start)
	taskCounter := 0
	errorCounter := 0
	for {
		err, _ := <- chError
		taskCounter++
		if err != nil {
			fmt.Errorf("ERROR: %v", err)
			errorCounter++
			if errorCounter > maxErrors {
				println("Много задач провалилось")
				os.Exit(1)
			}
		}
		if taskCounter == len(tasks) {
			println("Успешно")
			break
		}
	}
}

func main() {
	functions := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			println("a")
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			println("b")
			return errors.New("b fail")
		},
		func() error {
			time.Sleep(5 * time.Millisecond)
			println("c")
			return nil
		},
		func() error {
			time.Sleep(3 * time.Second)
			println("d")
			return errors.New("d fail")
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("e")
			return nil
		},
		func() error {
			time.Sleep(3 * time.Second)
			println("f")
			return errors.New("f fail")
		},
	}
	GoRun(functions, 4, 3)
}