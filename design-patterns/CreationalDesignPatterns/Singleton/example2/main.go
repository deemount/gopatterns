package plus

import (
	"fmt"
	"sync"
)

var once sync.Once

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = new(singleton) // or instance = &singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
