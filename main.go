package main

import "fmt"

func main() {
	// vars
	var a = "var"
	b := "second var"
	fmt.Println("adsasd " + a + " " + b)
	// slices (arrays) and maps
	var s = make([]string, 0)
	var m = make(map[string]string)
	s = append(s, "some string")
	m["some key"] = "some value"
	fmt.Println(s, m)
	//pointer operator to retrieve the address in memory of some variable (&) and dereference (*)
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)
	//if, else if, else
	var x = 2
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else if x == 2 {
		fmt.Println("x is equal to 2")
	} else {
		fmt.Println("X is not equal to 1")
	}
	//switch case
	switch x {
	case 1:
		fmt.Println("Found foo")
	case 2:
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//iterate over slice
	nums := []int{2,4,6,8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}

/*func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}*/
	//objects
/*type Person struct {
	Name string
	Age int
}
func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}
func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()
}*/

/*
type Dog struct {}
func (d *Dog) SayHello() {
 fmt.Println("Woof woof")
}
func main() {
 var guy = new(Person)
 guy.Name = "Dave"
 Greet(guy)
 var dog = new(Dog)
 Greet(dog)
}
 */