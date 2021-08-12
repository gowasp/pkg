package pkg

import (
	"context"
	"sync"
)

type Private struct {
	subMap sync.Map
}

type PvtSubFunc func(context.Context, []byte) error

func (ps *Private) Subscribe(topicID byte, f PvtSubFunc) {
	ps.subMap.Store(topicID, f)
}

func (ps *Private) Get(topicID byte) PvtSubFunc {
	if v, ok := ps.subMap.Load(topicID); ok {
		return v.(PvtSubFunc)
	}
	return nil
}
