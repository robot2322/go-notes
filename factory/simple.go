package factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hi{}
	}else if t == 2 {
		return &hello{}
	}
	return nil
}

type hi struct {}

func (*hi) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type hello struct {}

func (*hello) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}