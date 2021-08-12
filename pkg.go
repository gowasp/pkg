package pkg

import (
	"context"
	"sync"
)

type SubFunc func(context.Context, []byte) error

type Subscribe struct {
	subMap sync.Map
}

func (s *Subscribe) Subscribe(topicID string, f SubFunc) {
	s.subMap.Store(topicID, f)
}

func (s *Subscribe) Get(topic string) SubFunc {
	if v, ok := s.subMap.Load(topic); ok {
		return v.(SubFunc)
	}
	return nil
}

type Private struct {
	subMap sync.Map
}

func (ps *Private) Subscribe(topicID byte, f SubFunc) {
	ps.subMap.Store(topicID, f)
}

func (ps *Private) Get(topicID byte) SubFunc {
	if v, ok := ps.subMap.Load(topicID); ok {
		return v.(SubFunc)
	}
	return nil
}
