package txtproc

import (
	"fmt"
)

var (
	ErrEmptyText = fmt.Errorf("empty text")

	ErrMappedStringsNil = fmt.Errorf("mapped strings is nil")
	ErrMinReplacerScore = fmt.Errorf("minimum replacer score must between 0 to 1")
)
