package main

import (
	"fmt"
	"sort"
)

type mobile struct {
	make  string
	model int
	color string
}
type PairList []mobile

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].color < p[j].color }

func main() {

	m := make(PairList, 3)
	m[0] = mobile{
		make:  "Nokia",
		model: 216,
		color: "Black",
	}
	m[1] = mobile{
		make:  "Mi Max",
		model: 2,
		color: "Gold",
	}
	m[2] = mobile{
		make:  "Samsaung",
		model: 7,
		color: "Blue",
	}

	sort.Sort(m)
	//p is sorted

	for _, k := range m {
		fmt.Printf("%v\t%d\t%v\n", k.make, k.model, k.color)
	}

}
