package minetype

import (
	"testing"
	"fmt"

	"github.com/h2non/filetype/types"

)

func TestMinetype(t *testing.T) {
	tmime := types.Type{
		MIME: types.MIME{
			Type:"image",
			Subtype:"jpeg",
			Value:"image/jpeg"},
		Extension:"jpg"}
	
	kind := Minetype("../test_media/eagle_nebula.jpg")
	fmt.Printf("TEST: %s\n", kind)
	if kind != tmime {
		t.Fatalf("invalid subtype %s", kind)
	}
}
