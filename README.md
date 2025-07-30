# pointer
Generic pointer to value in Go

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
```
