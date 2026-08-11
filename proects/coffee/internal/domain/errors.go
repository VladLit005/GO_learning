package domain

import "errors"

var (
	ErrDrinkNotFound     = errors.New("Напиток не найден")
	ErrNotEnoughMoney    = errors.New("Недостаточно денег")
	ErrNotEnoughStock    = errors.New("Недостаточно ингредиентов")
	ErrInvalidParameters = errors.New("Некорректные параметры")
)
