package main

import (
	"aoc2025/info"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func begin() (*info.Application, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("could not read env file %w", err)
	}
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return nil, fmt.Errorf("no session cookie supplied")
	}
	return &info.Application{
		Session: session,
	}, nil
}

func main() {
	app, err := begin()
	if err != nil {
		panic(err)
	}
	data, err := app.GetData("https://adventofcode.com/2024/day/2/input")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println("booting up advent of code application")
}
