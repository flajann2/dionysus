package main

import (
	"os"
	"fmt"
	"flag"
	
	"github.com/flajann2/dionysus/minetype"
	"github.com/karrick/godirwalk"
	"github.com/h2non/filetype/types"
)

type metadata struct {
	path string
	minetype types.Type
	modetype os.FileMode
}


func scanDirectory(ch chan metadata, dirname string) {

	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if de.IsRegular() {
				go func() {
					ch <- metadata{path,
						minetype.Minetype(path),
						de.ModeType()}					
				}()
			}
			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()	
	paths := flag.Args()

	ch := make(chan metadata, 100000)
	for _, path := range paths {
		fmt.Printf("scanning path: %s\n", path)
		go scanDirectory(ch, path)
	}
	for m := range ch {
		fmt.Printf("ch: %s %s %s\n", m.modetype, m.path, m.minetype)
	}
}
