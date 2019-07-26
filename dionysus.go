package main

import (
	"os"
	"fmt"
	
	"github.com/flajann2/dionysus/minetype"
	"github.com/karrick/godirwalk"
	"github.com/h2non/filetype/types"
)

type metadata struct {
	path string
	minetype types.Type
	modetype os.FileMode
}

func main() {
	dirname := "/grohnde/torrent"
	ch := make(chan metadata, 100000)
	
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if de.IsRegular() {
				m := metadata{path,
					minetype.Minetype(path),
					de.ModeType()}
				
				fmt.Printf("%s %s %s\n",
					m.modetype, m.path, m.minetype)
			}
			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
