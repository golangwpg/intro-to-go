func main() {
	done := make(chan (bool))      // HL
	results := make(chan (string)) // HL

	sites := []string{"https://google.com", "https://yahoo.com"}

	for _, site := range sites {
		go fetch(site, results, done) // HL
	}

	for i := 0; i < len(sites); {
		select { // HL
		case data := <-results: // HL
			fmt.Printf("resp body size: %d\n", len(data))
		case <-done: // HL
			i++
		} // HL
	}

}
