package main

import (
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

func BenchmarkHello(b *testing.B) {
	var name = "zhang san"
	greeting, err := hello(name)
	if err != nil {
		b.Errorf("The error is nil, but it should not be. (name=%q)", name)
	}
	b.Logf(greeting)

}
