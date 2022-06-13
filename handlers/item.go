package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"josephbartley.dev/go-api-test/utils"
)

type Item struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

/*
Gets all items from the database that have not been marked as deleted in Id order
*/
func GetAllItemsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(404)
	}

	dat := utils.ReadFile("./data/data.json")
	var items map[string]Item
	unmarshalErr := json.Unmarshal(dat, &items)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	var results []Item
	for _, item := range items {
		if item.DeletedAt == "" {
			results = append(results, item)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Id < results[j].Id
	})

	data, marshalErr := json.Marshal(results)
	if marshalErr != nil {
		panic(marshalErr)
	}

	fmt.Fprint(w, string(data))
}

func ItemRouter(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		GetItemHandler(w, req)
		return
	} else if req.Method == "POST" {
		AddItemHandler(w, req)
		return
	} else if req.Method == "DELETE" {
		DeleteItemHandler(w, req)
		return
	}
}

func GetItemHandler(w http.ResponseWriter, req *http.Request) {
	dat := utils.ReadFile("./data/data.json")
	var items map[string]Item
	unmarshalErr := json.Unmarshal(dat, &items)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	id := strings.TrimPrefix(req.URL.Path, "/item/")
	item, ok := items[id]
	if !ok {
		w.WriteHeader(404)
		fmt.Fprint(w, string("{\"message\": \"Item not found\"}"))
		return
	}

	data, marshalErr := json.Marshal(item)
	if marshalErr != nil {
		panic(marshalErr)
	}
	fmt.Fprint(w, string(data))
}

func AddItemHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		panic(readErr)
	}

	var item Item
	unmarshalErr0 := json.Unmarshal(body, &item)
	if unmarshalErr0 != nil {
		panic(unmarshalErr0)
	}

	dat := utils.ReadFile("./data/data.json")
	var items map[int]Item
	unmarshalErr := json.Unmarshal(dat, &items)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	id := len(items) + 1
	item.Id = id
	items[id] = item

	itemsJson, marshalErr := json.Marshal(items)
	if marshalErr != nil {
		panic(marshalErr)
	}
	utils.WriteFile("./data/data.json", itemsJson)

	w.WriteHeader(200)
	fmt.Fprint(w, string("{\"message\": \"Item added\"}"))
}

func DeleteItemHandler(w http.ResponseWriter, req *http.Request) {
	dat := utils.ReadFile("./data/data.json")
	var items map[string]Item
	unmarshalErr := json.Unmarshal(dat, &items)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	id := strings.TrimPrefix(req.URL.Path, "/item/")
	item, ok := items[id]
	if !ok {
		w.WriteHeader(404)
		fmt.Fprint(w, string("{\"message\": \"Item not found\"}"))
		return
	}

	if item.DeletedAt != "" {
		w.WriteHeader(404)
		fmt.Fprint(w, string("{\"message\": \"Item not found or already deleted\"}"))
		return
	}

	item.DeletedAt = time.Now().Format("2006-02-01 15:04:05")
	items[id] = item

	itemsJson, marshalErr := json.Marshal(items)
	if marshalErr != nil {
		panic(marshalErr)
	}
	utils.WriteFile("./data/data.json", itemsJson)

	w.WriteHeader(200)
	fmt.Fprint(w, string("{\"message\": \"Item deleted\"}"))
}
