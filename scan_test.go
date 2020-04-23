package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	assert.Equal(t, []int{}, Left([]int{1, 2}, 0))
	assert.Equal(t, []int{1}, Left([]int{1, 2}, 1))
	assert.Equal(t, []int{1, 2}, Left([]int{1, 2}, 2))
	assert.Equal(t, []int{1, 2}, Left([]int{1, 2}, 3))
}

func TestLen(t *testing.T) {
	assert.Equal(t, 0, Len([]int{}))
	assert.Equal(t, 1, Len([]int{1}))
	assert.Equal(t, 1, Len([1]int{1}))
	assert.Equal(t, 2, Len([2]int{1, 2}))
	assert.Equal(t, 0, Len(map[string]int{}))
	assert.Equal(t, 1, Len(map[string]int{"one": 1}))
}

func TestForEach(t *testing.T) {
	is := assert.New(t)

	var results []int

	ForEach([]int{1, 2, 3, 4}, func(x int) {
		if x%2 == 0 {
			results = append(results, x)
		}
	})

	is.Equal(results, []int{2, 4})

	results = results[0:0]

	ForEach([]int{1, 2, 3, 4}, func(i, x int) {
		if i%2 == 0 {
			results = append(results, x)
		}
	})

	is.Equal(results, []int{1, 3})

	results = results[0:0]

	ForEach([]int{1, 2, 3, 4}, func(i, x int) bool {
		results = append(results, x)
		return i < 2
	})

	is.Equal(results, []int{1, 2, 3})

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	ForEach(mapping, func(k int, v string) {
		is.Equal(v, mapping[k])
	})

	j := 0

	ForEach(mapping, func(i, k int, v string) {
		is.Equal(v, mapping[k])
		is.Equal(i, j)
		j++
	})
}

func TestForEachRight(t *testing.T) {
	is := assert.New(t)

	results := []int{}

	ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		results = append(results, x*2)
	})

	is.Equal(results, []int{8, 6, 4, 2})

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	mapKeys := []int{}

	ForEachRight(mapping, func(k int, v string) {
		is.Equal(v, mapping[k])
		mapKeys = append(mapKeys, k)
	})

	is.Equal(len(mapKeys), 2)
	is.Contains(mapKeys, 1)
	is.Contains(mapKeys, 2)
}

func TestHead(t *testing.T) {
	is := assert.New(t)

	is.Equal(Head([]int{1, 2, 3, 4}), 1)
}

func TestLast(t *testing.T) {
	is := assert.New(t)

	is.Equal(Last([]int{1, 2, 3, 4}), 4)
}

func TestTail(t *testing.T) {
	is := assert.New(t)

	is.Equal(Tail([]int{}), []int{})
	is.Equal(Tail([]int{1}), []int{1})
	is.Equal(Tail([]int{1, 2, 3, 4}), []int{2, 3, 4})
}

func TestInitial(t *testing.T) {
	is := assert.New(t)

	is.Equal(Initial([]int{}), []int{})
	is.Equal(Initial([]int{1}), []int{1})
	is.Equal(Initial([]int{1, 2, 3, 4}), []int{1, 2, 3})
}
