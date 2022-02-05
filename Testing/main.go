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
	StudentID   string `json: "StudentID"`
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

type Class2 struct {
	ClassID    int    `json: "ClassID"`
	ModuleID   string `json: "ModuleID"`
	ClassDate  string `json: "ClassDate"`
	ClassStart string `json: "ClassStart"`
	ClassEnd   string `json: "ClassEnd"`
	Capacity   int    `json: "Capacity"`
	TutorfName string `json: "tutorName"`
	TutorID    int    `json: "TutorID"`
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
	var modlist []Module
	var tutor AssignedTutor
	tutor.TutorId = "1"
	tutor.ModuleCode = 1

	var student EnrolledStudent
	student.ClassId = 1
	student.Semester = "Sem 4"
	student.StudentID = "1"

	var mods Module
	mods.AssignedTutors = []AssignedTutor{AssignedTutor{"1", 1}, AssignedTutor{"2", 1}}
	mods.ModuleCode = "PRG1"
	mods.Synopsis = "Program python"
	mods.ModuleName = "Programming 1"
	mods.EnrolledStudents = []EnrolledStudent{student, EnrolledStudent{"2", 2, "Sem 4"}}
	mods.Classes = append(mods.Classes, 1)
	modlist = append(modlist, mods)

	var mods2 Module
	mods2.AssignedTutors = []AssignedTutor{AssignedTutor{"4", 1}, AssignedTutor{"3", 1}}
	mods2.ModuleCode = "PRG2"
	mods2.Synopsis = "Program C#"
	mods2.ModuleName = "Programming 2"
	mods2.EnrolledStudents = []EnrolledStudent{EnrolledStudent{"4", 2, "Sem 4"}, EnrolledStudent{"3", 2, "Sem 4"}}
	mods2.Classes = append(mods2.Classes, 2)
	modlist = append(modlist, mods2)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(modlist)
	return
}

func getRatingData(w http.ResponseWriter, r *http.Request) {
	ratingData := []RatingAndComments{RatingAndComments{1, 1, "good"}, RatingAndComments{2, 2, "Very good"}}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(ratingData)
}

func getClasses(w http.ResponseWriter, r *http.Request) {
	var classes []Class2
	classes = append(classes, Class2{1, "PRG1", "12 Feb 2021", "8pm", "9pm", 50, "John C maxwell", 1})
	classes = append(classes, Class2{2, "PRG2", "12 Feb 2021", "10pm", "11pm", 50, "John C maxwell", 1})
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(classes)
	return
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	var studentList []Student
	studentList = append(studentList, Student{"1", "John", "10 March 2000", "24 west c", "91234567"})
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
	router.HandleFunc("/api/v1/getClasses", getClasses).Methods("GET")
	router.HandleFunc("/api/v1/getStudent", getStudent).Methods("GET")

	fmt.Println("Listening at port 9032")
	log.Fatal(http.ListenAndServe(":9032", handlers.CORS(headers, methods, origins)(router)))
}
