package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type SeatInfo struct {
// 	class_id     int64
// 	row_count    int64
// 	column_count int64
// }
//
// type SeatingAdjustment struct {
// 	student_id int64
// 	row_no     int64
// 	column_no  int64
// }
//
// type SeatingData struct {
// 	seatInfo *SeatInfo
// 	seatList []*SeatingAdjustment
// }

type SeatInfo struct {
	ClassId     int64 `json:"class_id,string"`
	RowCount    int64 `json:"row_count,string"`
	ColumnCount int64 `json:"column_count,string"`
}

type SeatingAdjustment struct {
	StudentId int64 `json:"student_id,string"`
	RowNo     int64 `json:"row_no,string"`
	ColumnNo  int64 `json:"column_no,string"`
}

type SeatingData struct {
	SeatInfo SeatInfo            `json:"seatInfo"`
	SeatList []SeatingAdjustment `json:"seatList"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	http.HandleFunc("/adjustSeating", func(w http.ResponseWriter, r *http.Request) {
		var seatingData SeatingData
		err := json.NewDecoder(r.Body).Decode(&seatingData)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Parsed data: ")
		fmt.Printf("%#v", seatingData)
	})
	http.ListenAndServe(":8089", nil)
}
