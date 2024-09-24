package main

import (
	"filedownloader/cmd"
)

func main() {
	runErr := cmd.Run()

	if runErr != nil {
		println(runErr.Error())
	}
}
