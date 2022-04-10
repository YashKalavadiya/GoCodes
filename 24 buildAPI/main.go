package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course -file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("API is up and running")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId: "2", 
		CourseName: "React Js", 
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Hitesh Choudhary",
			Website: "lco.dev",
		},
	})
	courses = append(courses, Course{
		CourseId: "4", 
		CourseName: "MERN Stack", 
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Hitesh Choudhary",
			Website: "go.dev",
		},
	})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To Go API by yash</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from req
	params := mux.Vars(r)

	//loop through courses and find matching ID and return the response

	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found With given id: " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	//what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please Send Required Fields")
		return
	}

	for _, val := range courses {
		if val.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("No Duplicate course allowed")
			return
		}
	}

	//generate a UID, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab ID from REQ
	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("There is no course with this ID")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting Course....")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("There is no course with this ID")
}