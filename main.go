package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func getSwapiURL(endpoint, searchTerm string) ([]byte, int) {
    swapiURL := fmt.Sprintf("https://swapi.dev/api/%s/", endpoint)
    if searchTerm != "" {
        swapiURL += fmt.Sprintf("?search=%s", searchTerm)
    }

    resp, err := http.Get(swapiURL)
    if err != nil {
        log.Printf("Failed to fetch data from SWAPI: %v", err)
        return nil, http.StatusInternalServerError
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Failed to read response body: %v", err)
        return nil, http.StatusInternalServerError
    }

    return body, resp.StatusCode
}

func getStarships(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("search")
    response, status := getSwapiURL("starships", searchTerm)

    w.WriteHeader(status)
    w.Write(response)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("search")
    response, status := getSwapiURL("people", searchTerm)

    w.WriteHeader(status)
    w.Write(response)
}

func home(w http.ResponseWriter, r *http.Request) {
    apiList := map[string]interface{}{
        "_welcome_note": "Star Wars API Wrapper from SWAPI. Use the available APIs to access information about Star Wars starships and characters.",
        "available_apis": map[string]string{
            "/ships":       "Access information about Star Wars starships. Use the 'search' query parameter to filter results.",
            "/characters":  "Access information about Star Wars characters. Use the 'search' query parameter to filter results.",
            "/":            "Home page listing all available APIs.",
            "search_query": "Use the 'search' query parameter to filter results based on the name of the starship or character.",
        },
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(apiList)
}

func catchAll(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Not active API"})
}

func main() {
    http.HandleFunc("/ships", getStarships)
    http.HandleFunc("/characters", getPeople)
    http.HandleFunc("/", home)
    http.HandleFunc("/<path:path>", catchAll)

    fmt.Print("Listening on port 5002\n")
    log.Fatal(http.ListenAndServe(":5002", nil))
}
