package main

import "os"

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 3 {
		os.Stderr.Write([]byte("Usage           : ./lev <input> <input>\nValid Example   : ./lev xxx xxy; echo \"$?\"\n"))
		return
	}

	lev := Distance(os.Args[1], os.Args[2])

	os.Stdout.Write([]byte(string(lev + 48)))
	os.Exit(0)
}
