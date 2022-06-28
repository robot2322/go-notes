package factory

import (
	"fmt"
	"testing"
)

func TestAPI1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	fmt.Println(s)
}

func TestAPI2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	fmt.Println(s)
}
