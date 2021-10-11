package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	marks := map[string]int{
		"English":         88,
		"Science":         90,
		"Maths":           97,
		"Social sciences": 82,
	}

	p := make(PairList, len(marks))

	i := 0
	for k, v := range marks {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted

	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}

}
