package main

import(
    "net/http"
    "github.com/gorilla/mux"
)

func main(){
    r := mux.NewRouter()
    
    r.Handle("/", http.FileServer(http.Dir("./views/")))
    
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    
    http.ListenAndServe(":3000", r)
    
    
    r.Handle("/status", NotImplemented).Methods("GET")
    r.Handle("/products", NotImplemented).Methods("GET")
    r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Not Implemented"))
})

