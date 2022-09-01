package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "mongodbapi/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUrl = "mongodb://localhost:27017/"
const dbName = "netflix"
const collName = "watchList"

var collection *mongo.Collection

func init(){
	// client options 
	clientOption := options.Client().ApplyURI(dbUrl) 

	// connect to mongodb 
	client , err := mongo.Connect(context.TODO(),clientOption) 

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection Success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(collName)) 

	
}



func insertOneMovie(movie model.Netflix){
  inserted,err :=   collection.InsertOne(context.Background(),movie)

  if err != nil {
	log.Fatal(err)
  }

  fmt.Println("Inserted movie with id :",inserted.InsertedID) 
}


func updateMovie(movieId string ){
    id,_:= primitive.ObjectIDFromHex(movieId) 
	filter := bson.M{"_id":id} 
	update := bson.M{"$set":bson.M{"watched":true}}

	result,err :=  collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count : ",result.ModifiedCount)
}



func deleteOneMovie(movie_id string){
	id,_:= primitive.ObjectIDFromHex(movie_id) 

	filter := bson.M{"_id":id} 

    deleted, err := collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie got deleted with count :",deleted)
}


func deleteAllMovie(){
	
    res,err := collection.DeleteMany(context.Background(), bson.D{{}} )

	if err != nil{
     log.Fatal(err)
	}

	fmt.Println("Number of movies delete : ",res) 

}

func getAllMovie() [] primitive.A {

	cur,err := collection.Find(context.Background(),bson.D{{}}) 

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.A 

	for cur.Next(context.Background()) {
		var movie bson.A
		err := cur.Decode(&movie) 
		if err != nil {
			log.Fatal(err)
		} 
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background()) 
	return movies 
}



// actual controllers 

func GetAllMyMovies(w http.ResponseWriter,r *http.Request){
   w.Header().Set("Content-Type","application/x-www-form-urlencoded") 
   allMovies := getAllMovie() 
   json.NewEncoder(w).Encode(allMovies) 
}



func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded") 
    w.Header().Set("Allow-Control-Allow-Methods","POST") 

	var movie model.Netflix 

 	_ = json.NewDecoder(r.Body).Decode(&movie) 
	insertOneMovie(movie) 

	json.NewEncoder(w).Encode("New Movie created successfully") 

}


func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded") 
    w.Header().Set("Allow-Control-Allow-Methods","PUT") 

	params := mux.Vars(r) 

	updateMovie(params["id"]) 

	json.NewEncoder(w).Encode("Movie marked as watched") 
	
}

func DeleteOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded") 
    w.Header().Set("Allow-Control-Allow-Methods","DELETE") 

	params := mux.Vars(r) 

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode("Movie is deleted") 
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded") 
    w.Header().Set("Allow-Control-Allow-Methods","DELETE") 

	deleteAllMovie()

	json.NewEncoder(w).Encode("All Movies are deleted") 
}

