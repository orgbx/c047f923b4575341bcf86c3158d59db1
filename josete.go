package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	//"io"
)

type Person struct {
	name string
	age  int
}

func josete() {

	os.Stdin.Sync()
	fmt.Print("Hcj")
	body()
	//readFile()
}

func body() {
	var x int
	x = 3
	y := 4
	print(x + y)

	arra := []int{3, 4, 5}

	arra = append(arra, 3, 5, 6)
	fmt.Println(arra)

	fmt.Print(arra[0])

	mapa := map[string]int{"one": 2}

	print(mapa["one"])

	chim := func(x int) bool {
		return x > 0
	}

	chim(3)

	per := Person{age: 2, name: "s"}

	fmt.Println(per)
	per2 := per
	per2.name = "Jose"
	fmt.Println(per2)

	mapaArr := map[int][]string{1: {"s"}}

	fmt.Println(mapaArr)
	//file,_ := os.Create("te.txt")
	//file.Name()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() {
	path := "/Users/alcaljos/sampleFile.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(path)
	/*check(err)
	  b1 := make([]byte, 6)
	  n1, err := f.Read(b1)
	  check(err)
	  fmt.Printf("%d bytes: %s\n", n1, string(b1))

	  o3, err := f.Seek(6, 0)
	  b3 := make([]byte, 2)
	  n3, err := io.ReadAtLeast(f, b3, 2)
	  check(err)
	  fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	  r4 := bufio.NewReader(f)
	  b4, err := r4.Peek(5)
	  check(err)
	  fmt.Printf("5 bytes: %s\n", string(b4))*/
	fmt.Println("Starign new reading stuff")
	reader := bufio.NewReader(f)
	line, _, _ := reader.ReadLine()
	for line != nil {
		fmt.Println(string(line))
		line, _, _ = reader.ReadLine()
	}
	fmt.Println("Finishing new reading stuff")
}
