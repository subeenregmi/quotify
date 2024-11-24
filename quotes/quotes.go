package quotes

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Response struct {
    Quote   string  `json:"q"`
    Author  string  `json:"a"`
    Length  int64   `json:"c"`
}

const QuoteAPI      = "https://zenquotes.io/api/random/"
const CSVFile       = "quotes/quotes.csv"

func GetRandomQuote() (*Response, error) {
    resp, err := http.Get(QuoteAPI)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    var qs []Response
    err = json.Unmarshal([]byte(body), &qs)

    if err != nil {
        return nil, err
    }

    return &qs[0], nil
}

func GetCSVQuotes() (*[]Response, error) {
    file, err := os.Open(CSVFile)

    if err != nil {
        return nil, err
    }

    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()

    if err != nil {
        return nil, err
    }

    responses := []Response{}

    for _, rec := range records {
        v := Response{Quote: rec[1], Author: rec[0]}
        responses = append(responses, v)
    }

    return &responses, nil
}
