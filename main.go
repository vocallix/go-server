package main

import (
	"fmt"
	"net/http"
)

func HandleHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "전적!\n")
	fmt.Fprintf(w, "1게임 승\n")
	fmt.Fprintf(w, "2게임 승\n")
	fmt.Fprintf(w, "3게임 패 \n")
	fmt.Fprintf(w, "4게임 승\n")
	fmt.Fprintf(w, "5게임 승\n")
}

func HandleFriends(w http.ResponseWriter, r *http.Request) {

	// 데이터는 db에서 가져와야되고

	fmt.Fprintf(w, "친구 목록\n")
	fmt.Fprintf(w, "제갈량\n")
	fmt.Fprintf(w, "하우돈\n")
	fmt.Fprintf(w, "장비\n")
	fmt.Fprintf(w, "여포\n")
}

func main() {
	// 전적 검색
	http.HandleFunc("/history", HandleHistory)

	// 함께하는 친구 검색
	http.HandleFunc("/friends", HandleFriends)

	http.ListenAndServe(":4000", nil)
}
