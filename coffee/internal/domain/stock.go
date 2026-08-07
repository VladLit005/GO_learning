package domain

import (
	"maps"
)

type Stock struct {
	items map[string]int
}

func NewStock() *Stock {
	return &Stock{
		items: map[string]int{
			"beans": 0,
			"water": 0,
			"milk":  0,
			"sugar": 0,
		},
	}
}

func (s Stock) GetAll() map[string]int {
	result := make(map[string]int, len(s.items))
	maps.Copy(result, s.items)
	return result
}

func (s *Stock) Add(name string, count int) error {
	if _, ok := s.items[name]; !ok {
		return ErrInvalidParameters
	}

	if count < 1 {
		return ErrInvalidParameters
	}

	s.items[name] += count
	return nil
}

func (s *Stock) Set(name string, count int) error {
	if _, ok := s.items[name]; !ok {
		return ErrInvalidParameters
	}

	if count < 0 {
		return ErrInvalidParameters
	}

	s.items[name] = count
	return nil
}
