// package intro
// defined by existing in a folder called intro

package intro

import (
	"fmt"

	"os"
	"os/signal"

	"github.com/gopherjs/gopherjs/compiler/natives/src/sync"
)

//--------------------------------
// Variable and Const declarations
//--------------------------------

// keyword var, name, type
var anInt int

// available from packages outside this package
var ExportedInt int

// extra examples of different types and stuff

// groups of vars
var (
	aString string
	aFloat  float64
)

// variable declaration with assignment
var newString = "this is a string"

// groups with assignments
var (
	newerString = "this is a newer string"
	newInt      = 42
)

// mix
var (
	evenNewerString = "this is an even newer string"
	newerInt        int
)

// constants must have a value when they are declared, but the syntax should look familiar
const constInt = 42

const (
	newConstInt = 42
	constString = "This is a constant string"
)

// iota - autoincrementing declaration
const (
	code200 = iota + 200
	code201
)

//--------
// Structs
//--------

/*
 * A struct is used for organizing data, and for defining methods that can be called to do something
 * with that data
 */

// struct definition - uppercase in go means it is exported outside the package
type MyStruct struct {
	// struct fields: name, type
	MyInt    int
	MyString string
}

// only usable inside the package
type myStruct struct {
}

// defining a method on a struct pointer
func (s *MyStruct) PrintInt() {
	fmt.Println(s.MyInt)
}

func (s *MyStruct) PrintString() {
	fmt.Println(s.MyString)
}

//-----------
// Interfaces
//-----------

/*
 * An interface is an abstract type that defines functions. Any concrete type (a struct) that implements
 * all of the functions defined in the interface, can then be used in functions that expect that interface.
 */

type IntPrinter interface {
	PrintInt()
}

type StringPrinter interface {
	PrintString()
}

// the empty interface{}

//----------
// Functions
//----------

/*
 * Functions in go are first class - meaning that they can also be variables
 * not just static function declarations
 */

// static function definition
func staticFunc() {
	fmt.Println("Hello!")
}

// function that expects an interface
func PrintInt(i IntPrinter) {
	i.PrintInt()
}

func PrintString(s StringPrinter) {
	s.PrintString()
}

// actually passing a concrete type to PrintInt()

func testPrintInt() {
	// := == type inference - the compiler knows that printer, is a pointer to MyStruct
	// := may only be used inside of a function, otherwise you must specify the type, or
	// assign the variable
	// & is the address of operator - it means we're taking the pointer of something
	printer := &MyStruct{MyInt: 42, MyString: "my string"}
	PrintInt(printer)
}

func testPrintString() {
	printer := &MyStruct{MyInt: 42, MyString: "my string"}
	PrintString(printer)
}

//------------
// Collections
//------------

// Slices - Dynamic Arrays

// initializes strSlc as a slice of strings
var strSlc []string

// make is the memory allocation keyword in go - it is used to allocate memory for slices, maps, and channels
// when using it with a slice type, the second parameter is the capacity
var newStrSlc = make([]string, 0)

// initializing a slice with default values
var intSlc = []int{1, 2, 3}

// iterating over slices
func Iter(slc []int) {
	// i == index, j == integer in slice
	for i, j := range slc {
		fmt.Println("index: ", i)
		fmt.Println("integer: \n", j)

		fmt.Println("by index: ", slc[i])
	}

	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}

	for i := len(slc) - 1; i >= 0; i-- {
		fmt.Println(slc[i])
	}
}

// MAPS! (Hash tables, KV stores)
var m map[string]string

// make only requires a single argument with maps - the map - but, can also accept a capacity argument
var nm = make(map[string]string)

// with a capacity of 2
var nmc = make(map[string]string, 2)

// no capacity, but allocated
var nnm = map[string]string{"Hello!": "Goodbye!", "You say": "I say"}

// mixing types
var im map[int]string

// iterating over maps
func MapIter(m map[string]string) {
	// k == key, v == value
	// order is not preserved - map iteration is RANDOM
	for k, v := range m {
		fmt.Println(k, " ", v)
	}
	fmt.Println(m["Hello!"])
}

// channels

// A channel is a construct in Go that is used to pass memory between Goroutines - or threads
var bc chan bool

var sc = make(chan string)

func ChanIter(c chan string) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for s := range c {
			fmt.Println(s)
			wg.Add(-1)
		}
	}()
	c <- "Hello"
	c <- "Channel!"
	wg.Wait()
	// race condition
}

func BetterChannelExample(c chan string) {
	b := make(chan bool)
	go func() {
		for s := range c {
			fmt.Println(s)
		}
	}()

	go func() {
		c <- "Hello!"
		c <- "Goodbye!"
		b <- true
	}()

	select {
	case _ := <-b:
	}
}

// select - blocks on select until first condition returns, or forever if select is empty
func Select() {
	c := make(chan bool)

	// Blocks until something is received on c and then exits the select
	select {
	case _ = <-c:
		fmt.Println("Exiting!")
	}

	// Will block forever
	select {}
}

// assume this actually runs an http server
func serveHttp() error {
	return nil
}

// "serve http" and handle cleanup with signals
func ServeHttp() error {
	// run http server in separate thread
	go serveHttp()

	// do more work
	c := make(chan<- os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	select {
	case s := <-c:
		fmt.Printf("Received signal %d, exiting\n", s)
		// close a bunch of files - or maybe a database connection, do cleanup work, etc
		os.Exit(0)
	}
}
