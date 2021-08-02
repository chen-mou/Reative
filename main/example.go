package main

import (
	"RxAny/main/reative"
	"log"
)

func main() {

	task := reative.Create()

	a := 10
	b := 40
	var c int

	task.AddTask("getA", func() {
		a *= b
		task.AddValue("a", []byte{byte(a)})
	}).AddTask("getB", func() {
		b *= a + b
		task.AddValue("b", []byte{byte(b)})
	}).AddTask("getC", func() {
		c = 100
		if c <= 10 {
			c *= int(task.GetValue("a")[0])
		} else {
			c *= int(task.GetValue("b")[0])
		}
		log.Print(c)
	})

	task.Start()

	for {

	}

}
