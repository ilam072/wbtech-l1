package main

import "fmt"

func main() {
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}, 7: {}, 8: {}}
	set2 := map[int]struct{}{4: {}, 5: {}, 3: {}, 7: {}, 9: {}}

	in := intersection(set1, set2)
	fmt.Println(in)
}

func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	res := make(map[int]struct{})
	for v := range set1 {
		if _, ok := set2[v]; ok {
			res[v] = struct{}{}
		}
	}

	return res
}
