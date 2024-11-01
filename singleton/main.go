package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type weekday string

var (
	test weekday = "monday"
)

var lock = sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating singleton", rand.Intn(100))
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

type DriverPg struct {
	conn string
}

var (
	once     sync.Once
	instance *DriverPg
)

func Connect() *DriverPg {
	once.Do(func() {
		fmt.Println("here once")
		instance = &DriverPg{conn: "PostGresDriver" + strconv.Itoa(rand.Intn(1000))}
	})
	return instance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// delayed call to connect
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	for i := 0; i < 100; i++ {
		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", Connect().conn)
		}(i)
	}

	fmt.Scanln()
}
