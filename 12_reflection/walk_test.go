package reflection

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("First Case, simple value", func(t *testing.T) {
		expected := "Someone"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}

		if got[0] != expected {
			t.Errorf("got %q want %q", got[0], expected)
		}
	})

	t.Run("Second Case, more advanced value", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Someone"},
				[]string{"Someone"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Someone", "Earth"},
				[]string{"Someone", "Earth"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Someone", 33},
				[]string{"Someone"},
			},
			{
				"nested fields",
				Person{
					"Someone",
					Profile{30, "Earth"},
				},
				[]string{"Someone", "Earth"},
			},
			{
				"pointers to things",
				&Person{
					"Someone",
					Profile{30, "Earth"},
				},
				[]string{"Someone", "Earth"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			// TODO: Handle Arrays
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string

				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

}
