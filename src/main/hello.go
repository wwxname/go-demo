package main

import "fmt"
import "utils"

const v1 = iota
const v2 = iota

const (
	v3 = iota
	v4 = iota
)

var sp [2]struct {
	x int
	y int
}
var i int = 1
var sp1 []int = make([]int, 5, 10)
var sp2 = make(map[string]string)

func GetName() (string, string, string) {
	return "first", "second", "third"
}

func main() {

	sp[0].x = 1
	sp[0].y = 2
	fmt.Println(sp[0].x)
	fmt.Println(GetName())
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(utils.Now())
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错1:", err)

		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错2:", err)
			//utils.MustThrowError()
		}
	}()
	defer func() {
		fmt.Println("出了错3:")
	}()

	utils.MustThrowError()

	fmt.Println(len(sp1))
	fmt.Println(cap(sp1))
	sp2["name"] = "wwx"
	fmt.Println(sp2)
	delete(sp2, "name")
	fmt.Println(sp2)

	var j int = 5
	var a func() = func() func() {
		var i int = 10
		return func() {
			fmt.Println(i)
			fmt.Println(j)
		}
	}()
	fmt.Println(a)
	a()

	j = j * 2
	a()

}
