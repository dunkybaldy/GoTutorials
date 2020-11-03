/*
Program Name: GoLang Tutorials
Version: 1.0
Description: For noobs just starting https://www.youtube.com/watch?v=C8LgvuEBraI
Author: github.com/u/dunkybaldy
Creation Date: 20/10/20
*/
package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// test_out_go_syntax()
	// concurrencyTestWithWaitGroup()
	// concurrencyTestWithChannel()
	selectChannel()

	// fmt.Scanln() // blocking call - allows gorountines to execute, press enter to exit
} // End of program

func selectChannel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select { // Whenever a channel is ready, receive from that channel
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func concurrencyTestWithChannel() {
	c := make(chan string)
	go countWithChannel("sheep", c) // go keyword -> function call becomes a gorountine// immediately exectue function wrapper
	// go count("fish")  // if both a gorountines the program will immediately exit because the "main" gorountine will have finished all it's line

	for msg := range c { // Written like this we don't need to manually check for close messages anymore
		fmt.Println(msg)
	}
}

func countWithChannel(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500) // 0.5 seconds
	}

	close(c) // never close a channel as a receiver, because you don't know when the recipient has finished sending all the messages
}

func concurrencyTestWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep") // go keyword -> function call becomes a gorountine
		wg.Done()
	}() // immediately exectue function wrapper
	// go count("fish")  // if both a gorountines the program will immediately exit because the "main" gorountine will have finished all it's lines

	wg.Wait()
}

func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500) // 0.5 seconds
	}
}

func test_out_go_syntax() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Max(73.15, 92.46))
	// Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	// Printing the value of `𝜋`
	fmt.Println(math.Pi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

	// var <name> <type> =
	// var <name> =
	var x int = 20
	var aByte byte = 'a'
	var aRune rune = 'ä'
	var aString string = "blahablah\n"
	// var result string = fmt.Sprintf("%a %a %a", x, aByte, aRune)
	fmt.Printf("%d %b %U %s", x, aByte, aRune, aString)

	y := 5

	if y > 5 {
		fmt.Println("y is bigger than 5")
	} else if y == 5 {
		fmt.Println("y is equal to 5")
	} else {
		fmt.Println("what")
	}

	var array1 [5]int
	array1[2] = 7
	fmt.Println(array1)

	array2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(array2)

	slice1 := []int{0, 1, 2, 3, 4}
	fmt.Println(slice1)
	slice2 := append(slice1, 5)
	fmt.Println(slice1, slice2)

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["pentagon"] = 5
	fmt.Println(vertices)
	fmt.Println(vertices["square"], vertices["triangle"])
	delete(vertices, "pentagon")
	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d", i)
	}

	j := 0
	for j < 5 {
		fmt.Printf("j = %d\n", j)
		j++
	}

	slice3 := []string{"a", "b", "c"}
	for index, value := range slice3 {
		fmt.Println("Index", index, "Value", value) // space separated
		slice3 = append(slice3, "d")                // no runtime error, doesn't output d because line 87 slice3 and line 85 slice3 are different scopes, different objects

	}

	map2 := make(map[string]string)
	map2["a"] = "alpha"
	map2["b"] = "beta"
	for key, value := range map2 {
		fmt.Println("Key:", key, "Value", value)
	}

	fmt.Println(sum(2, 7))

	k := -1

	result, err := sqrt(float64(k))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	p := person{Name: "Duncan", Age: 25}
	fmt.Println(p)
	fmt.Println(p.Name)

	num := 0
	fmt.Println(num)
	fmt.Println(&num)
	increment(&num)
	fmt.Println(num)
}

func increment(x *int) { // asterisk to accept a pointer
	*x++ // dereference the pointer to access the value at the pointer
}

func sum(x int, y int) int { // lowercase sum means unexported (other modules/packages cannot access this function)
	return x + y
}

func sqrt(x float64) (float64, error) { // go dodesn't have exceptions
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

type person struct {
	Name string
	Age  int
}
