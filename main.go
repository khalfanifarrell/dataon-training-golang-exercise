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
