package filter

import (
	"fmt"
)

type Limits struct {
	Offset int
	Limit  int
}

func (l Limits) String() string {
	return fmt.Sprintf("%d_%d", l.Limit, l.Offset)
}
