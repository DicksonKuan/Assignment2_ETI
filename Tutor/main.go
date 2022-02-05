package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Official Classes
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
	ClassID    int    `json: "ClassID"`
	ModuleID   string `json: "ModuleID"`
	ClassDate  string `json: "ClassDate"`
	ClassStart string `json: "ClassStart"`
	ClassEnd   string `json: "ClassEnd"`
	Capacity   int    `json: "Capacity"`
	TutorfName string `json: "tutorName"`
	TutorID    int    `json: "TutorID"`
}

type Module struct {
	Code              string    `json: "Code"`
	Name              string    `json: "Name"`
	LearningObjective string    `json: "LearningObjective"`
	Classes           []Class   `json: "Classes"`
	AssignedTutor     string    `json: "AssignedTutor"`
	EnrolledStudent   []Student `json: "EnrolledStudent"`
}

//API Classes
type EnrolledStudent struct {
	StudentID string `json: "student_id"`
	ClassId   int    `json: "class_id"`
	Semester  string `json: "semester"`
}
type AssignedTutor struct {
	TutorId    string `json: "tutorid"`
	ModuleCode int    `json: "modulecode"`
}
type getModule struct {
	ModuleCode         string            `json:"modulecode"`
	ModuleName         string            `json:"modulename"`
	Synopsis           string            `json:"synopsis"`
	LearningObjectives string            `json:"learningobjective"`
	Classes            []int             `json:"classes"`
	AssignedTutors     []AssignedTutor   `json:"assigned_tutors"`
	EnrolledStudents   []EnrolledStudent `json:"enrolled_students"`
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
func getTutor(tutorID int) Tutor {
	//url := fmt.Sprintf("http://localhost:9181/api/v1/tutor/GetaTutorByEmail/%d", tutorID)
	url := fmt.Sprintf("http://localhost:9032/api/v1/getTutor/%d", tutorID)
	response, err := http.Get(url)
	var tutor Tutor
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err == nil {
			err = json.Unmarshal(responseData, &tutor)
		}
		println(err)
	}
	return tutor
}

func checkTutorExsist(tutorID int) bool {
	//To check if tutor exsists and information is accurate
	url := fmt.Sprintf("http://localhost:9032/api/v1/getTutor/%d", tutorID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == 202 {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
		} else {
			var tutor Tutor
			err = json.Unmarshal([]byte(responseData), &tutor)
			if err == nil {
				return true
			}
		}
	}
	return false
}

func putUser(tutor Tutor) bool { //Update tutor's profile
	jsonValue, _ := json.Marshal(tutor)
	//URL := "http://localhost:9181/api/v1/tutor/UpdateTutorAccountByEmail/" + tutor.Email
	URL := "http://localhost:9032/api/v1/putTutor"

	request, err := http.NewRequest(http.MethodPut,
		URL,
		bytes.NewBuffer(jsonValue))

	if err != nil {
		println(err.Error())
	} else {
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode == http.StatusCreated {
				response.Body.Close()
				return true
			}
		}
		response.Body.Close()
	}
	return false
}

func getMod(tutorID string) []Module { //get mod from mod microservice
	//URL := http://10.31.11.12:9061/module/v1/modules/+tutorID
	URL := "http://localhost:9032/api/v1/getMod"
	response, err := http.Get(URL)
	if err != nil {
		fmt.Print(err.Error())
	} else if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err)
		} else {
			//Get module
			var newMods []getModule
			err := json.Unmarshal(responseData, &newMods)
			if err != nil {
				panic(err.Error())
			}

			var modList []Module
			for _, data := range newMods {
				var mods Module
				var allClasses []Class
				var allStudent []Student
				mods.Code = data.ModuleCode
				mods.Name = data.ModuleName
				mods.LearningObjective = data.Synopsis
				mods.AssignedTutor = tutorID
				//GET all classes
				response, err := http.Get("http://localhost:9101/api/v1/class?key=2c78afaf-97da-4816-bbee-9ad239abb296")
				if err != nil {
					fmt.Print(err.Error())
				} else if response.StatusCode == http.StatusAccepted {
					responseData, err := ioutil.ReadAll(response.Body)
					if err != nil {
						println(err)
					}
					err = json.Unmarshal(responseData, &allClasses)
					if err != nil {
						panic(err.Error())
					}
				}

				//GET class details
				for _, classID := range data.Classes {
					for _, Classes := range allClasses {
						if Classes.ClassID == classID {
							mods.Classes = append(mods.Classes, Classes)
							break
						}
					}
				}
				//GET all Student
				response, err = http.Get("http://10.31.11.12:9211/api/v1/students/")
				if err != nil {
					fmt.Print(err.Error())
				} else if response.StatusCode == http.StatusAccepted {
					responseData, err := ioutil.ReadAll(response.Body)
					if err != nil {
						println(err)
					}
					err = json.Unmarshal(responseData, &allStudent)
					if err != nil {
						panic(err.Error())
					}
				}

				//GET Student Details
				for _, Student := range data.EnrolledStudents {
					for _, StudentDetails := range allStudent {
						if Student.StudentID == StudentDetails.StudentID {
							mods.EnrolledStudent = append(mods.EnrolledStudent, StudentDetails)
						}
					}
				}
			}

			return modList
		}
	}
	return nil
}

func getClassAssigned(tutorID string) []Class {
	//Get Assigned mods
	mods := getMod(tutorID)

	//Get all classes
	var classesInfo []Class
	for _, modules := range mods {
		classesInfo = append(classesInfo, modules.Classes...)
	}
	return classesInfo
}

func getEnrolledStudent(tutorID string) []Student {
	mods := getMod(tutorID)
	var studentList []Student
	//Get modules from mods list
	for _, module := range mods {
		//Get students from the student list
		for _, student := range module.EnrolledStudent {
			//Check if student exsist in student list
			checkStudentExsist := true
			for _, stud := range studentList {
				if stud.StudentID == student.StudentID {
					checkStudentExsist = false
				}
			}
			if checkStudentExsist {
				studentList = append(studentList, student)
			}
		}
	}
	return studentList
}
func getAllTutor() []Tutor {
	var tutorList []Tutor
	url := "http://localhost:9032/api/v1/getTutorList"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil || json.Unmarshal([]byte(responseData), &tutorList) != nil {
			println(err)
		}
	}
	return tutorList
}
func getOtherTutor(tutorEmail string) Tutor {
	//url := "http://localhost:5000/api/v1/tutor/" + tutorEmail
	url := "http://localhost:9032/api/v1/getTutor/1"
	response, err := http.Get(url)
	var tutor Tutor

	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil || json.Unmarshal([]byte(responseData), &tutor) != nil {
			println(err)
		}
	}
	return tutor
}

func viewTutorProfile(tutorEmail string) Tutor {
	//url := "http://localhost:5000/api/v1/tutor/" + tutorEmail
	url := "http://localhost:9032/api/v1/getTutor/1"
	response, err := http.Get(url)
	var tutor Tutor

	if err != nil {
		fmt.Print(err.Error())
	}
	if response.StatusCode == http.StatusAccepted {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil || json.Unmarshal([]byte(responseData), &tutor) != nil {
			println(err)
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

	//Get information from JSON and validation
	var tutor Tutor
	params := mux.Vars(r)
	//password := params["Password"]
	tutorIDParam := params["TutorID"]
	tutorID, err := strconv.Atoi(tutorIDParam)
	if tutorID == 0 || !checkTutorExsist(tutorID) || err != nil { //To check for information not empty
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Please supply a valid tutor's tutorID"))
		return
	}
	//Check method
	if r.Method == "GET" {
		//To get tutor's profile
		tutor = getTutor(tutorID)
		if tutor == (Tutor{}) { //Check if tutor is empty
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Could not retrieve tutor"))
		} else {
			json.NewEncoder(w).Encode(tutor)
			w.WriteHeader(http.StatusAccepted)
		}
		return
	} else if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "PUT" { //To update tutor's profile
			//Unmarshal JSON
			reqBody, err1 := ioutil.ReadAll(r.Body)
			if err1 != nil {
				println(err1.Error())
			}
			defer r.Body.Close()
			var newTutorData Tutor
			err = json.Unmarshal(reqBody, &newTutorData)
			if !putUser(newTutorData) || err != nil { //Check if password is empty
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("User fail to update"))
				return
			} else {
				//Update tutor's profile
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Account updated"))
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
	//Get and convert parameter
	params := mux.Vars(r)
	method := params["method"]
	tutorID := params["TutorID"]

	//To check if param is empty
	if method == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Please supply tutor's information and valid method"))
		return
	} else {
		//To run function according to the method selected
		switch string(method) {
		case "getMod":
			mods := getMod(tutorID)
			if len(mods) == 0 {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Mod list Empty"))
			} else {
				json.NewEncoder(w).Encode(mods)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getClassAssigned":
			classes := getClassAssigned(tutorID)
			if len(classes) == 0 {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"class list Empty"))
			} else {
				json.NewEncoder(w).Encode(classes)
				w.WriteHeader(http.StatusAccepted)
			}
		case "enrolledStudent":
			students := getEnrolledStudent(tutorID)
			if len(students) == 0 {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Student list Empty"))
			} else {
				json.NewEncoder(w).Encode(students)
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
		case "getOtherTutor":
			tutor := getOtherTutor(email)
			if tutor == (Tutor{}) {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Cannot find tutor"))
			} else {
				json.NewEncoder(w).Encode(tutor)
				w.WriteHeader(http.StatusAccepted)
			}
		case "viewTutorProfile":
			tutor := viewTutorProfile(email)
			if tutor == (Tutor{}) {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Cannot find tutor"))
			} else {
				json.NewEncoder(w).Encode(tutor)
				w.WriteHeader(http.StatusAccepted)
			}
		case "getAllTutor":
			var tutorList []Tutor
			tutorList = getAllTutor()
			json.NewEncoder(w).Encode(tutorList)
		}
	}
}

//Main
func main() {
	//API
	router := mux.NewRouter()
	//Web front-end CORS
	headers := handlers.AllowedHeaders([]string{"X-REQUESTED-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Test API status
	router.HandleFunc("/api/v1/tutor", test)

	//3.6.1 View particular.
	//3.6.2 Update particular.
	//JSON, get data using tutorID and tutorPassword
	router.HandleFunc("/api/v1/tutor/profile/{TutorID}/{Password}", profile).Methods("GET", "PUT")

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
	//checkMicroservices()
	fmt.Println("Listening at port 9031")
	log.Fatal(http.ListenAndServe(":9031", handlers.CORS(headers, methods, origins)(router)))

}
