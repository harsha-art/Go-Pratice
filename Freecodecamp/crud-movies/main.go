package main

import (
"fmt"
"log"
"encoding/json"
"github.com/gorilla/mux"
"strconv"
"math/rand"
"net/http"
)

type Movie struct{
    Id string `json:"Id"`
    Isid string `json:"Isid"`
    title string `json:"title"`
    Director *Director `json:"Director"`
}

type Director struct{
    fname string `json:"fname"`
    lname string `json:"lname"`
}

var movies[] Movie