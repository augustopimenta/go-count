package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
    "encoding/json"
    "os"
    "time"
)

type Response struct {
    Id int
    Count int
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    resp := Response{Id: rand.Int(), Count: 0}

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        resp.Count++

        enc := json.NewEncoder(w)

        if err := enc.Encode(resp); err != nil {
            fmt.Errorf("Ocorreu um erro ao retornar")
            os.Exit(1)
        }
    })

    fmt.Println("Servidor iniciado na porta :80")
    log.Fatal(http.ListenAndServe(":80", nil))
}
