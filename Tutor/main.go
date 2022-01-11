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

//Database Functions
func getTutor(db *sql.DB, email string) Tutor {
	query := fmt.Sprintf("Select * FROM Customer WHERE EmailAddress= '%s'", email)
	results, err := db.Query(query)
	if err != nil {
		return err.Error()
	}
	var tutor Tutor
	for results.Next() {
		err = results.Scan(&tutor.ID, &tutor.FirstName,
			&tutor.LastName, &tutor.EmailAddress, &tutor.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	return tutor
}

func checkTutorExsist(db *sql.DB, email string) {
	//To check if tutor exsists and information is accurate
	query := fmt.Sprintf("Select EmailAddress FROM Tutor WHERE EmailAddress= '%s'", email)
	results, err := db.Query(query)
	if err != nil {
		return false
	}
	var tutor Tutor
	for results.Next() {
		// map this type to the record in the table
		err = results.Scan(&tutor.Email)
		if err != nil {
			panic(err.Error())
		} else if tutor.Emailtutor.Email == email {
			return true
		}
	}
	return false

}

func putUser(db *sql.DB, tutor Tutor) bool {
	query := fmt.Sprintf("UPDATE Tutor SET FirstName = '%s', LastName = '%s', MobileNumber= '%s' WHERE EmailAddress = '%s';", tutor.FirstName, tutor.LastName, tutor.MobileNumber, tutor.EmailAddress)
	_, err := db.Query(query)
	if err != nil {
		return false
	}
	return true
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
		db = database()

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
			if !checkTutorExsist(db, tutor.Email) { //To check if tutor exsists in the DB
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("There is no exsiting account for " + tutor.Email))
				return
			}
		}
		//Check method
		if r.Method == "GET" {
			//To get tutor's profile
			json.NewEncoder(w).Encode(getTutor(db, tutor.Email))
			w.WriteHeader(http.StatusAccepted)
			return
		} else if r.Method == "PUT" { //To update tutor's profile
			if tutor.password == "" {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please enter password"))
				return
			} else {
				putUser(db, tutor) //Update tutor's profile
				w.WriteHeader(http.StatusAccepted)
				return
			}
		}
	} else {
		w.WriteHeader(
			http.StatusUnprocessableEntity)
		w.Write([]byte(
			"422 - Please supply tutor's information"))
		return
	}
}
func mod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Email := params["Email"]
	Password := params["Password"]
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
	router.HandleFunc("/api/v1/tutor/details/{method}", details).Methods("GET")

	//Establish port
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))

}