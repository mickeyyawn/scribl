package main

import "fmt"

func HandleError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))

	// TODO: add logging
	// TODO; return string of error ?

	//os.Exit(1)
}
