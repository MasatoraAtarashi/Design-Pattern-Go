package main

import "main/pkg/si"

func main() {
	ch := make(chan interface{})
	go func() { ch <- si.GetSingletonInstance() }()
	go func() { ch <- si.GetSingletonInstance() }()
	go func() { ch <- si.GetSingletonInstance() }()
	<-ch
	<-ch
	<-ch
}
