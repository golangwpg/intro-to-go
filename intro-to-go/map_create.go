	var ageMap map[string]int

	ageMap = make(map[string]int) // HL

	// add to map
	ageMap["Mark"] = 43 // HL

	// check if key exists
	_, ok := ageMap["Andrew"] // HL
	if !ok {
		fmt.Println("Andrew not in map")
	}
