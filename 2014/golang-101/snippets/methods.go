package main

import "fmt"

type Coords struct {
	lat float64
	lng float64
}

func (c Coords) String() string {
	return fmt.Sprintf("%v,%v", c.lat, c.lng)
}

func main() {
	c := Coords{
		lat: 3.12,
		lng: 101.3,
	}
	fmt.Println(c.String())
}
