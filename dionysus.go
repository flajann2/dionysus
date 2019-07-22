package main

import (
	"fmt"
	"os"
	
	"github.com/karrick/godirwalk"
	"github.com/flajann2/dionysus/minetype"
)

func main() {
	dirname := "/grohnde/torrent"
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if de.IsRegular(){
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
