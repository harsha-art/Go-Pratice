package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandlefunc(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() err : %v", err)
		return 
	}
	fmt.Fprintf(w,"Form Parse Successful !!")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Fprintf(w,"Your Name is %s %s ",fname,lname)
}

func hellohandlefunc(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"method not support",http.StatusNotFound)
		return
	}
	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello ")
}

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formhandlefunc)
	http.HandleFunc("/hello",hellohandlefunc)
	fmt.Println("Server being set up at Port 8000")
	if err := http.ListenAndServe(":8000",nil);err != nil{
		log.Fatal(err)
	}
}