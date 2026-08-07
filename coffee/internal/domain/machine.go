package domain

type Machine struct {
	stock   Storage
	recipes map[string]Recipe
	orders  int
	revenue int
}

func NewMachine(s Storage) *Machine {
	return &Machine{
		recipes: DefaultRecipes(),
		stock:   s,
	}
}

func (m Machine) GetRecipes() map[string]Recipe {
	return m.recipes
}

func (m *Machine) Brew(drink string, payment int) ([]string, error) {
	recipe, ok := m.recipes[drink]
	if !ok {
		return nil, ErrDrinkNotFound
	}
	if payment < recipe.Price {
		return nil, ErrNotEnoughMoney
	}

	if !m.stock.IsEnough(recipe.Ingredients) {
		return nil, ErrNotEnoughStock
	}

	m.stock.Deduct(recipe.Ingredients)

	m.orders++
	m.revenue += recipe.Price

	return recipe.Steps, nil
}

func (m Machine) Stats() (int, int) {
	return m.orders, m.revenue
}
