package domain

type Storage interface {
	GetAll() map[string]int
	Add(name string, count int) error
	Set(name string, count int) error
	IsEnough(ingredients map[string]int) bool
	Deduct(ingredients map[string]int)
}
