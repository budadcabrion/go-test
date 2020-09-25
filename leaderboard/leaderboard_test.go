package main

import (
	"testing"
	"log"
)

var l *Leaderboard

func init() {
	l = MakeLeaderboard()
	l.SetScore("alice", 12)
	l.SetScore("bob", 2)
	l.SetScore("craig t nelson", 14)
	l.SetScore("david", 22)
	l.SetScore("elizabeth", 9)
	l.SetScore("frank", 0)
	l.SetScore("gerard", 9)
}

func TestLeaderboardSorting(t *testing.T) {
	l.updateSortedScores()

	log.Printf("%v", l.sortedScores)

	for i := 0; i < len(l.scores)-1; i++ {
		if l.sortedScores[i].Score < l.sortedScores[i+1].Score {
			t.Errorf("not actually sorted %d %d %d", i, l.sortedScores[i].Score, l.sortedScores[i+1].Score)
		}
		if l.sortedScores[i].Score == l.sortedScores[i+1].Score {
			if l.sortedScores[i].Name < l.sortedScores[i+1].Name {
				t.Errorf("ties should sort by name ascending %d %v %v", i, l.sortedScores[i].Name, l.sortedScores[i+1].Name)
			}
		}
	}
}


func TestGetScores(t *testing.T) {
	scores1 := l.GetScores(0, 3)
	scores2 := l.GetScores(1, 1)
	scores3 := l.GetScores(5, 5)
	scores4 := l.GetScores(100, 5)

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