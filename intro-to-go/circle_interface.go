type Shape interface { // HL
	Area() float64 // HL
} // HL

func main() {
	var s Shape // HL

	s = circle.New(3.0) // HL

	area := s.Area() // HL

	fmt.Printf("area of shape: %f\n", area)
}
