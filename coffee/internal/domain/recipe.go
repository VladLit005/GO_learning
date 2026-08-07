package domain

type Recipe struct {
	Name        string
	Price       int
	Ingredients map[string]int
	Steps       []string
}

func DefaultRecipes() map[string]Recipe {
	return map[string]Recipe{
		"espresso": {
			Name:  "espresso",
			Price: 150,
			Ingredients: map[string]int{
				"beans": 8,
				"water": 30,
			},
			Steps: []string{
				"grind",
				"tamp",
				"brew 25s",
			},
		},
		"americano": {
			Name:  "americano",
			Price: 180,
			Ingredients: map[string]int{
				"beans": 8,
				"water": 150,
			},
			Steps: []string{
				"grind",
				"brew 25s",
				"add water",
			},
		},
		"latte": {
			Name:  "latte",
			Price: 220,
			Ingredients: map[string]int{
				"beans": 8,
				"water": 30,
				"milk":  150,
			},
			Steps: []string{
				"grind",
				"brew 25s",
				"steam milk",
				"mix",
			},
		},
	}
}
