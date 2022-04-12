package main

import (
	"fmt"
	"sync"
)

const a9 int16 = 17

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return result
}

// interface defines behaviours
// type Writer interface {
// 	Write([]_byte (int, error))
// }

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	// Arrays and Slices
	grades := [...]int{97, 99, 93}
	fmt.Println(grades)

	var students [3]string
	fmt.Println(students)

	students[0] = "S1"
	fmt.Println(students)
	fmt.Println(students[0])
	fmt.Println(len(students))

	var iMatrix [3][3]int
	iMatrix[1] = [3]int{1, 2, 3}
	fmt.Println(iMatrix)

	// deep copy by default
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 900000
	fmt.Println(a, b)

	// soft copy using pointer
	c := &a
	c[1] = 5555555
	fmt.Println(a, c)

	// array : [...]
	// slice : []
	// most things common about the two

	// slicing : works for both array and slice
	s := []int{1, 2, 5, 3, 9}
	fmt.Println(s[:])
	fmt.Println(s[0:])
	fmt.Println(s[:5])
	// [inclusive : exclusive]

	// builtin make() : makes slice
	t := make([]int, 3)
	fmt.Println(t, len(t), cap(t))

	// arrays have fixed size
	// slices have variable size
	t = append(t, 1, 1024, 777)
	fmt.Println(t, len(t), cap(t))

	// if using slice : spread operator
	t = append(t, []int{1, 100, 945, 223}...)
	fmt.Println(t, len(t), cap(t))
	t[5] = 7
	fmt.Println(t, len(t), cap(t))

	// remove from begin
	u := t[1:]
	fmt.Println(u, len(u), cap(u))

	// remove from end
	v := t[:len(t)-1]
	fmt.Println(v, len(v), cap(v))

	// remove from somewhere else
	w := append(t[:2], t[3:]...)
	fmt.Println(w, len(w), cap(w))

	// Channels

	// since go was born in world of multi-core and multi-thread world
	// channels allow data between multiple go routines
	// to avoid race conditions and memory issues

	var wg = sync.WaitGroup{}

	ch := make(chan int)
	wg.Add(2)
	// anon function
	go func() {
		// direction of data flow : channel to i
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// direction of data flow : to channel
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

	// Constants
	const a1 int32 = 9 * iota
	fmt.Println(a1)

	// Control Flow
	if 3 > 5 {
		fmt.Println("The test is True.")
	} else {
		fmt.Println("False.")
	}

	// Map
	statePopulation := map[string]int{
		"Cali":     1000,
		"Texas":    90,
		"Flo":      811,
		"New York": 192,
	}
	// init; condition
	if pop, ok := statePopulation["Flo"]; ok {
		fmt.Println(pop)
	}

	val := 69
	// break(s) aren't required
	switch val {
	case 69:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	default:
		fmt.Println("l")
	}

	// tagless switch
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than eq 10")
		fallthrough // anti-break (break was implicit)
	case i <= 20:
		fmt.Println("Less than eq 20")
	}

	// Defer-Panic-Recover
	// fmt.Println("Start")
	// defer fmt.Println("Middle")
	// fmt.Println("end")

	// // last defer is executed first after synchronous program is done
	// // FILO and LIFO
	// defer fmt.Println("Middle1")
	// defer fmt.Println("Middle2")
	// defer fmt.Println("Middle3")

	// // defers are executed towards the end

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // closing resource before using it in body but with defer
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// // defer isn't good for loops

	// //
	// a2 := "start"
	// defer fmt.Println(a2)
	// a2 = "end"

	// panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	// panic("error happened")

	// panic in web server
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err = http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// panics happen after defer
	// normal -> defer -> panic

	// anonyonomous fnc
	// fmt.Println("Start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error : ", err)
	// 	}
	// }()
	// panic("something bad")
	// fmt.Println("done")

	// Functions
	sayMessage("Hello Go!")
	sayMessage("Hello")
	fmt.Println(sum(1, 1, 2, 3, 4, 5, 6, 7, 3, 21, 2, 4, 12, 32, 21, 4))

	// GoRoutines
	var wg1 = sync.WaitGroup{}

	// os threads aren't used as in other languages
	// layer of abstraction over threads
	// go keyword for things to work in thread
	msg := "1"
	wg1.Add(1)
	go func() {
		fmt.Println(msg)
		wg1.Done()
	}()
	msg = "2"
	// time.Sleep(100 * time.Millisecond) // bad practice
	// hence using wait group
	wg1.Wait() // good practice

	// Interfaces

	// Looping
	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	statePopulation1 := map[string]int{
		"Cali":     1000,
		"Texas":    90,
		"Flo":      811,
		"New York": 192,
	}
	for k, v := range statePopulation1 {
		fmt.Println(k, v)
	}

	for k, v := range "GeeksForGeeks" {
		fmt.Println(k, v)
	}

	for _, v := range "GeeksForGeeks" {
		fmt.Println(v)
	}
}

// var wg = sync.WaitGroup{} // used globally here
// var counter = 0
// var m = sync.RWMutex{} // to maintain mutual exclusion
// // any reader will wait for all writers to get done
// // any writer will wait for all readers to get done

// func main() {
// 	runtime.GOMAXPROCS(100) // setting number of threads
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		m.RLock() // works
// 		go sayHello()
// 		m.Lock() // works
// 		go increment()
// 		// number of threads
// 		fmt.Println("Threads : ", runtime.GOMAXPROCS(-1))
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	// m.RLock() // didn't work as required
// 	fmt.Println("Hello Cnt : ", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	// m.Lock() // didn't work as required
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }

// Variables
// // all variable must be used : unused variables are to be removed
// package main

// import "fmt"

// var i float64 = 100 // shadowed by local variable i

// var I int = 420 // global exportable variable //

// // 3 possible scope:
// // - block
// // - lowercase : scoped to package
// // - uppercase
// // NO private scope

// var (
// 	actorName string = "SRK"
// 	companion string = "KS"
// 	number    int    = 10
// )

// func main() {
// 	// decl 1
// 	var i int
// 	i = 42
// 	fmt.Println(i)

// 	// decl 2
// 	var j float32 = 50
// 	j = float32(i)
// 	fmt.Println(j)

// 	// decl 3
// 	p := 69.0
// 	fmt.Println(p)

// 	// %value, %Type
// 	fmt.Printf("%v, %T\n", i, i)
// 	fmt.Printf("%v, %T\n", j, j)
// 	fmt.Printf("%v, %T\n", p, p)
// }

// Primitives
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n bool // default value is 0
// 	fmt.Printf("%v, %T\n", n, n)

// 	// int : int8, int64

// 	// unary, binary, shift operations
// }

// Pointers
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// copy by value
// 	a := 42
// 	b := a
// 	fmt.Println(a, b)
// 	a = 10
// 	fmt.Println(a, b)

// 	// copy by ref
// 	var c int = 42
// 	var d *int = &c
// 	fmt.Println(c, *d)
// 	c = 10
// 	fmt.Println(c, *d)
// 	*d = 1e6
// 	fmt.Println(c, *d)

// 	// pointer arithmetic to be avoided
// 	// if inevitable, import "unsafe"

// 	// uninitialized pointers = nil

// }

// Maps_and_Structs
// package main

// import (
// 	"fmt"
// 	"reflect" // for tags
// )

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	companions []string //slice, not array
// }

// type Animal struct {
// 	Name   string `required max : "100"` // tag : required + max length
// 	Origin string
// }

// type Bird struct {
// 	Animal   // composition : has a : relationship
// 	SpeedKPH float32
// 	CanFly   bool
// }

// func main() {
// 	statePopulation := map[string]int{
// 		"Cali":     1000,
// 		"Texas":    90,
// 		"Flo":      811,
// 		"New York": 192,
// 	}

// 	m := make(map[string]int, 10)

// 	// add
// 	statePopulation["Geo"] = 10092018
// 	// order of keys will not be maintained
// 	fmt.Println(statePopulation, m)

// 	// del
// 	delete(statePopulation, "Geo")
// 	fmt.Println(statePopulation, m)

// 	fmt.Println(statePopulation["Geo"])

// 	// check if presence
// 	pop, ok := statePopulation["Flo"]
// 	fmt.Println(pop, ok)

// 	// STRUCTS
// 	aDoctor := Doctor{
// 		number:    3,
// 		actorName: "John",
// 		companions: []string{
// 			"Liz",
// 			"Jo",
// 			"Sarah",
// 		},
// 	}
// 	fmt.Println(aDoctor)

// 	// Way 1
// 	b := Bird{}
// 	b.Name = "Emu"
// 	b.Origin = "Aus"
// 	b.SpeedKPH = 60
// 	b.CanFly = false
// 	fmt.Println(b)

// 	// Way 2
// 	c := Bird{
// 		Animal:   Animal{Name: "Emu", Origin: "Aus"},
// 		SpeedKPH: 60,
// 		CanFly:   false,
// 	}
// 	fmt.Println(c)

// 	// tagging - fectching tag
// 	t := reflect.TypeOf(Animal{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag)
}