package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/labstack/echo"
	"database/sql"
)


type EventResponse struct {
	Count  int `json:"count" `
	Type string `json:"type" `
  }

func yallo(c echo.Context) error {
	return c.String(http.StatusOK,"Try to use as /events?from=1234&to=1334&type=my_event ")
}

func getEvents(c echo.Context) error {
	v_from := c.QueryParam("from")
	v_to := c.QueryParam("to")
	v_type := c.QueryParam("type")

	connStr := "user=postgres dbname=postgres password=123456go host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	// select sum(count) from events where type='my_event' and timestamp between 1234 and 3000;
	statement :="select sum(count) from events where type=$1 and timestamp between $2 and $3"
	//prepare statement for sql
	stmt , err := db.Prepare(statement)
	rows, err := stmt.Query(v_type,v_from,v_to)
	
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"error":"There is something wrong",
		})
	}
	defer rows.Close()

	for rows.Next(){
		//assign values to variables
		var Count int
		err := rows.Scan(&Count)
		if err != nil {
			fmt.Print(err)
		}

		u := &EventResponse{
			Count:  Count,
			Type: v_type,
		  }
		return c.JSON(http.StatusOK, u)
	}//end of for loop

	return c.JSON(http.StatusOK, map[string]string{
		"error":"There is something wrong",
	})
}

func addEvent(c echo.Context) error{
	event := &Events{} // event proto is used here!

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&event)
	if err != nil {
		log.Printf("Failed processing event request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
		
	}
	connStr := "user=postgres dbname=postgres password=123456go host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	//sql to insert event information
	statement :="INSERT INTO events(count, type, timestamp ) VALUES($1, $2, $3)"
	//prepare statement for sql
	stmt , err := db.Prepare(statement)
	if err != nil {
		fmt.Print(err)
	}
	defer stmt.Close()
	//call a instant of event

	stmt.QueryRow(event.Count,event.Type,event.Timestamp)

	//select event first and last name
	rows, err := db.Query("Select count, type, timestamp from events")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Count Type Timestamp")
	//loop through all event results
	for rows.Next(){
		//assign values to variables
		var Count int
		var Type string
		var Timestamp int
		err := rows.Scan(&Count, &Type, &Timestamp)
		if err != nil {
			fmt.Print(err)
		}
		//print results to console
		fmt.Printf("%d %s %d\n",Count,Type,Timestamp)
	}//end of for loop


	log.Printf("this is your event: %#v", event)
	log.Printf("this is your event.Type: %s", event.Type)

	return c.String(http.StatusOK, "we got your event! ")
}
