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
type EnrolledStudent struct {
	StudentID string `json: "student_id"`
	ClassId   int    `json: "class_id"`
	Semester  string `json: "semester"`
}
type AssignedTutor struct {
	TutorId    string `json: "tutorid"`
	ModuleCode int    `json: "modulecode"`
}
type Module struct {
	ModuleCode         string            `json:"modulecode"`
	ModuleName         string            `json:"modulename"`
	Synopsis           string            `json:"synopsis"`
	LearningObjectives string            `json:"learningobjective"`
	Classes            []int             `json:"classes"`
	AssignedTutors     []AssignedTutor   `json:"assigned_tutors"`
	EnrolledStudents   []EnrolledStudent `json:"enrolled_students"`
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
	var tutor AssignedTutor
	tutor.TutorId = "S1234567G"
	tutor.ModuleCode = 1

	var student EnrolledStudent
	student.ClassId = 1
	student.Semester = "Sem 4"
	student.StudentID = "S2345678A"

	var mods Module
	mods.AssignedTutors = append(mods.AssignedTutors, tutor)
	mods.ModuleCode = "PRG1"
	mods.Synopsis = "Program python"
	mods.ModuleName = "Programming 1"
	mods.EnrolledStudents = []EnrolledStudent{student, EnrolledStudent{"S1234567C", 2, "Sem 4"}}
	mods.Classes = append(mods.Classes, 1)

	var mods2 Module
	mods.AssignedTutors = append(mods.AssignedTutors, tutor)
	mods.ModuleCode = "PRG2"
	mods.Synopsis = "Program C#"
	mods.ModuleName = "Programming 2"
	mods.EnrolledStudents = []EnrolledStudent{student, EnrolledStudent{"S1234567C", 2, "Sem 4"}}
	mods.Classes = append(mods.Classes, 1)

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
