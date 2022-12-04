package main

import "time"

func main() {
	dur, err := time.ParseDuration("10s")
	if err != nil {
		panic(err)
	}

	time.Sleep(dur)
}
