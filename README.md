# pointer
Generic pointer to arbitrary value in Go

Usage example:
```Go
import (
	"github.com/DmitriyVTitov/pointer"
)

type T struct {
	BoolPointer   *bool
	StringPointer *string
}

var V = T{
	BoolPointer:   pointer.To(true),
	StringPointer: pointer.To("string"),
}

var v any
var nilPointer = pointer.To(v)
```
