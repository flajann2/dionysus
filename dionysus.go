package main

import (
	"fmt"
	"os"
	
	"github.com/karrick/godirwalk"
)

func main() {
	dirname := "/grohnde/torrent"
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
