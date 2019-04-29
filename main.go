package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ActionModel struct {
	DB *mongo.Client
}

type Action struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID int                `json:"userID,omitempty" bson:"userID,omitempty`
	GameID int                `json:"gameID,omitempty" bson:"gameID,omitempty`
}

type Application struct {
	actionsql *ActionModel
}

func (m *ActionModel) Insert(userID, gameID int) (interface{}, error) {
	var action Action
	action.UserID = userID
	action.GameID = gameID
	collection := m.DB.Database("game").Collection("action")
	insertResult, err := collection.InsertOne(context.TODO(), action)
	if err != nil {
		return 0, err
	}

	return insertResult.InsertedID, nil
}

func (app *Application) CreateActionEndpoint(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	UserID, err := strconv.Atoi(r.Form.Get("UserID"))
	if err != nil {
		fmt.Println(err)
	}
	GameID, err := strconv.Atoi(r.Form.Get("GameID"))
	if err != nil {
		fmt.Println(err)
	}

	id, err := app.actionsql.Insert(UserID, GameID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}

func main() {

	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Disconnect(context.TODO())

	app := &Application{
		actionsql: &ActionModel{DB: db},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/action", app.CreateActionEndpoint).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func openDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
