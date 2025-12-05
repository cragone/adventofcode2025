package main

import (
	"aoc2025/env"
	"aoc2025/info"
	"aoc2025/twentyfive"
	"fmt"
	"log"
	"os"
)

func begin() (*info.Application, error) {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, fmt.Errorf("no session cookie supplied")
	}
	return &info.Application{
		Session: session,
	}, nil
}

func main() {
	fmt.Println("Booting up advent of code application...")
	if err := env.ReadENV(".env"); err != nil {
		log.Fatal(err)
	}
	app, err := begin()
	if err != nil {
		log.Fatal(err)
	}
	data, err := app.GetData("https://adventofcode.com/2025/day/2/input")
	if err != nil {
		log.Fatal(err)
	}

	// test := "1000-2000,500-1500"
	count := twentyfive.InvalidIdCount(data)

	fmt.Println(count)

}
