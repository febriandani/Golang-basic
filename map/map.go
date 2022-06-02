package main

import "fmt"

func main() {
	var ayam map[string]int
	ayam = map[string]int{}

	ayam["januari"] = 50
	ayam["februari"] = 40

	fmt.Println("januari",ayam["januari"])
	fmt.Println("mei",ayam["mei"])


	fmt.Println("Iterasi map menggunakan for - range")
	var chicken = map[string]int{
		"januari" : 10,
		"februari" : 50,
		"maret" : 30,
		"april" : 40,
	}

	for key, val := range chicken{
		fmt.Println(key," \t",val)
	}

	fmt.Println("\nMenghapus item map")
	var cat = map[string]int{
		"februari" : 17,
		"maret" : 18,
		"april" : 19,
	}

	fmt.Println(len(cat))
	fmt.Println(cat)

	delete(cat, "februari")

	fmt.Println(len(cat))
	fmt.Println(cat)


	fmt.Println("\n\nkombinasi slice dan map")
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue",   "gender": "male"},
		map[string]string{"name": "chicken red",    "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println("\n\ndata yang berbeda")
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, val := range data{
		fmt.Println(val["gender"],val["name"],val["address"],val["id"],val["color"])
	}
}