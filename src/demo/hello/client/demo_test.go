package main

import (
	"testing"
	"fmt"
	"time"
	list2 "container/list"
	"container/heap"
)

func TestDefer(t *testing.T) {
	start := time.Now()
	end := time.Since(start)
	fmt.Println(end)
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

type Runnable interface {
	run()
}
type Thread struct {
}

func (th Thread) run() {
	fmt.Println("running...")
}
func TestIf(t *testing.T) {
	var s int =9

}
