package pkg

import (
	"sync"
)

type SubFunc func([]byte) error

type Subscribe struct {
	subMap sync.Map
}

func (s *Subscribe) Subscribe(topic string, f SubFunc) {
	s.subMap.Store(topic, f)
}

func (s *Subscribe) Get(topic string) SubFunc {
	if v, ok := s.subMap.Load(topic); ok {
		return v.(SubFunc)
	}
	return nil
}

func (s *Subscribe) GetTopics() []string {
	strs := make([]string, 0)
	s.subMap.Range(func(key, value interface{}) bool {
		strs = append(strs, key.(string))
		return true
	})
	return strs
}
