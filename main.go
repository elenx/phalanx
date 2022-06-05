package main

import (
	"log"
	"phalanx/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%s", err)
	}
}
