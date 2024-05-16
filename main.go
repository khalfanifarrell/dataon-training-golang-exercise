package main

import "fmt"

const (
	tyreTypeRubber = "Karet"
	tyreTypeWood   = "Kayu"
	tyreTypeMetal  = "Besi"

	doorTypeRight = "right"
	doorTypeLeft  = "left"

	soundTypeKnock = "Knock"
	soundTypeBeep  = "Beep"

	errInvalidDoorType = "Invalid door type"
)

type Car struct {
	totalTyre int
	tyreType  string
}

func main() {
	// Define Car
	var Brio Car

	Brio.totalTyre = 4
	Brio.tyreType = tyreTypeWood

	Brio.printTotalTyre()
	Brio.printTyreType()

	// Change tyre type
	Brio.changetyreType(tyreTypeRubber)

	Brio.printTyreType()

	// Knock and open doors
	Brio.knockDoor(doorTypeRight)
	Brio.knockDoor(doorTypeLeft)
	Brio.openDoor(doorTypeRight)
	Brio.openDoor(doorTypeLeft)
}

func (c *Car) changetyreType(tyreType string) {
	fmt.Printf("Changing tyre type from %s to %s", c.tyreType, tyreType)
	fmt.Println()
	fmt.Println()
	c.tyreType = tyreType

}

func (c *Car) knockDoor(doorType string) {
	if doorType == doorTypeRight {
		fmt.Printf("Knocking %s door...", doorTypeRight)
		fmt.Println()
		fmt.Println(soundTypeKnock)
		fmt.Println()
		return
	}

	if doorType == doorTypeLeft {
		fmt.Printf("Knocking %s door...", doorTypeLeft)
		fmt.Println()
		fmt.Println(soundTypeBeep)
		fmt.Println()
		return
	}

	fmt.Println(errInvalidDoorType)
}

func (c *Car) openDoor(doorType string) {
	if doorType == doorTypeRight {
		fmt.Printf("Opening %s door...", doorTypeRight)
		fmt.Println()
		fmt.Println(soundTypeBeep)
		fmt.Println()
		return
	}

	if doorType == doorTypeLeft {
		fmt.Printf("Opening %s door...", doorTypeLeft)
		fmt.Println()
		fmt.Println(soundTypeKnock)
		fmt.Println()
		return
	}

	fmt.Println(errInvalidDoorType)
}

func (c *Car) printTotalTyre() {
	fmt.Printf("The current total tyres: %d", c.totalTyre)
	fmt.Println()
	fmt.Println()
}

func (c *Car) printTyreType() {
	fmt.Printf("The current tyre type is: %s", c.tyreType)
	fmt.Println()
	fmt.Println()
}

// func main() {
// 	var a = 4
// 	var b *int = &a

// 	*b = 10

// 	fmt.Println(a)

// 	var total = 0

// 	var m sync.Mutex

// 	for i := 0; i < 100; i++ {
// 		for j := 0; j < 100; j++ {
// 			go func() {
// 				m.Lock()
// 				total++
// 				m.Unlock()
// 			}()
// 		}
// 	}

// 	// time.Sleep(5 * time.Second)
// 	fmt.Println(total)
// }

//
// Previous Training
//
// func main() {
// 	// evenOrOdd := checkOddOrEven(1)
// 	// grading := gradeTest(101)

// 	// fmt.Println("Even or Odd? ", evenOrOdd)
// 	// fmt.Println("Your Grade? ", grading)

// 	// printEvenNumbers(9)
// 	// printLastSliceValue([]string{"John", "Jean"})
// 	// sumRangeNumber(3)
// 	// sumRangeNumberOfOddAndMultipleOfFive([]int{1, 2})
// 	var square Square
// 	square.sideLength = 5
// 	fmt.Println(square.calcArea())

// 	var alquran Book
// 	alquran.pages = 300
// 	fmt.Println(alquran.calcThickness())

// 	var bible Book
// 	bible.pages = 299
// 	fmt.Println(bible.calcThickness())
// }

// func checkOddOrEven(number int) (res string) {
// 	if number%2 == 0 {
// 		return "even"
// 	}

// 	return "odd"
// }

// func gradeTest(number int) (res string) {
// 	if number > 100 {
// 		return "S"
// 	}

// 	if number >= 90 && number <= 100 {
// 		return "A"
// 	}

// 	if number >= 80 && number <= 89 {
// 		return "B"
// 	}

// 	if number >= 70 && number <= 79 {
// 		return "C"
// 	}

// 	if number >= 60 && number <= 69 {
// 		return "D"
// 	}

// 	if number <= 59 {
// 		return "E"
// 	}

// 	return "F"
// }

// func printEvenNumbers(number int) {
// 	for i := 1; i < number; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func printLastSliceValue(slice []string) {
// 	fmt.Println(slice[len(slice)-1])
// }

// func sumRangeNumber(number int) {
// 	var res int

// 	for i := 1; i <= number; i++ {
// 		res += i
// 	}

// 	fmt.Println(res)
// }

// func sumRangeNumberOfOddAndMultipleOfFive(slice []int) {
// 	var res int

// 	for _, v := range slice {
// 		if v%2 != 0 && v%5 == 0 {
// 			res += v
// 		}
// 	}

// 	fmt.Println(res)
// }

// type (
// 	Square struct {
// 		sideLength int
// 	}

// 	Book struct {
// 		title  string
// 		author string
// 		pages  int
// 	}
// )

// func (s Square) calcArea() int {
// 	return s.sideLength * s.sideLength
// }

// func (b Book) calcThickness() bool {
// 	if b.pages >= 300 {
// 		return true
// 	}

// 	return false
// }
