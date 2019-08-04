package main

import (
	"os"
	"fmt"
	"flag"
	"log"
	
	"github.com/flajann2/dionysus/minetype"
	"github.com/karrick/godirwalk"
	"github.com/h2non/filetype/types"
	bolt "go.etcd.io/bbolt"
)

type metadata struct {
	path string
	minetype types.Type
	modetype os.FileMode
}

const dbname = "dionysus.db"
const media = "media"

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

func ensure(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(media))
		if err != nil{
			log.Fatal(err)
		}
		return nil
	})
}

func main() {
	flag.Parse()	
	paths := flag.Args()

	db, err := bolt.Open(dbname, 0600, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		ensure(db)
		ch := make(chan metadata, 100000)
		for _, path := range paths {
			fmt.Printf("scanning path: %s\n", path)
			go scanDirectory(ch, path)
		}
		for	m := range ch {
			fmt.Printf("ch: %s %s %s\n", m.modetype, m.path, m.minetype)
		}
	}
	db.Close()
}
