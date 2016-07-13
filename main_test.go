package main

import "testing"

func TestAll(t *testing.T) {
	allProv := all.String()
	allProvT := allProverbs()
	if allProv != allProvT {
		t.Fail()
	}
}

func TestRandom(t *testing.T) {
	randomProverbT := randomProverb()
	found := false
	for _, p := range all {
		if p == randomProverbT {
			found = true
			break
		}
	}
	if !found {
		t.Fail()
	}
}
