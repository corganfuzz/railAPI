package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rudglazier/dbutils"
)

//DB Driver visible to whole program

var DB *sql.DB

type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

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

func main() {
	// Connect to DB

	db, err := sql.Open("sqlite3", "railapi.db")

	if err != nil {
		log.Println("Driver Creation failed!")
	}

	// Create Tables

	dbutils.Initialize(db)

}
