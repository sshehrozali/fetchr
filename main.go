package main

import (
	"fetchr/cmd"
)

func main() {
	runErr := cmd.Run()

	if runErr != nil {
		println(runErr.Error())
	}
}
