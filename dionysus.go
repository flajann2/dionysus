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
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			ftype := minetype.Minetype(osPathname)
			fmt.Printf("%s %s %s\n",
				de.ModeType(), osPathname, ftype)
			
			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
