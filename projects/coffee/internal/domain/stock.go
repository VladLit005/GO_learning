package domain

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
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

func (s Stock) IsEnough(ingredients map[string]int) bool {

	for name, need := range ingredients {
		if have := s.items[name]; have < need {
			return false
		}
	}
	return true
}

func (s *Stock) Deduct(ingredients map[string]int) {
	for name, count := range ingredients {
		s.items[name] -= count
	}
}

func (s Stock) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range s.items {

		fmt.Fprintf(file, "%s %d\n", k, v)
	}

	return nil
}

func Load(path string) (*Stock, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	stock := NewStock()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		if err := stock.Set(name, count); err != nil {
			return nil, err
		}
	}

	return stock, nil
}
