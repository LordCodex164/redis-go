package store

import "errors"

var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyKey = errors.New("No Keys are found")
var ErrLimitExceeded = errors.New("Max limit exceeded")

type IStore interface {
	Set(key, val string) (error)
	Get(key string)		 (string, error)
	//Keys()				 []string
	Pop(key string)		 (string, bool)
	Len()				 int
}