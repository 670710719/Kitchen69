package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Menu struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Type  string  `json:"type"`
}

var menus = []Menu{
	{ID: 1, Name: "ต้มยำกุ้ง", Price: 120, Type: "soup"},
	{ID: 2, Name: "พิซซ่าฮาวายเอี้ยน", Price: 199, Type: "pizza"},
	{ID: 3, Name: "พิซซ่าเห็ด", Price: 179, Type: "pizza"},
}

var nextID = 4

func listMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func getMenu(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "เลขจานต้องเป็นตัวเลข", http.StatusBadRequest)
        return
    }
    for _, m := range menus {
        if m.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(m)
            return
        }
    }
    http.Error(w, "ไม่พบเมนูหมายเลขนี้", http.StatusNotFound)
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /menu", listMenu)
	mux.HandleFunc("GET /menu/{id}", getMenu)
	http.ListenAndServe(":8080", mux)
	mux.HandleFunc("POST /menu", createMenu)
	
}

func createMenu(w http.ResponseWriter, r *http.Request) {
    var m Menu
    if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
        http.Error(w, "อ่านกล่อง JSON ไม่ออก", http.StatusBadRequest)
        return
    }
    if m.Name == "" || m.Price <= 0 {
        http.Error(w, "ต้องมีชื่อเมนู และราคาต้องมากกว่าศูนย์",
            http.StatusBadRequest)
        return
    }
    m.ID = nextID
    nextID++
    menus = append(menus, m)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(m)
}




	








