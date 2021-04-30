package main

import "testing"

func AddThreeNumbersTest(t *testing.T) {
    n1 := 3
    n2 := 4
    n3 := 5
    got := AddThreeNumbers(n1, n2, n3)
    want := 12

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

