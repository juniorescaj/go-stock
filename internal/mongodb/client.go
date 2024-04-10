package mongodb

import "sync"

var lock = &sync.Mutex

type single struct {
}

var mongodbClient *single

func getInstance() *single
