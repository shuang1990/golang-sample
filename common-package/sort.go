package main

import (
	"sort"
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age int8
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (a ByAge) Less(i, j int) bool {
	return a[i].Age > a[j].Age
}
// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var persons ByAge = []Person{
		{
			Name: "zhangsan",
			Age: 27,
		},
		{
			Name: "lisi",
			Age: 20,
		},
		{
			Name: "wangwu",
			Age: 26,
		},
		{
			Name: "zhaoliu",
			Age: 30,
		},
	}
	sort.Sort(persons)

	for _, val := range persons {
		fmt.Printf("%+#v\n", val)
	}

	rand.Seed(time.Now().UnixNano())
	intList := make([]int, 10)
	for i := 0; i < 10; i++ {
		data := rand.Intn(100)
		intList[i] = data
	}

	fmt.Println("ints sort before: ", intList)
	sort.Ints(intList)
	fmt.Println("ints sort after: ", intList)

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println("ints reverse after: ", intList)

	floatList := []float64{12.4, 34.3, 56.4, 17.4, 18.8, 20.8}
	fmt.Println("floats sort before: ", floatList)
	sort.Float64s(floatList)
	fmt.Println("floats sort before: ", floatList)

	sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))
	fmt.Println("floats reverse after: ", floatList)

	stringList := []string{"smith", "jack", "wilson", "lily", "green", "tom"}
	fmt.Println("strings sort before: ", stringList)
	sort.Strings(stringList)
	fmt.Println("strings sort before: ", stringList)

	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println("strings reverse after: ", stringList)
}
