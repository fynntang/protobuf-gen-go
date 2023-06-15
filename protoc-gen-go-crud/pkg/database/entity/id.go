package entity

import (
	"strconv"
)

type ID int64

func (id ID) Int64() int64 {
	return int64(id)
}

func (id ID) String() string {
	return strconv.FormatInt(id.Int64(), 10)
}

func (id ID) FromInt64(i int64) ID {
	return IDFromInt64(i)
}

func (id ID) FromString(s string) ID {
	return IDFromString(s)
}

func IDFromString(s string) ID {
	id, _ := IDFromStringError(s)
	return id
}
func IDFromStringError(s string) (ID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return ID(i), err
}
func IDFromInt64(i int64) ID {
	return ID(i)
}
