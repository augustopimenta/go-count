package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
    "encoding/json"
    "os"
)

type Response struct {
    Id int
    Count int
}

var resp = Response{Id: rand.Int(), Count: 0}

func main() {

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        resp.Count++

        data, err := json.Marshal(resp)

        if err != nil {
            fmt.Errorf("Ocorreu um erro ao retornar")
            os.Exit(1)
        }

        fmt.Fprintf(w, string(data))
    })

    fmt.Println("Servidor iniciado na porta :80")
    log.Fatal(http.ListenAndServe(":80", nil))
}
