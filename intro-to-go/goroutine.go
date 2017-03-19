func fetch(uri string) error { // HL
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
} // HL

func main() { // HL
	go fetch("https://www.google.com") // HL

	time.Sleep(2000)
	fmt.Println("done")
} // HL
