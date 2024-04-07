package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"John"},
			[]string{"John"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"John", "Paris"},
			[]string{"John", "Paris"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"John", 36},
			[]string{"John"},
		},
		{
			"struct with nested fields",
			Person{
				"John",
				Profile{Age: 36, City: "Paris"},
			},
			[]string{"John", "Paris"},
		},
		{
			"with a pointer",
			&Person{
				"John",
				Profile{36, "Paris"},
			},
			[]string{"John", "Paris"},
		},
		{
			"with a pointer",
			Person{
				"John",
				Profile{36, "Paris"},
			},
			[]string{"John", "Paris"},
		},
		{
			"with slices",
			[]Profile{
				{21, "Paris"},
				{25, "Rome"},
			},
			[]string{"Paris", "Rome"},
		},
		{
			"with arrays",
			[2]Profile{
				{21, "Paris"},
				{25, "Rome"},
			},
			[]string{"Paris", "Rome"},
		},
		{
			"with map",
			map[string]string{"Hello": "World", "Go": "Lang"},
			[]string{"World", "Lang"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string

			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("%s: expected %s to equal %s", test.Name, actual, test.ExpectedCalls)
			}
		})
	}

}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
