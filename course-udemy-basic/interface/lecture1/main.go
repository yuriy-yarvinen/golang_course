package main

import "fmt"

type SomeInterface interface {
	SomeMethod(in string)
}

type SomeStruct struct{}

func (s SomeStruct) SomeMethod(in string) {}

type Stringer interface {
	String() string
}

type ByValueExample struct {
	Name string
}

func (e ByValueExample) String() string {
	e.Name = "Name by value"
	return e.Name
}

type ByPointerExample struct {
	Name string
}

func (e *ByPointerExample) String() string {
	e.Name = "Name by pointer"
	return e.Name
}

type Nameable interface {
	GetName() string
}
type User struct {
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}
func (u User) String() string {
	return fmt.Sprintf("Юзеру %s столько лет %d", u.Name, u.Age)
}

type City struct {
	Name string
	Addr string
}

func (c City) String() string {
	return fmt.Sprintf("Город %s адрес %s", c.Name, c.Addr)
}

func Logger(s Stringer) {
	fmt.Println(s.String())
}

type KeyValueStorage interface {
	Get(key string) string
	Set(key, value string)
}

type InMemStorage struct {
	data map[string]string
}

func (s *InMemStorage) Get(key string) string {
	return s.data[key]
}

func (s *InMemStorage) Set(key, value string) {
	s.data[key] = value
}
func (s *InMemStorage) Delete(key string) {
	delete(s.data, key)
}

func PrintThisValue(value interface{}) {
	switch casted := value.(type) {
	case int:
		fmt.Println("this is int:", casted)
	case string:
		fmt.Println("this is string:", casted)
	case Nameable:
		fmt.Println("this is something with name", casted.GetName())
	default:
		fmt.Println("WTF is this")
	}
}

func main() {

	kv := InMemStorage{data: make(map[string]string)}
	kv.Set("test", "123")
	kv.Set("test2", "12sds3")
	kv.Delete("test2")
	fmt.Println(kv.Get("test"))
	fmt.Println(kv.Get("test2"))

	var v SomeInterface

	fmt.Printf("value %#+v\n", v)
	fmt.Printf("pointer %#+v\n", &v)

	var s *SomeStruct

	fmt.Println(s == nil)

	var i SomeInterface
	fmt.Println(i == nil)

	i = s

	fmt.Println(i == nil)
	fmt.Printf("value %#v\n", i)

	u := User{
		Name: "Yuriy",
		Age:  36,
	}
	c := City{
		Name: "Spb",
	}
	Logger(u)
	Logger(c)

	var obj Stringer
	obj = ByValueExample{Name: "asd"}
	// obj = &ByValueExample{}

	// obj = &ByPointerExample{}
	// obj = ByPointerExample{} // not work
	fmt.Println(obj.String())

	// any or empty interface
	var a interface{} // any
	a = 1
	a = "sdfsdf"
	a = []string{"one", "two", "three"}

	sliceStr, ok := a.([]string)

	fmt.Println(a)
	fmt.Println(sliceStr, ok)

	someUser := User{Name: "sadfas"}
	PrintThisValue(someUser)
}
