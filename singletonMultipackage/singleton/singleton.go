package singleton

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Singleton struct {
	Name string
}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "random name" + strconv.Itoa(rand.Intn(100))}
		fmt.Println("singleton created")
	})
	return instance
}
