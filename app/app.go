package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sr/quotify/quotes"
)

type QuotesCSV struct {
    Quotes []quotes.Response
}

type QuoteLike struct {
    Quote string    `json:"quote"`
    Likes int       `json:"likes"`
}

func (q *QuotesCSV) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    i := rand.Intn(len(q.Quotes))

    jsonB, err := json.Marshal(q.Quotes[i])
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }

    w.Write(jsonB)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
    quote, err := quotes.GetRandomQuote()
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }
    jsonB, err := json.Marshal(quote)
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }
    w.Write(jsonB)
}

func AddLike(w http.ResponseWriter, r *http.Request) {

    body, err := io.ReadAll(r.Body)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }
    defer r.Body.Close()

    var q QuoteLike
    err = json.Unmarshal(body, &q)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }
    
    likes[q.Quote]++
    q.Likes = likes[q.Quote]

    jsonB, err := json.Marshal(q)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }

    w.Write(jsonB)
}

func ViewLikes(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }

    defer r.Body.Close()

    var q QuoteLike
    err = json.Unmarshal(body, &q)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }

    q.Likes = likes[q.Quote]
    fmt.Println(q.Likes)

    jsonB, err := json.Marshal(q)

    if err != nil {
        fmt.Fprintln(w, err)
        return 
    }

    w.Write(jsonB)
}

var likes = make(map [string]int)

func RunServer() {

    var qsCSV QuotesCSV 
    qs, err := quotes.GetCSVQuotes()

    if err != nil {
        log.Fatal(err)
    }

    qsCSV.Quotes = *qs

    http.HandleFunc("/api/like/add", AddLike)
    http.HandleFunc("/api/like/view", ViewLikes)
    http.Handle("/api/quote", &qsCSV)
    http.Handle("/", http.FileServer(http.Dir("static/")))
    http.ListenAndServe(":8989", nil)
}
