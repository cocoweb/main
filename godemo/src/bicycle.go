package main

import (
	"fmt"
)

type Part struct {
	name        string
	Description string
	NeedsSpare  bool
}
func (part Part) Name() string {
	return part.name
}
func (part Part) String() string {
	return fmt.Sprintf("%s: %s", part.name, part.Description)
}



func (part Part) SetName(name string) {
	part.name = name
}

type Parts []Part

func (parts Parts) Spares() (spares Parts) {
	for _, part := range parts {
		if part.NeedsSpare {
			spares = append(spares, part)
		}
	}
	return spares
}

type Bicycle struct {
	Size string
	Parts
}

var (
	RoadBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "23", true},
		{"tape_color", "red", true},
	}

	MountainBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "2.1", true},
		{"front_shock", "Manitou", false},
		{"rear_shock", "Fox", true},
	}

	RecumbentBikeParts = Parts{
		{"chain", "9-speed", true},
		{"tire_size", "28", true},
		{"flag", "tall and orange", true},
	}
)

func main() {
	roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
	mountainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
	recumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

	fmt.Println(roadBike.Spares())
	fmt.Println(mountainBike.Spares())
	fmt.Println(recumbentBike.Spares())

	comboParts := Parts{}
	comboParts = append(comboParts, mountainBike.Parts...)
	comboParts = append(comboParts, roadBike.Parts...)
	comboParts = append(comboParts, recumbentBike.Parts...)

	fmt.Println(len(comboParts), comboParts[9:])
	fmt.Println(comboParts.Spares())
}
