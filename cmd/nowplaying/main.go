package main

import (
	"os"
	"log"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying"
	flags "github.com/jessevdk/go-flags"
	"fmt"
)

const version = "0.0.1"

var option nowplaying.Option

func main() {
	parser := flags.NewParser(&option, flags.Default)
	parser.Name = "nowplaying"
	_, err := parser.Parse()

	if err != nil {
		log.Fatalf("option not loaded. %v", err)
	}

	switch {
	case option.Version:
		fmt.Printf("%s \n", version)
	default:
		run(option.GetFilename())
	}
}


func run(filename string) {
	config, err := config.NewConfig(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	nowplaying.Run(config)
}