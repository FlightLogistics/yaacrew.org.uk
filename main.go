package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"yaacrew.org.uk/datecheck"
)

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", hello1)
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":3000", nil)
}

func hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my new awesome site! THIS IS GREAT</h1>")

	testDate := time.Now()

	var d1, d2 datecheck.ExpiryDate
	d1 = datecheck.GetExpiryDate(testDate, 6, datecheck.Month, true)
	d2 = datecheck.GetExpiryDate(testDate, 3, datecheck.Day, true)

	fmt.Fprint(w, "<p>Date of test: "+testDate.String()+"</p>")
	fmt.Fprint(w, "<p>Expiry Date (6 months to end of month): "+datecheck.ConvertExpiryDateToString(d1)+"</p>")
	fmt.Fprint(w, "<p>Expiry Date (3 days): "+datecheck.ConvertExpiryDateToString(d2)+"</p>")
	return
}
