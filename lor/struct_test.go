package lor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Person1 struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
}

func TestStructToStruct(t *testing.T) {
	p1 := Person1{
		Name: "John",
		Age:  30,
	}

	res, err := StructToStruct[Person1, Person2](&p1)
	require.NoError(t, err)
	require.True(t, res.Name == "John")
}

type a1 struct {
	CampaignId string
}

type a2 struct {
	CampaignId string
	Name       string
}

type a3 struct {
	CampaignId string
	Name       string
	Age        int
}

func TestMapStructs(t *testing.T) {
	a1 := a1{CampaignId: "123"}

	arr := []interface{}{a2{}, a3{}}

	res, err := MapStructs(a1, arr)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, len(res) == 2)

	a2V, err := StructToStruct[interface{}, a2](&res[0])
	require.NoError(t, err)
	require.True(t, a2V.CampaignId == "123")

	a3V, err := StructToStruct[interface{}, a3](&res[1])
	require.NoError(t, err)
	require.True(t, a3V.CampaignId == "123")
}
