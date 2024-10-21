package lor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Bob", 40},
	}

	names := Map(people, func(p Person) string {
		return p.Name
	})
	require.True(t, names[0] == "John" && names[1] == "Jane" && names[2] == "Bob")
}

func TestFilter(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Bob", 40},
	}

	adults := Filter(people, func(p Person) bool {
		return p.Age >= 30
	})

	require.True(t, len(adults) == 2 && adults[0].Name == "John" && adults[1].Name == "Bob")
}

func TestReduce(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Bob", 40},
	}

	sum := Reduce(people, 0, func(acc int, p Person) int {
		return acc + p.Age
	})

	require.True(t, sum == 95)
}
