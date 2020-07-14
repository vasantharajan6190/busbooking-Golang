package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
    "../models"
	"../drivers"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
)

//users
//login
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var sum = models
	var users []models.user
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", params["name"]))
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)
	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2, queryJs)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	searchService := esclient.Search().Index("users").SearchSource(searchSource)
	ctx := context.Background()
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	if len(searchResult.Hits.Hits) == 0 {
		fmt.Println("")
		json.NewEncoder(w).Encode("Empty")
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		var ss user
		err := json.Unmarshal(hit.Source, &ss)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		users = append(users, ss)
	}
	fmt.Println("Users", users)
	json.NewEncoder(w).Encode(users[0])
}

//user registering
func addusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var use model.user
	_ = json.NewDecoder(r.Body).Decode(&use)
	ctx := context.Background()
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	dataJSON, err := json.Marshal(use)
	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("users").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful", ind)
	json.NewEncoder(w).Encode(use)
}

//travels
func gettravels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var travels []model.travel
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "travels"))
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)
	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2, queryJs)
	}
	//fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	searchService := esclient.Search().Index("travels").SearchSource(searchSource)
	ctx := context.Background()
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	if len(searchResult.Hits.Hits) == 0 {
		fmt.Println("")
		json.NewEncoder(w).Encode("Empty")
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		var ss travel
		err := json.Unmarshal(hit.Source, &ss)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		travels = append(travels, ss)
	}
	fmt.Println("students", travels)
	json.NewEncoder(w).Encode(travels)
}

//adding new buses
func addtravels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var use model.travel
	_ = json.NewDecoder(r.Body).Decode(&use)
	ctx := context.Background()
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	dataJSON, err := json.Marshal(use)
	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("travels").
		BodyJson(js).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful", ind)
	json.NewEncoder(w).Encode(use)
}

//bookings
func getallbookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var bookings []model.booking
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	searchService := esclient.Search().Index("bookings")
	ctx := context.Background()
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	if len(searchResult.Hits.Hits) == 0 {
		fmt.Println("")
		json.NewEncoder(w).Encode("Empty")
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		var ss booking
		err := json.Unmarshal(hit.Source, &ss)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		bookings = append(bookings, ss)
	}
	json.NewEncoder(w).Encode(bookings)
}

func getbookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var bookings []model.booking
	params := mux.Vars(r)
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", params["name"]))
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)
	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2, queryJs)
	}
	//fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	searchService := esclient.Search().Index("bookings").SearchSource(searchSource)
	ctx := context.Background()
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	if len(searchResult.Hits.Hits) == 0 {
		fmt.Println("")
		json.NewEncoder(w).Encode("Empty")
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		var ss booking
		err := json.Unmarshal(hit.Source, &ss)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		bookings = append(bookings, ss)
	}
	fmt.Println("bookings", params["name"])
	length := len(bookings)
	json.NewEncoder(w).Encode(bookings[length-1])
}
func addbookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var x model.booking
	_ = json.NewDecoder(r.Body).Decode(&x)
	ctx := context.Background()
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	dataJSON, err := json.Marshal(x)
	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("bookings").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful", ind)
	json.NewEncoder(w).Encode(x)
}
func updatebookings() {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var x model.booking
	_ = json.NewDecoder(r.Body).Decode(&x)
	ctx := context.Background()
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	dataJSON, err := json.Marshal(x)
	js := string(dataJSON)
	ind, err := esclient.Update().
		Index("bookings").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful", ind)
	json.NewEncoder(w).Encode(x)
}

//travels
func gettravelsbyname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var bookings []model.travel
	params := mux.Vars(r)
	esclient, err := drivers.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", params["name"]))
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)
	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2, queryJs)
	}
	//fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	searchService := esclient.Search().Index("travels").SearchSource(searchSource)
	ctx := context.Background()
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	if len(searchResult.Hits.Hits) == 0 {
		fmt.Println("")
		json.NewEncoder(w).Encode("Empty")
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		var ss travel
		err := json.Unmarshal(hit.Source, &ss)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		bookings = append(bookings, ss)
	}
	fmt.Println("bookings", bookings)
	// length := len(bookings)
	json.NewEncoder(w).Encode(bookings)
}
