// package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Classes
type Tutor struct {
	TutorID   int    `json: "TutorID"`
	FirstName string `json: "FirstName"`
	LastName  string `json: "LastName"`
	Email     string `json: "Email"`
	password  string `json: "Password"`
}

type Student struct {
	StudentID   int    `json: "StudentID"`
	Name string `json: "Name"`
	DateOfBirth  string `json: "DateOfBirth"`
	Address     string `json: "Address"`
	Number  string `json: "Number"`
}

type Class struct {
	Code   		int    `json: "Code"`
	Schedule 	string `json: "Schedule"`
	Capacity  	int `json: "Capacity"`
}

type RatingAndComments struct {
	TutorID   	int    `json: "TutorID"`
	Rating 		int `json: "Rating"`
	Comments  	string `json: "Comments"`
}

type Module struct {
	Code   int    `json: "Code"`
	Name string `json: "Name"`
	LearningObjective  string `json: "LearningObjective"`
	Classes     []Class `json: "Classes"`
	AssignedTutor  int `json: "AssignedTutor"`
	EnrolledStudent  []Student `json: "EnrolledStudent"`
	RatingsAndComments []RatingAndComments `json: "RatingsAndComments"`
}

type Timetable struct {
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

func getTutor(tutorID int) Tutor {
	url := "http://localhost:5000/api/v1/tutor/" + tutorID
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
	return nil
}

func checkTutorExsist(tutorID int) bool {
	//To check if tutor exsists and information is accurate
	url := "http://localhost:5000/api/v1/tutor/checkTutor/" + tutorID
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
func getMod(tutorID int) []Module { //get mod from mod microservice
	URL := "http://localhost:5000/api/v1/CheckUser/" + tutorID

	response, err := http.Get(URL)
	if err != nil {
		fmt.Print(err.Error())
		return nil	
	} else if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
		} else {
			// mods := strings.Split(string(responseData), ",")
			// replacer := strings.NewReplacer(",", "")
			var newMods []Module
			err := json.Unmarshal(responseData, &newMods)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Error encoding the json."))
				return
			}
			// for i := range mods {
			// 	newMods = append(newMods, replacer.Replace(mods[i]))
			// }
			return newMods
		}
	}
	return nil
}

func getClassAssigned(tutorID int) []Class{
	//Get Assigned mods
	var mods = []Module
	mods := getMod(tutorID)

	//Get all classes 
	var classesInfo = []Class
	for i, classCode := range mods.classes{
		classURL := "http://localhost:5000/api/v1/CheckUser/" + classCode
		response, err := http.Get(classURL)
		if err != nil {
			fmt.Print(err.Error())
			return nil
		} else if response.StatusCode == http.StatusAccepted {
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				println(err)
				return nil
			} else {
				classesInfo = append(classesInfo, responseData)
			}
		}
	}
	return classesInfo
}

func getTimetable(){
	//Work in progress
}

func getEnrolledStudent(tutorID string) []Student{
	mods := getMod(tutorID)
	var studentList = []Student
	//Get modules from mods list
	for i, module := range mods{
		//Get students from the student list
		for x, student := range module.EnrolledStudent{
			//Check if student exsist in student list
			checkStudentExsist := true
			for y, stud := range studentList{
				if(stud.StudentID == student.StudentID){
					checkStudentExsist = false
				}
			}
			if(checkStudentExsist==false){
				studentList = append(studentList, student)
			}
		}
	}
	return studentList
}

func getListTutorAndRating() []Tutor{
	response, err := http.Get("http://localhost:4000/api/v1/GetAllDriver")
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
		} else {
			var tutors []Tutor
			err := json.Unmarshal(responseData, &tutors)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Error encoding the json."))
			}else{
				return tutors
			}
		}
	}
	return nil
})

func getOtherTutor(tutorEmail string){
	url := "http://localhost:5000/api/v1/tutor/" + tutorEmail	
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
			var tutors []Tutor
			err = json.Unmarshal([]byte(responseData), &tutors)
			return tutors
		}
	}
	return tutor
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
			w.Write([]byte("Please provide tutorID or password"))
			return
		} else {
			json.Unmarshal(reqBody, &tutor)
			if tutor.tutorID == 0 { //To check for information not empty
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Please supply tutor's tutorID"))
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
			tutor = getTutor(tutor.TutorID)
			if tutor == ""{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("Could not retrieve tutor"))
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
	//Get parameter
	params := mux.Vars(r)
	method := params["method"]
	tutorID := params["TutorID"]

	//To check if param is empty
	if method == "" || tutorID == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Please supply tutor's information and valid method"))
		return
	} else {
		//To run function according to the method selected
		switch method {
		case "getMod":
			mods := getMod(tutorID)
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
			classes := getClassAssigned(tutorID)
			if len(classes) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"class list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getTimetable":
			timetable := getTimetable(tutorID)
			if len(timetable) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"timetable list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		case "enrolledStudent":
			students := getEnrolledStudent(tutorID)
			if len(students) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Student list Empty"))
			}else{
				json.NewEncoder(w).Encode(JSONObject)
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
}
func details(w http.ResponseWriter, r *http.Request) {
	//Get params value
	params := mux.Vars(r)
	method := params["method"]
	email := params["email"]

	//Check param is empty
	if method == "" || email == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Please supply tutor's information and valid method"))
		return
	} else {
		switch method {
		case "getListTutorAndRating":
			mods := getListTutorAndRating()
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"tutor list Empty"))
			}else{
				json.NewEncoder(w).Encode(mods)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getOtherTutor":
			classes := getOtherTutor(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"class list Empty"))
			}else{
				json.NewEncoder(w).Encode(classes)
				w.WriteHeader(http.StatusAccepted)
			}
		case "viewTutorProfile":
			tutor := viewTutorProfile(email)
			if len(mods) == 0{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"timetable list Empty"))
			}else{
				json.NewEncoder(w).Encode(tutor)
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
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
	//JSON, get data using tutorID and tutorPassword
	router.HandleFunc("/api/v1/tutor/profile", profile).Methods("GET", "PUT")

	//3.6.3 View mod assigned.
	//3.6.4 View class assigned.
	//3.6.5 view timetable.
	//3.6.6 view enrolled students.
	router.HandleFunc("/api/v1/tutor/mod/{method}/{TutorID}", mod).Methods("GET")

	//3.6.7 List all tutors with ratings.
	//3.6.8 Search for other tutors.
	//3.6.9 View other tutor's profile, modules, class, timetable, ratings and comments.
	router.HandleFunc("/api/v1/tutor/details/{method}/{email}", details).Methods("GET")

	//Establish port
	checkMicroservices()
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))

}