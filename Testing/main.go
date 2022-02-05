package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Classes
type Tutor struct {
	TutorID     int    `json: "TutorID"`
	FirstName   string `json: "firstname"`
	LastName    string `json: "lastname"`
	Email       string `json: "email"`
	Description string `json: "descriptions"`
}

type Student struct {
	StudentID   int    `json: "StudentID"`
	Name        string `json: "Name"`
	DateOfBirth string `json: "DateOfBirth"`
	Address     string `json: "Address"`
	Number      string `json: "Number"`
}

type Class struct {
	Code     int    `json: "Code"`
	Schedule string `json: "Schedule"`
	Capacity int    `json: "Capacity"`
}

type RatingAndComments struct {
	TutorID  int    `json: "TutorID"`
	Rating   int    `json: "Rating"`
	Comments string `json: "Comments"`
}

type Module struct {
	Code              int       `json: "Code"`
	Name              string    `json: "Name"`
	LearningObjective string    `json: "LearningObjective"`
	Classes           []Class   `json: "Classes"`
	AssignedTutor     int       `json: "AssignedTutor"`
	EnrolledStudent   []Student `json: "EnrolledStudent"`
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: Tutor API Is working")
}

func getTutor(w http.ResponseWriter, r *http.Request) {
	var tutor Tutor
	tutor.TutorID = 1
	tutor.FirstName = "John"
	tutor.LastName = "Lee"
	tutor.Email = "JohnLee@np.com"
	tutor.Description = "Hello world"
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(tutor)
	return
}

func getTutorList(w http.ResponseWriter, r *http.Request) {
	var tutorList []Tutor
	var tutor Tutor
	tutor.TutorID = 1
	tutor.FirstName = "John"
	tutor.LastName = "Lee"
	tutor.Email = "JohnLee@np.com"
	tutor.Description = "Hello world"

	var tutor1 Tutor
	tutor1.TutorID = 1
	tutor1.FirstName = "John"
	tutor1.LastName = "Lee"
	tutor1.Email = "JohnLee@np.com"
	tutor1.Description = "Hello world"

	tutorList = append(tutorList, tutor)
	tutorList = append(tutorList, tutor1)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(tutorList)
	return
}
func putTutor(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") == "application/json" {
		reqBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var tutor Tutor
		if err == nil {
			err := json.Unmarshal(reqBody, &tutor)
			if err != nil {
				fmt.Printf("There was an error encoding the json. err = %s", err)
			} else if tutor.Email != "" || tutor.Description != "" {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("Account has bee created successfully"))
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Please enter Email"))
			}
		}
	}

	return
}

func getMod(w http.ResponseWriter, r *http.Request) {
	var mods Module
	mods.AssignedTutor = 1
	mods.Code = 1
	mods.LearningObjective = "Math"
	mods.Name = "Math"
	mods.EnrolledStudent = []Student{Student{1, "john", "28 July", "25 west coast", "1234567"}, Student{2, "Susan", "28 July", "25 west coast", "1234567"}}
	mods.Classes = []Class{Class{1, "8.00Am - 9.00Am", 1}, Class{2, "8.30Am - 9.30Am", 1}}

	var mods2 Module
	mods2.AssignedTutor = 2
	mods2.Code = 2
	mods2.LearningObjective = "English"
	mods2.Name = "English"
	mods2.EnrolledStudent = []Student{Student{1, "john", "28 July", "25 west coast", "1234567"}, Student{2, "Susan", "28 July", "25 west coast", "1234567"}}
	mods2.Classes = []Class{Class{1, "8.15Am - 9.15Am", 1}, Class{2, "8.45Am - 9.45Am", 1}}

	modList := [2]Module{mods2, mods}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(modList)
	return
}

func getRatingData(w http.ResponseWriter, r *http.Request) {
	ratingData := []RatingAndComments{RatingAndComments{1, 1, "good"}, RatingAndComments{2, 2, "Very good"}}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(ratingData)
}

func main() {
	//API
	router := mux.NewRouter()
	//Web front-end CORS
	headers := handlers.AllowedHeaders([]string{"X-REQUESTED-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Test API status
	router.HandleFunc("/api/v1/tutor", test)

	router.HandleFunc("/api/v1/getTutor/{tutorID}", getTutor).Methods("GET")
	router.HandleFunc("/api/v1/putTutor", putTutor).Methods("PUT")
	router.HandleFunc("/api/v1/getMod", getMod).Methods("GET")
	router.HandleFunc("/api/v1/getRatingData", getRatingData).Methods("GET")
	router.HandleFunc("/api/v1/getTutorList", getTutorList).Methods("GET")

	fmt.Println("Listening at port 9032")
	log.Fatal(http.ListenAndServe(":9032", handlers.CORS(headers, methods, origins)(router)))
}
