const Pi = 3.14

// circle struct
type Circle struct { // HL
	radius float64
} // HL

// return a pointer to a circle
func New(radius float64) *Circle { // HL
	return &Circle{radius}
} // HL

// function on a circle pointer receiver
func (c *Circle) Area() float64 { // HL
	return Pi * c.radius * c.radius
} // HL
