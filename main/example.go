package main

import (
	"RxAny/main/reative"
	"encoding/json"
	"log"
)

func main() {

	task := reative.Create()

	var a = 10
	var d = 10
	var b = 40
	var c int

	task.AddTask("getA", func() {
		a *= b
		by, _ := json.Marshal(a)
		task.AddValue("a", by)
	}).AddTask("getB", func() {
		b = d + b
		by1, _ := json.Marshal(b)
		task.AddValue("b", by1)
	}).AddTask("getC", func() {
		c = 10
		var temp int
		if c <= 10 {
			json.Unmarshal(task.GetValue("a"), &temp)
		} else {
			json.Unmarshal(task.GetValue("b"), &temp)
		}
		c *= temp
		log.Print(c)
	})

	task.Start()

	for {

	}

}
