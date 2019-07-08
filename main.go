package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Bunda Wow super cool
type Bunda interface {
	run()
	validate() bool
}

// BundaImpl lala
type BundaImpl struct {
	isValid bool
}

func (b BundaImpl) validate() bool {
	return !b.isValid
}

func (b BundaImpl) run() {
	fmt.Printf("Running: %+v", b)
}

func (b BundaImpl) additional() {
	fmt.Printf("NOT IN THE INTERFACE")
}

func rodar(b Bunda) {
	if b.validate() {
		b.run()
	} else {
		panic("Bunda falhou")
	}
}

func test() {
	i := 0

	defer fmt.Println(">>> defer", i)
	i++
	defer fmt.Println(">>> defer", i)
	i++
	fmt.Println(">>>", i)
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {

	rodar(BundaImpl{isValid: !true})

	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("REREREcover", r)
		}
	}()
	user1 := User{ID: 1, name: "Will", Email: "will@pm.me"}
	fmt.Printf("Sum: %d\n", sum(2, 2))
	fmt.Printf("Sub: %d\n", sub(2, 2))
	fmt.Printf("Mul: %d\n", mul(2, 2))
	fmt.Printf("Div: %d\n", div(2, 2))
	fmt.Printf("Struct: %+v\n", user1)
	fmt.Printf("Struct (no plus): %v\n", user1)
	fmt.Printf("user1.name: %s\n", user1.name)

	resp, err := http.Get("http://pelota.local:8080")

	if err != nil {
		fmt.Println("error", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	test()
	ret := c()
	fmt.Println(ret)

	var list []string

	// list[0] = "Apple"
	// list[1] = "Banana"
	// list[2] = "Pear"
	// list[3] = "Oak"

	list = append(list, "apple")
	list = append(list, "banana")
	list = append(list, "oak")
	list = append(list, "pear")
	list = append(list, "bunda", "abacate")
	fmt.Println(list[2:3])

	var lista2 [][]string
	lista2 = append(lista2, list)
	fmt.Println(lista2)

	func() { fmt.Println("tanabata") }()

	lala(13)
}

func lala(a int) {
	if a > 5 {
		panic("bigger than five")
	}
}
