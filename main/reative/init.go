package reative

import (
	"RxAny/main/config"
	"RxAny/main/cron"
)

type Task struct {
	tasks map[string]func()

	resultChan map[string]chan []byte

	isClose bool

	isDelete bool
}

var nowThread = 0

func (t *Task) AddValue(name string, value []byte) *Task {
	if t.resultChan[name] == nil {
		t.resultChan[name] = make(chan []byte, 100)
	}
	t.resultChan[name] <- value
	return t
}

func (t Task) GetValue(name string) []byte {
	var result []byte
	if t.resultChan[name] == nil {
		t.resultChan[name] = make(chan []byte, config.DefaultBuffSize)
	}
	result = <-t.resultChan[name]
	return result
}

func (t *Task) getTask() {
	for key, value := range t.tasks {
		delete(t.tasks, key)
		value()
		nowThread--
	}
}

func (t *Task) AddTask(task func(), name string) *Task {
	t.isClose = false
	if t.tasks[name] != nil {
		panic("该名字已被使用")
	}
	t.tasks[name] = task
	return t
}

func (t *Task) Wait() {
	t.isClose = true
}

func (t *Task) Close() {
	t.isDelete = true
}

func (t *Task) start() {
	for {
		if nowThread < config.MaxThread {
			if !t.isClose {
				go t.getTask()
			}
			nowThread++
		}
		if t.isDelete {
			break
		}
	}
}

func (t *Task) Start() {

	go t.start()

	cron.AddTask("close", "* 5 * * * *", func() {
		t.Wait()
	})
}
