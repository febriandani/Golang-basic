package switchcase

import "fmt"

func main(){
	var point = 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	fmt.Print("\nPermanfaatan case banyak kondisi\n\n")

	var point1 = 6

	switch point1 {
	case 8:
		fmt.Println("Perfect")
	case 7,6,5,4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	fmt.Print("\nKurung kurawal pada case & default\n\n")

	var point2 = 3

	switch point2 {
	case 8:
		fmt.Println("Perfect")
	case 7,6,5,4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("you can do better")
		}
	}

	fmt.Print("\nSwitch case dengan gaya if - else\n\n")
	var point3 = 6

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}


}