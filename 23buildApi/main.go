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

// model for file

type Course struct {
   CourseId  string `json:"course_id"`
   Name      string `json:"name"`
   Price     int `json:"price"`
   Author    *Author `json:"author"`
}


type Author struct {
    Fullname string `json:"fullname"`
	Website  string `json:"website"` 
}

// fake db 

var courses = []Course{}

// middleware or helper function 

func (c *Course) IsEmpty() bool  {
	// return c.CourseId == "" && c.Name == "" 
	return  c.Name == "" 
}


// controllers 

func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>Welcome to Api </h1>"))
}


func getAllCourses(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json") 
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json") 

	params := mux.Vars(r) 
	fmt.Println(params) 

	// loop through all courses and find the matching course 
	for _,course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id") 
	return 
}

func createCourse(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json") 
	
	// dealing if body is empty 
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data 1")
		return 
	}

	// dealing empty object 
	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course) 

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return 
	}

	// generate unique course id and append to course
	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100)) 
	courses = append(courses, course) 
	 json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json") 
    
	params := mux.Vars(r) 

	for i,course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i],courses[i+1:]... )
			var _course Course 
			_ = json.NewDecoder(r.Body).Decode(&_course) 
			course.CourseId = params["id"] 
			courses = append(courses, _course) 
			json.NewEncoder(w).Encode(_course)
			return 
		}
	}
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json") 
	
    params := mux.Vars(r) 

	for i,c:=range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:i],courses[i+1:]... ) 
			break 
		}
	}
	 json.NewEncoder(w).Encode("Deleted successfully")
}

func main() {
	fmt.Println("------- Build API -------")

	r := mux.NewRouter() 

	// seeding data 
	courses = append(courses, Course{
		CourseId: "1",
		Name: "JS Bootcamp",
		Price: 250,
		Author: &Author{
			Fullname: "Mohammed Shahul",
			Website: "shahul.dev",
		},
	})
	courses = append(courses, Course{
		CourseId: "2",
		Name: "React Bootcamp",
		Price: 300,
		Author: &Author{
			Fullname: "Mohammed Shahul",
			Website: "shahul.dev",
		},
	})

	// routing 
	r.HandleFunc("/",serveHome).Methods("GET") 
	r.HandleFunc("/courses",getAllCourses).Methods("GET") 
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET") 
	r.HandleFunc("/course",createCourse).Methods("POST") 
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT") 
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE") 

	// listening on a port 
	log.Fatal(http.ListenAndServe(":4000",r)) 
}