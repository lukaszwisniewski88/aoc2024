package main

import (
	"bufio"
	"encoding/json"
	"lukaszwisniewski88/aoc2024/first"
	"net/http"
	"strings"
)

func main() {
	// This is a placeholder for the main function of the API.
	// http server with routes
	mux := http.NewServeMux()

	mux.Handle("POST /day/{number}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url_segments := strings.Split(r.URL.String(), "/")
		day_num := url_segments[len(url_segments)-1]
		switch day_num {
		case "1":
			// Process day 1
			scanner := bufio.NewScanner(r.Body)
			input, err := first.GetInputLines(scanner)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pairs, err := first.ProcessLines(input)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			respond := struct {
				SumDiffs        int
				SimilarityScore int
			}{
				SumDiffs:        pairs.GetSumDiffs(),
				SimilarityScore: pairs.GetSimilarityScore(),
			}
			js, err := json.Marshal(respond)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		default:
			http.Error(w, "Day not implemented", http.StatusNotImplemented)
		}
	}))
	// Start the server
	http.ListenAndServe(":8080", mux)
}
