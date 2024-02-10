package soda

import (
	"errors"
)

var ErrIteratorDone = errors.New("iterator done")
var ErrNotFound = errors.New("not found")
