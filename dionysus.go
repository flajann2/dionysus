package main

import (
	"os"
	"fmt"
	
	"github.com/flajann2/dionysus/minetype"
	"github.com/karrick/godirwalk"
)

func main() {
	dirname := "/grohnde/torrent"
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if de.IsRegular() {
				ftype := minetype.Minetype(path)
				fmt.Printf("%s %s %s\n",
					de.ModeType(), path, ftype)
			}

			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
