package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://apidojo-forever21-v1.p.rapidapi.com/products/v2/detail?productId=2000383198"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "1aa5b6fa64msh21e535afc50d574p1e0eadjsnd37ae1fe22f2")
	req.Header.Add("X-RapidAPI-Host", "apidojo-forever21-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}