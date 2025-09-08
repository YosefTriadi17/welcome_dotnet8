package main

import (
	"fmt"
	"net/http"
)

var counter = 0

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/health", HealthCheck)

	http.ListenAndServe(":8080", nil)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	counter = counter + 1
	if counter > 5 {
		w.WriteHeader(500)
		fmt.Fprintf(w, "KO")
	} else {
		fmt.Fprintf(w, "OK")
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Versi 1.0.1")
}

// Release 366
// changes 1

// git checkout -b feat/new2

// git checkout production

// git commit -am "feat: Tusesday 11 Feb 2025" -m "Penambahan Fitur, Dengan menambahkan langkah ini, Anda dapat memverifikasi bahwa pesan commit telah dikumpulkan dengan benar sebelum melanjutkan ke langkah pembuatan tag dan rilis."

// git commit -am "fix: Tusesday 11 Feb 2025" -m "Perbaikan Fitur, Dengan menambahkan langkah ini, Anda dapat memverifikasi bahwa pesan commit telah dikumpulkan dengan benar sebelum melanjutkan ke langkah pembuatan tag dan rilis."

// git commit -am "chore: Tusesday 11 Feb 2025" -m "Merapikan kode yang amburadul"

// git push --set-upstream origin feat/new2 -f

// git checkout main

// git branch -d feat/new2 -f



// git pull

// git commit -am "fix: revisi github action"

// git push


// git log --pretty=format:"- %s %H%n  %b" $(git merge-base origin/production origin/main)..origin/main --no-merges

// git log --pretty=format:"- %s %H%n  %b" $(git merge-base origin/main origin/production)..origin/production --no-merges

// git log --pretty=format:"- %s %H%n  %b" origin/main..origin/feat/new --no-merges