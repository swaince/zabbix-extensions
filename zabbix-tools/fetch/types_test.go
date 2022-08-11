package fetch

import (
	"fmt"
	"testing"
)

func TestFieldObject_ParseKey(t *testing.T) {
	f := FieldObject{
		//RawKey: "delay（必需）",
		RawKey: "itemid",
	}

	fmt.Printf("%s\n", ParseKey(f.RawKey))
}
