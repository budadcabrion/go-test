package db

import (
	"log"
	"testing"
	"fmt"
)

func init() {
	InitDB("test", true)

	SetScore("alice", 12)
	SetScore("bob", 2)
	SetScore("craig t nelson", 14)
	SetScore("david", 22)
	SetScore("elizabeth", 9)
	SetScore("frank", 0)
	SetScore("gerard", 6)
	SetScore("gerard", 7)
	SetScore("gerard", 8)
	SetScore("gerard", 9)

	fmt.Println("init done")
}

func TestListThings(t *testing.T) {
	scores1 := GetScores(0, 3)
	scores2 := GetScores(1, 1)
	scores3 := GetScores(5, 5)
	scores4 := GetScores(100, 5)

	log.Printf("%v", scores1)
	log.Printf("%v", scores2)
	log.Printf("%v", scores3)
	log.Printf("%v", scores4)

	if len(scores1) != 3 {
		t.Fail()
	}
	if len(scores2) != 1 {
		t.Fail()
	}
	if len(scores3) != 2 {
		t.Fail()
	}
	if len(scores4) != 0 {
		t.Fail()
	}

	if scores1[0].Name != "david" {
		t.Fail()
	}
	if scores2[0].Name != "craig t nelson" {
		t.Fail()
	}
	if scores3[0].Name != "bob" {
		t.Fail()
	}
}