package main

import (
	"testing"
	"fmt"
	"time"
	"sync"
	"runtime"
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

type Person struct {
	Name string
}

var count int = 0

func TestIf(t *testing.T) {
	lock := &sync.Mutex{}
	cond := sync.NewCond(lock)

	for i := 0; i < 10; i++ {
		go condCopy(cond, i)
	}
	time.Sleep(11e9)
	cond.Signal()
	time.Sleep(1e9)
	cond.Signal()
	time.Sleep(1e9)
	cond.Signal()
	_, file, line, _ := runtime.Caller(0)
	_, file1, line1, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	fmt.Println(file, "	", line)
	fmt.Println(file2, "	", line2)
	fmt.Println(file1, "	", line1)
	time.Sleep(30e9)
}
func condCopy(cond *sync.Cond, id int) {
	duration := time.Duration(1e9 * id)
	time.Sleep(duration)

	cond.L.Lock()
	count++
	for count < 10 {
		fmt.Println("id:		", id, "count:		", count)
		cond.Wait()
	}
	fmt.Println("count 10	", id)
	cond.L.Unlock()
}
