package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"WELCOME TO PRAVEEN FIRST WEB APP GO PROGRAM \n\n") //optional
	fmt.Fprintf(w,"THIS WEB APP HAVE ONLY ONE WEB PAGE \n\n")  //optional
	fmt.Fprintf(w,"WEB PAGE WILL DISPLAY CURRENT YEAR, MONTH , DATE AND TIME  \n\n") //optional
	fmt.Fprintf(w,"######################## \n\n ") //optional
	t:= time.Now()
	y := t.Year()
	mo := t.Month()
	d := t.Day()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	fmt.Fprintf(w,"Current Year : %d\n\n",y)
	fmt.Fprintf(w,"Current Month : %d\n\n",mo)
	fmt.Fprintf(w,"Current Date : %d\n\n",d)
	fmt.Fprintf(w,"Current hour : %d\n\n",h)
	fmt.Fprintf(w,"Current minute : %d\n\n",m)
	fmt.Fprintf(w,"Current seconds: %d\n\n",s)
	fmt.Fprintf(w,"######################## \n\n ")
}


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6543", nil)) // this will run on port 6543 , if you run on local system means go to browser type localhost:6543 

}
