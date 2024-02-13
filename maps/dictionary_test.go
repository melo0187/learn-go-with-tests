package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrDefinitionNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	assertNoError(t, err)
	assertStrings(t, got, want)
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Did get an error but did not want one")
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Did not get error but wanted one")
	}

	if err != want {
		t.Errorf("got error %q, want %q", err, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
