package pkg

import "sync"

type Private struct {
	subMap sync.Map
}

type pvtSubFunc func(int, []byte)

func (ps *Private) Subscribe(topicID byte, f pvtSubFunc) {
	ps.subMap.Store(topicID, f)
}

func (ps *Private) Get(topicID byte) pvtSubFunc {
	if v, ok := ps.subMap.Load(topicID); ok {
		return v.(pvtSubFunc)
	}
	return nil
}
