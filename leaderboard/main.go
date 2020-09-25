package main

import "fmt"
import "net/http"
import "strconv"

var leaderboard *Leaderboard

func setScore(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if name == "" {
		http.Error(w, "name must be present", 503)
		return
	}
	scoreString := req.FormValue("score")
	score, err := strconv.ParseInt(scoreString, 0, 64)
	if err != nil {
		http.Error(w, "score must be present and an integer", 503)
		return
	}
	leaderboard.SetScore(name, score)
	fmt.Fprintf(w, "confirmed %v = %d\n", name, score)
}

func getScores(w http.ResponseWriter, req *http.Request) {

	startString := req.FormValue("start")
	start, err := strconv.ParseInt(startString, 0, 64)
	if err != nil {
		http.Error(w, "start must be present and an integer", 503)
		return
	}
	countString := req.FormValue("count")
	count, err := strconv.ParseInt(countString, 0, 64)
	if err != nil {
		http.Error(w, "count must be present and an integer", 503)
		return
	}

	scores := leaderboard.GetScores(int(start), int(count))

	for _, s := range scores {
		fmt.Fprintf(w,"%v %d\n", s.Name, s.Score)
	}
}

func main() {
	leaderboard = MakeLeaderboard()

	http.HandleFunc("/setscore", setScore)
	http.HandleFunc("/getscores", getScores)

	http.ListenAndServe(":1234", nil)
}
