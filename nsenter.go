package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gokrazy/gokrazy"
)

func main() {
	gokrazy.DontStartOnBoot()

	f, err := os.OpenFile("/tmp/nsenter.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%v: %v\n", time.Now(), os.Args)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	os.Exit(61) // arbitrary, intentionally non-standard
}
