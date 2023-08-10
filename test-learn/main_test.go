package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var name = "meiken gu"
	greeting, err := hello(name)
	if err != nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)", name)
	}
	t.Logf(greeting)
}

//
//func TestFail(t *testing.T) {
//	t.Log("Test Fail Func ")
//	t.FailNow()
//}

func BenchmarkHello(b *testing.B) {
	var name = "zhang san"
	greeting, err := hello(name)
	if err != nil {
		b.Errorf("The error is nil, but it should not be. (name=%q)", name)
	}
	b.Logf(greeting)

}

func ExampleHello() {

	var name = "example test"
	greeting, err := hello(name)
	if err != nil {
		fmt.Println("The error is nil, but is should not bt. (name)", name)
	}
	fmt.Println(greeting)
}
