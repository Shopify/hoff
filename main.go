package main

import (
	"fmt"

	"github.com/Shopify/hof/set"
)

func main() {
	setA := set.FromSlice([]int{1, 2, 3, 1, 2, 3})
	setB := set.FromSlice([]int{3, 4, 5})

	// remember, sets are unordered, so the output might be too!
	fmt.Println("setA", setA.Values())
	fmt.Println("setB", setB.Values())

	fmt.Println("A.Len() == len(A)", setA.Len(), len(setA))
	fmt.Println("A.Has(1), A.Has(999)", setA.Has(1), setA.Has(999))

	// add 8, 9, 10 to A (1, 2, 3, 8, 9, 10)
	setA.Add([]int{8, 9, 10}...)
	fmt.Println("Added 8, 9, 10", setA.Values())

	// delete 8, 9, 10 from A (1, 2, 3)
	setA.Delete([]int{8, 9, 10}...)
	fmt.Println("Removed 8, 9, 10", setA.Values())

	// combination of both sets (1, 2, 3, 4, 5)
	fmt.Println("Union A ∪ B", setA.Union(setB).Values())

	// things that exist in both sets (3)
	fmt.Println("Intersection A ∩ B", setA.Intersection(setB).Values())

	// produces the things in setA that are not in setB (1, 2)
	fmt.Println("Difference A-B", setA.Difference(setB).Values())
	// and vice versa (4, 5)
	fmt.Println("Difference B-A", setB.Difference(setA).Values())

	// things that are _not_ in both sets - opposite of Intersection (1, 2, 4, 5)
	fmt.Println("Symm Diff A △ B", setA.SymmetricalDifference(setB).Values())

	// Clone A as C, edit C, C != A
	setC := setA.Clone()
	setC.Add(999)
	fmt.Println("A, C", setA.Values(), setC.Values())

	// run a callback on each element of A, multiplying by 10
	setX := make(set.Set[int], len(setA))
	setA.ForEach(
		func(v int) {
			setX.Add(v * 10)
		},
	)
	fmt.Println("A values * 10", setX.Values())
}
