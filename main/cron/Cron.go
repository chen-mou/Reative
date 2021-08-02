package cron

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	name     string
	task     func()
	index    int
	taskTime []string
}

var tasks []Task

func isEqual(index int, value string) bool {
	if value == "*" {
		return true
	}
	atoi, _ := strconv.Atoi(value)
	nowTime := time.Now()
	switch index {
	case 0:
		return nowTime.Second() == atoi
	case 1:
		return nowTime.Minute() == atoi
	case 2:
		return nowTime.Hour() == atoi
	case 3:
		return nowTime.Day() == atoi
	case 4:
		return nowTime.Month().String() == value
	case 5:
		return nowTime.Year() == atoi
	}
	return false
}

func isNow(tim []string) bool {
	for index, value := range tim {
		if !isEqual(index, value) {
			return false
		}
	}
	return true
}

func AddTask(name string, times string, task func()) {
	temp := strings.Split(times, " ")
	var Index int
	for index := range temp {
		if temp[index] != "*" {
			Index = index - 1
			break
		}
		temp[index] = "0"
	}
	value := Task{
		name:     name,
		task:     task,
		taskTime: temp,
		index:    Index,
	}
	tasks = append(tasks, value)
}

func thread() {
	MinTime := time.Second
	minIndex := 7
	for _, index := range tasks {
		if index.index < minIndex {
			minIndex = index.index
		}
	}
	for index := 1; index <= minIndex; index++ {
		switch index {
		case 1:
		case 2:
			MinTime *= 60
			break
		case 3:
			MinTime *= 24
			break
		case 4:
			MinTime *= 30
			break
		}
	}
	for {
		for _, index := range tasks {
			if isNow(index.taskTime) {
				fmt.Print("[CRON-INFO] Task:" + index.name + " running")
				go index.task()
			}
			time.Sleep(time.Duration(1) * MinTime)
		}
	}
}

func Start() {
	go thread()
	fmt.Print("[CRON-INFO] Cron is run\n")

}
