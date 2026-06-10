package main

type intStringMap map[int]string

func (m intStringMap) print() {
	for id, name := range m {
		println("ID:", id, "Name:", name)
	}
}

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

	userIDs := make([]int, 0, 5)
	userIDs = append(userIDs, 101, 102, 103)
	println("User IDs:", userIDs[0], userIDs[1], userIDs[2])
	println("Length:", len(userIDs), "Capacity:", cap(userIDs))

	userNames := make(intStringMap, 50)
	userNames[1] = "Alice"
	userNames[2] = "Bob"
	userNames[3] = "Bob Smith"

	println("User Names:")
	userNames.print()

	for index, value := range userNames {
		println("Index:", index, "Value:", value)
	}
}
