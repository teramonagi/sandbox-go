package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/teramonagi/sandbox-go/hoge"
	"github.com/teramonagi/sandboxgo2"
)

func hello() string {
	return ("Hello, world!")
}

func add(x int, y int) int {
	return (x + y)
}

func swap(x string, y string) (string, string) {
	return y, x
}

func pow(x float64, n float64, limit float64) float64 {
	if value := math.Pow(x, n); value < limit {
		return value
	}
	return limit
}
func deferHelloWorld() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferPrint() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

type Vertex struct {
	X int
	Y int
}

// Method v.s. Function
func (v Vertex) onlyX1() int {
	return v.X
}
func onlyX2(v Vertex) int {
	return v.X
}
func (v *Vertex) scale(x int) {
	v.X = v.X * x
	v.Y = v.Y * x
}

// Interface
type I interface {
	hoge() int
}

type Int int

func (x Int) hoge() int {
	return (1)
}
func (v *Vertex) hoge() int {
	if v == nil {
		return -1
	}
	return v.X
}

func main() {
	{
		fmt.Println(hello())
		fmt.Println("The time is", time.Now())
		fmt.Println("My favorite number is", rand.Intn(10))
		fmt.Println(math.Pi)
		fmt.Println("sqrt(2) is equal to", math.Sqrt(2))
		fmt.Println("1 + 1 =", add(1, 1))
		x, y := swap("Hello", "World")
		fmt.Println(x, y)
	}
	{
		var i int = 3
		j := 5
		fmt.Printf("%v + %v = %v\n", i, j, add(i, j))
		var (
			a string = "hoge"
			b string = "moge"
		)
		fmt.Println(a + b)
		const World = "世界"
		fmt.Println(World)
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(pow(10, 3, 10))
	fmt.Println(pow(10, 3, 100000))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good!")
	}

	deferHelloWorld()

	deferPrint()

	// Exernal
	{
		fmt.Println(sandboxgo2.Sqrt(2.0))
	}
	ip := 10
	p := &ip
	*p = 100
	*p = *p + 1
	fmt.Println(ip)

	fmt.Println(Vertex{1, 2})
	{
		v := Vertex{X: 100}
		fmt.Println(v)
		p := &v
		p.Y = 222
		fmt.Println(v)
	}

	{
		// Array and Slices
		var a [3]int = [3]int{1, 2, 3}
		var s []int = []int{1, 2, 3, 4, 5}
		fmt.Println(a)
		fmt.Println(s, len(s), cap(s))

		ss := make([]int, 0, 100)
		fmt.Println(ss, len(ss), cap(ss))
		ss = append(ss, 1, 2, 3, 4, 777)
		fmt.Println(ss, len(ss), cap(ss))
	}

	// range
	for i, v := range []int{1, 2, 4, 8} {
		fmt.Println(i, v)
	}

	// map data structure
	m := make(map[string]Vertex)
	m["hoge"] = Vertex{1, 2}
	fmt.Println(m)

	mm := map[string]Vertex{
		"hoge": {1, 2},
		"hage": {3, 7},
		"mage": {8, 9},
	}
	fmt.Println(mm["hage"])

	// closure
	iterator := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	iter := iterator()
	fmt.Println(iter())
	fmt.Println(iter())
	fmt.Println(iter())
	fmt.Println(iter())

	//Methods
	{
		fmt.Println(onlyX2(Vertex{3, 3}))
		v := Vertex{3, 3}
		fmt.Println(v.onlyX1())
		v.scale(7)
		fmt.Println(v)
	}

	//interface
	{
		var iii I = Int(111)
		fmt.Println(iii)
		var v *Vertex
		fmt.Println(v.hoge())
		v = &Vertex{1, 2}
		fmt.Println(v.hoge())
	}

	{
		switchWithType := func(i interface{}) {
			switch v := i.(type) {
			case int:
				fmt.Println(v, v*2)
			case string:
				fmt.Println(v, len(v))
			default:
				fmt.Println(v)
			}
		}
		var i interface{} = "hello"
		s, ok := i.(string)
		fmt.Println(s, ok)
		switchWithType(21)
		switchWithType("hello")
	}

	{
		fmt.Println("Go routine")
		say := func(s string) {
			for i := 0; i < 5; i++ {
				time.Sleep(100 * time.Millisecond)
				fmt.Println(s)
			}
		}

		go say("world")
		say("hello")
	}

	{
		sum := func(s []int, c chan int) {
			sum := 0
			for _, v := range s {
				sum += v
			}
			c <- sum
		}
		x := []int{1, 2, 3, 4}
		c := make(chan int)
		go sum(x[len(x)/2:], c)
		go sum(x[:len(x)/2], c)
		fmt.Println(c)
		a, b := <-c, <-c
		fmt.Println(a, b)
	}

	//
	{

		boring := func(msg string, c chan string) {
			for i := 0; ; i++ {
				c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}

		c := make(chan string)
		go boring("boring!", c)
		for i := 0; i < 5; i++ {
			//When the main function executes <–c, it will wait for a value to be sent.
			//Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
			fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
		}
		fmt.Println("You're boring; I'm leaving.")
	}
	fmt.Println(hoge.Hoge1(3))
}
