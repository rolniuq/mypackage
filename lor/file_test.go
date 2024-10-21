package lor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Name string
	Age  int
}

func TestReadJsonFile(t *testing.T) {
	os.WriteFile("input.json", []byte(`{"name": "John", "age": 30}`), 0644)

	res, err := ReadJsonFile[Person]("input.json")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Name == "John" && res.Age == 30)

	os.Remove("input.json")
}

func TestWriteJsonFile(t *testing.T) {
	person := Person{
		Name: "John",
		Age:  30,
	}

	err := WriteJsonFile("output.json", person)
	require.NoError(t, err)

	os.Remove("output.json")
}
