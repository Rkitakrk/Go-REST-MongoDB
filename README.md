# Go-REST-MongoDB

## **_Go_**

You have to download my file "main.go"

We have to import all packages to `mux` and `mongo-driver`, so use this command belows:

`go get github.com/gorilla/mux`

`go get go.mongodb.org/mongo-driver/bson/primitive`

`go get go.mongodb.org/mongo-driver/mongo`

`go get go.mongodb.org/mongo-driver/mongo/options`

Then we have to run file "main.go" with command:

`go run main.go`

---

## \***\*Send POST Request\*\***

We can use `curl` or `Postman`

**CURL**

`curl -d '{"UserID": 1, "GameID": 2, "Action": {"X": 12, "Y": 12, "Z": 12}}' -H "Content-Type: application/json" -X POST http://localhost:8000/action`

**POSTMAN**

postman is a program which help send requests and is very helpful.

![postman](/images/postman.png)

---

## **_Router_**

/ - HomePage

/action - save data to our MongoDB

![sql](/images/sql.png)
