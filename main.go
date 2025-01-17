package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter,r *http.Request){
	
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found", http.StatusNotFound)
		return 
	}
	
	if r.Method != "Get" {
		http.Error(w,"method is not supported", http.StatusNotFound)
		return 

	}


	fmt.Fprintf(w, "hello")

}

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"Parse form error: %v", err)
		return 
	}
	name := r.FormValue("name")

	fmt.Fprintf(w,"Greeting %v", name)
	fmt.Fprintf(w," Post request is successfull")

	


}


func main(){

	fileServer := http.FileServer(http.Dir("./static"));
	

	http.Handle("/",fileServer);

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting server at port %v\n","8080")


	

	if err:=  http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}