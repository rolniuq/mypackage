package lor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type person struct {
	name string
	age  int
}

var persons = []person{
	{
		name: "1",
		age:  1,
	},
	{
		name: "2",
		age:  2,
	},
	{
		name: "3",
		age:  3,
	},
}

func TestFindFirst(t *testing.T) {
	res := Find(persons, FindFirst)
	require.NotNil(t, res)
	require.True(t, res.age == 1)
	require.True(t, res.name == "1")
}

func TestFindLast(t *testing.T) {
	res := Find(persons, FindLast)
	require.NotNil(t, res)
	require.True(t, res.age == 3)
	require.True(t, res.name == "3")
}

func TestFindWithCondition(t *testing.T) {
	res := FindWithCondition(persons, func(p person) bool {
		return p.age == 2
	})
	require.NotNil(t, res)
	require.True(t, res.age == 2)
}
