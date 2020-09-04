package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WiseSaying struct
type WiseSaying struct {
	Saying string `json:"saying"`
	Name   string `json:"name"`
}

func main() {
	db := []WiseSaying{}
	db = append(db, WiseSaying{"모든 사업은 계속해서 젊어져야 한다. 고객이 당신과 함께 늙어간다면 당신은 Woolworth가 될 것이다.", "제프 베조스"})
	db = append(db, WiseSaying{"실패는 또 하나의 옵션이다. 실패를 겪지 않는다면 충분히 혁신적이지 않았다는 증거이다.", "일론 머스크"})
	db = append(db, WiseSaying{"정말 중요한 일이라면 역경이 닥쳐도 그 일을 계속해야 하는 것이다.", "일론 머스크"})

	http.HandleFunc("/api/wise-sayings", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.Encode(&db)
	})

	fmt.Println("server is running http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
