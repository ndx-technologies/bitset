package bitset

import "strconv"

type ErrInvalidDataLength struct {
	Length int
}

func (e ErrInvalidDataLength) Error() string { return "bad data length: " + strconv.Itoa(e.Length) }
