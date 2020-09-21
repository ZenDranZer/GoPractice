package main

import "fmt"

const max16 = 65535
const kmh_multiple = 1.60934

type car struct {
	gas   uint16
	brake uint16
	steer int16
	speed float64
}

func (c car) kmp() float64 {
	return float64(c.gas) * (c.speed / max16)
}

func (c car) mph() float64 {
	return float64(c.gas) * (c.speed / max16 / kmh_multiple)
}

func main() {
	car_one := car{65000, 12, 11, 225}
	fmt.Println(car_one.gas)
	fmt.Println(car_one.kmp())
	fmt.Println(car_one.mph())
}
