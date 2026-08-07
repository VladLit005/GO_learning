package main

import (
	"bufio"
	"coffee/internal/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stock = domain.NewStock()

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		switch args[0] {

		case "help":
			{
				fmt.Println("справка")
				os.Exit(0)
			}

		case "menu":
			{
				for name, recipe := range domain.DefaultRecipes() {
					fmt.Printf("%s %d\n", name, recipe.Price)
				}
			}

		case "stock":
			{
				if len(args) < 2 {
					fmt.Println("usage")
					os.Exit(2)
				}
				switch args[1] {

				case "get":
					{
						for k, v := range stock.GetAll() {
							fmt.Printf("%s = %d\n", k, v)
						}
					}

				case "add":
					{
						if len(args) < 3 {
							fmt.Println("usage")
							os.Exit(2)
						}

						ingredient := args[2]
						count, err := strconv.Atoi(args[3])

						if err != nil {
							fmt.Println(domain.ErrInvalidParameters)
							os.Exit(1)
						}

						if err = stock.Add(ingredient, count); err != nil {
							fmt.Println(err)
							os.Exit(1)
						} else {
							fmt.Println("ok")
						}
					}

				case "set":
					{
						if len(args) < 3 {
							fmt.Println("usage")
							os.Exit(2)
						}

						ingredient := args[2]
						count, err := strconv.Atoi(args[3])

						if err != nil {
							fmt.Println(domain.ErrInvalidParameters)
							os.Exit(1)
						}

						if err = stock.Set(ingredient, count); err != nil {
							fmt.Println(err)
							os.Exit(1)
						} else {
							fmt.Println("ok")
						}
					}
				}
			}
		}
	}
}
