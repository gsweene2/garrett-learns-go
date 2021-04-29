package main

import "testing"

func ConcatStringsTest(t *testing.T) {
    s1 := "Concat"
    s2 := "enation"
    got := ConcatStrings(s1,s2)
    want := "Concatenation"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

