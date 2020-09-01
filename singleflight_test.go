package singleflight

import (
	"fmt"
	"testing"
	"golang.org/x/sync/singleflight"
)

var g = Group{}

func TestDoWithFnPanic(t *testing.T){
	middle()
	if len(g.m) != 0 {
		t.FailNow()
	}
}

var g2 = singleflight.Group{}

func TestDoWithFnPanic2(t *testing.T){
	middle2()
	fmt.Println(g2) // {{0 0} map[1:0xc0000662a0]}
}

func middle() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	g.Do("1", willPanic)
}

func middle2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	g2.Do("1", willPanic)
}

func willPanic() (interface{}, error) {
	panic("panic!")
}