package minetype

import (
	"os"
	
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Minetype (path string) types.Type {
	buf := make([]byte, 512)
	f, err := os.Open(path)
	check(err)
	
	_, err = f.Read(buf)
	check(err)

	kind, _ := filetype.Match(buf)
	f.Close()
	
	return kind
}
