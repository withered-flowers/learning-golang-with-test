package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
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
			{
				"arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Berlin"}
			aChannel <- Profile{34, "Java"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Java"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{20, "Hello"}, Profile{18, "Almost"}
		}

		var got []string
		want := []string{"Hello", "Almost"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
