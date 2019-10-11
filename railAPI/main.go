package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rudglazier/dbutils"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource comment
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource comment
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

//ScheduleResource comment
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

//Register comment
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)

}

// Function Handlers from here

// GET http://localhost:8000/v1/trains/1

func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?, id").Scan(&t.ID, &t.DriverName, &t.OperatingStatus)

	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains

func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)

	// Omitting Error handling here

	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values(?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)

	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

}

func main() {
	// Connect to DB

	db, err := sql.Open("sqlite3", "railapi.db")

	if err != nil {
		log.Println("Driver Creation failed!")
	}

	// Create Tables

	dbutils.Initialize(db)

}
