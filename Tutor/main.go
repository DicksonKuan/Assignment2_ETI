// package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Tutor struct {
	TutorID   int    `json: "TutorID"`
	FirstName string `json: "FirstName"`
	LastName  string `json: "LastName"`
	Email     string `json: "Email"`
	password  string `json: "Password"`
}

//Database
func database() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3308)/rideshare") //Connecting to database
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db
}

//Key
func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//Microservice functions
func checkMicroservices() {
	//To check if microservice is online
	url := [6]string{
		"http://localhost:5000/api/v1/Tutor/",
		"http://localhost:5000/api/v1/Modules/",
		"http://localhost:5000/api/v1/Class/",
		"http://localhost:5000/api/v1/Student/",
		"http://localhost:5000/api/v1/RatingAndComments/",
		"http://localhost:5000/api/v1/Timetable/"
	}
	type = := [5]string{"Tutor","Modules","Class","Student","RatingAndComments","Timetable"}
	for i, s:= range url{
		response, err := http.Get(s)
		if err == nil{
			fmt.println(fmt.Sprintf("'%s' is working"))
		}else{
			fmt.println(fmt.Sprintf("'%s' is not working"))
		}
	}
}

func getTutor(email string) string {
	url := "http://localhost:5000/api/v1/tutor/" + email
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
			return nil
		} else {
			var tutor Tutor
			err = json.Unmarshal([]byte(responseData), &tutor)
			return responseData
		}
	}
	return ""
}

func checkTutorExsist(email string) bool {
	//To check if tutor exsists and information is accurate
	url := "http://localhost:5000/api/v1/tutor/checkTutor/" + email
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
			return false
		} else {
			var tutor Tutor
			err = json.Unmarshal([]byte(responseData), &tutor)
			if err == nil || tutor.Email == email{
				return true
			}else{
				return false
			}
		}
	}
	return false

}

func putUser(tutor Tutor) bool { //Update tutor's profile
	jsonValue, _ := json.Marshal(tutor)
	URL := "http://localhost:5000/api/v1/CheckUser/" + tutorEmail
	
	request, err := http.NewRequest(http.MethodPut,
		URL,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return false
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		return true
		response.Body.Close()
	}
}
func getMod(tutorEmail string) []string { //get mod from mod microservice
	URL := "http://localhost:5000/api/v1/CheckUser/" + tutorEmail

	response, err := http.Get(URL)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	} else if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
		} else {
			mods := strings.Split(string(responseData), ",")
			replacer := strings.NewReplacer(",", "")
			var newMods []string
			for i := range mods {
				newMods = append(newMods, replacer.Replace(mods[i]))
			}
			return newMods
		}
	}
	return nil
}

//API Functions
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: Tutor API Is working")
}
func profile(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}
	if r.Header.Get("Content-type") == "application/json" {
		//Get information from JSON and validation
		var tutor Tutor
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil { //To check if parameters are empty
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Please provide Email or password"))
			return
		} else {
			json.Unmarshal(reqBody, &tutor)
			if tutor.Email == " " { //To check for information not empty
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Please supply tutor's email"))
				return
			}
			if !checkTutorExsist(tutor.Email) { //To check if tutor exsists in the DB
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("There is no exsiting account for " + tutor.Email))
				return
			}
		}
		//Check method
		if r.Method == "GET" {
			//To get tutor's profile
			tutor = getTutor(tutor.Email)
			if tutor == ""{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Could not retrieve tutor"))
			}else{
				json.NewEncoder(w).Encode(tutor)
				w.WriteHeader(http.StatusAccepted)
			}
			return
		} else if r.Method == "PUT" { //To update tutor's profile
			if tutor.password == "" { //Check if password is empty
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please enter password"))
				return
			} else {
				putUser(tutor) //Update tutor's profile
				w.WriteHeader(http.StatusAccepted)
				return
			}
		}
	} else {
		w.WriteHeader(
			http.StatusUnprocessableEntity)
		w.Write([]byte(
			"Please supply tutor's information"))
		return
	}
}
func mod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	method := params["method"]
	email := params["tutorEmail"]

	if method == "" || email == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Please supply tutor's information and valid method"))
		return
	} else {
		switch method {
		case "getMod":
			mods := getMod(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Mod list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getClassAssigned":
			classes := getMod(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"class list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getTimetable":
			timetable := getMod(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"timetable list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		case "enrolledStudent":
			students := getMod(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"timetable list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
}
func details(w http.ResponseWriter, r *http.Request) {

}

//Main
func main() {
	//API
	router := mux.NewRouter()
	//Web front-end CORS
	headers := handlers.AllowedHeaders([]string{"X-REQUESTED-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Test API status
	router.HandleFunc("/api/v1/tutor", test)

	//3.6.1 View particular.
	//3.6.2 Update particular.
	router.HandleFunc("/api/v1/tutor/profile", profile).Methods("GET", "PUT")

	//3.6.3 View mod assigned.
	//3.6.4 View class assigned.
	//3.6.5 view timetable.
	//3.6.6 view enrolled students.
	router.HandleFunc("/api/v1/tutor/mod/{info}/{TutorID}", mod).Methods("GET")

	//3.6.7 List all tutors with ratings.
	//3.6.8 Search for other tutors.
	//3.6.9 View other tutor's profile, modules, class, timetable, ratings and comments.
	router.HandleFunc("/api/v1/tutor/details/{method}/{tutorEmail}", details).Methods("GET")

	//Establish port
	checkMicroservices()
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))

}