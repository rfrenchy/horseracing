package main

import "testing"

func TestHelloWorld(t *testing.T) {
        name := "Hello World"
        if name != "Hello World" {
                t.Errorf("name = %s; want Hello World", name)
        }
}

func FuzzHorse(f *testing.F) {
        f.Fuzz(func(t *testing.T, _ []byte) {
                if 1 == 0 {
                        f.Errorf("the worlds ending")
                }
        })
}
