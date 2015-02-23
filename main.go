package main

import (
	"log"
)

func init() {
	LoadAdventure()
}

func main() {
	log.Println("Let's do this!")
	Adventure()
	log.Println("The adventure is over...")
}
