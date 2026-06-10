package main

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"GitHub": "https://www.github.com",
		"Go":     "https://golang.org",
	}

	for name, url := range websites {
		println(name, "->", url)
	}

	delete(websites, "GitHub")
	println("After deletion:")
	for name, url := range websites {
		println(name, "->", url)
	}

	websites["StackOverflow"] = "https://stackoverflow.com"
	println("After adding StackOverflow:")
	for name, url := range websites {
		println(name, "->", url)
	}
}
