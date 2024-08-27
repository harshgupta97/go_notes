package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	fmt.Println("Hello")
}
