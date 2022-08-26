package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"www.github.com/KangkanaBaruah/finalProjectJosh/handler"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			handler.Getstatus(w)
		}

	case "POST":
		{
			handler.QParam = r.FormValue("name")
			if handler.QParam == "" {
				handler.Postall(w, r)
			} else {
				handler.Postquerry(w, handler.QParam)
			}
		}
	default:
		fmt.Println("Unexpected command")
	}
}
func PrintInMinute(mapData map[string]string) {
	for {
		if handler.Siteslist.SitesName != nil {
			handler.CheckStatus()
			for site, status := range mapData {
				fmt.Printf("%s : %s\n", site, status)
			}
			time.Sleep(6 * time.Second)
		}
	}
}
func main() {
	fmt.Println("Starting server...")
	go PrintInMinute(handler.Webstatus)
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
