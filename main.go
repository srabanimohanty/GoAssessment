package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "8000", "Port to listen on")
)

type userDetails struct{
	Name string  `json:"name"`
	Email string  `json:"email"`
	Mobile string  `json:"mobile"`
	Notificationtype int  `json:"notificationtype"`
}

func NotificationSent(ud userDetails) int {
	
	var status int =0
	var err1 error
	var msgType = ud.Notificationtype

	switch msgType {
	case 1:
		status,err1 = SmsSent(ud)
		if err1 !=nil{
			//log data
		}
		break;
	case 2:
		status,err1 = EmailSent(ud)
		if err1 !=nil{
			//log data
		}
		break;
	case 3:
		status,err1 = PhoneCall(ud)
		if err1 !=nil{
			//log data
		}
		break;
	default :
	status,err1 =0,nil

	}

	return status
}

// SmsSent : SMS sent function
func SmsSent(ud userDetails) (status int,msgErr error){

	// sms sent func
	fmt.Println("Welcome to SmsSent")
	 
	return 1, nil
}

// EmailSent : EMAIL sent function
func EmailSent(ud userDetails) (status int,msgErr error){

	// Email sent func
	fmt.Println("Welcome to EmailSent")
 
	return 2, nil
}

// PhoneCall : Message sent in Phone Call function
func PhoneCall(ud userDetails) (status int,msgErr error){

	// phone call func
	fmt.Println("Welcome to PhoneCall")
	 
	return 3, nil
}

var results []string

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" { 
		var statusNotification int 		 
		decoder := json.NewDecoder(req.Body)
		var data userDetails
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
 
		statusNotification =NotificationSent(data)
	
		if statusNotification == 0{
			fmt.Println("error while sending message from:",data.Notificationtype)
		}
		
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))   
		fmt.Fprint(w,statusNotification)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main(){ 

	results = append(results, time.Now().Format(time.RFC3339))

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandler)
	mux.HandleFunc("/notificationSent", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
	
}
