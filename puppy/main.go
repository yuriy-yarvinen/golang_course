package main

import (
	puppy "github.com/yuriy-yarvinen/go-puppy-for-course"
)

func main() {
	Puppy := puppy.NewPuppy()
	Puppy.Name = "Buddy"
	Puppy.Age = 2
	Puppy.BarkText = "Woof\nWoof\nWoof"
	Puppy.MakeBark()
}
