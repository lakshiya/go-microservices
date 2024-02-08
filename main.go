package main
// Application that acts as a HTTP Server
import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "time"
    "log"
    "fmt"
    details "github.com/lakshiya/go-microservices/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Checking application health")
    response := map[string]string{
        "status" : "UP",
        "timestamp" : time.Now().String(),
    }
    json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Serving the homepage")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Fetching the details")
    hostName, err := details.GetHostName()
    if err != nil {
        panic(err)
    }
    IP, _ := details.GetIP()
    fmt.Println(hostName, IP)
    response := map[string]string{
        "hostname" : hostName,
        "IP" : IP.String(),
    }
    json.NewEncoder(w).Encode(response)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/health", healthHandler)
    r.HandleFunc("/", rootHandler)
    r.HandleFunc("/details", detailsHandler)
    log.Println("Server has started!!!")
    log.Fatal(http.ListenAndServe(":80", r))
}