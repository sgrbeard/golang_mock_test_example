package service

import "sync"

type exampleService struct{}

var singleton ExampleService

var once sync.Once

func GetExampleService() ExampleService {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &exampleService{}
	})
	return singleton
}

func SetExampleService(service ExampleService) ExampleService {
	original := singleton
	singleton = service
	return original
}

type ExampleService interface {
	Get(no int) int
}

func (c *exampleService) Get(no int) int {
	if no == 1 {
		return 1
	}
	return 2
}
