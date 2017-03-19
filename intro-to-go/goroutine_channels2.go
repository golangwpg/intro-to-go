func fetch(uri string, results chan string, done chan bool) { // HL
	defer func() { // HL
		done <- true // HL
	}() // HL

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Printf("error fetching uri: %s\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading resp body: %s\n", err)
		return
	}

	results <- string(body) // HL
} // HL

