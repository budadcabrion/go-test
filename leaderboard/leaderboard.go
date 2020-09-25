package main

import (
	"sort"
)

type PlayerScore struct {
	Name string
	Score int64
}

type Leaderboard struct {
	scores map[string]int64
	sortedScores []PlayerScore
	sorted bool
}

func MakeLeaderboard() *Leaderboard {
	var l Leaderboard
	l.scores = make(map[string]int64)
	return &l
}

func (l *Leaderboard) SetScore(name string, score int64) {
	l.scores[name] = score
	l.sorted = false
}

func (l* Leaderboard) GetScore(name string) (score int64) {
	score, _ = l.scores[name]
	return
}

func (l* Leaderboard) updateSortedScores() {
	s := make([]PlayerScore, len(l.scores))
	i := 0
	for k, v := range l.scores {
		s[i] = PlayerScore{k, v}
		i = i+1
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].Score == s[j].Score {
			return s[i].Name > s[j].Name
		}
		return s[i].Score > s[j].Score
	})

	l.sortedScores = s
	l.sorted = true
}

func (l* Leaderboard) GetScores(start int, count int) []PlayerScore {
	if start >= len(l.scores) {
		return []PlayerScore{}
	}
	if !l.sorted {
		l.updateSortedScores()
	}
	end := start+count
	if end > len(l.scores) {
		end = len(l.scores)
	}
	return l.sortedScores[start:end]
}