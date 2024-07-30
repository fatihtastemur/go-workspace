package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start Validation : ", time.Now())
	list := []int{13, 28, 39, 41, 55, 66, 78, 83, 94, 110, 121, 194}

	listChannel := make(chan *int, len(list))

	var successList []int

	wg := sync.WaitGroup{}
	for _, item := range list {
		wg.Add(1)
		go validateNumber(&wg, listChannel, item)
	}

	wg.Wait()

	for done := false; !done; {
		select {
		case successItem := <-listChannel:
			successList = append(successList, *successItem)
		default:
			done = true
		}
	}

	close(listChannel)

	fmt.Println("VALIDATED : ", successList)
	fmt.Println("End Validation : ", time.Now())
	process(successList)
}

func validateNumber(wg *sync.WaitGroup, listChannel chan *int, item int) {
	defer wg.Done()

	if item < 100 {
		listChannel <- &item
	}
}

func process(list []int) {
	fmt.Println("Start Process : ", time.Now())
	processChannel := make(chan *int, len(list))

	var processList []int

	wg := sync.WaitGroup{}
	for _, item := range list {
		wg.Add(1)
		go processNumber(&wg, processChannel, item)
	}

	wg.Wait()

	for done := false; !done; {
		select {
		case processItem := <-processChannel:
			processList = append(processList, *processItem)
		default:
			done = true
		}
	}

	close(processChannel)

	fmt.Println("PROCESSED : ", processList)
	fmt.Println("End Process : ", time.Now())
}

func processNumber(wg *sync.WaitGroup, processChannel chan *int, successItem int) {
	defer wg.Done()

	successItem = successItem + 1
	processChannel <- &successItem
}
