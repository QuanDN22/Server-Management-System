package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	// run log
	// logFile := logger.LogFile()
	// defer logFile.Close()

	// // Connection to database
	// db := models.Connection_DB()

	// // Initialize database
	// models.Init_DB(db)

	// // add_data into database
	// models.AddData_Init(db)

	// // export data from database into excel file
	// models.ExportData_Example(db)

	// c, err := config.NewConfig("./pkg/config", ".env.auth")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Println("config parsed...")
	// fmt.Println(c)
	// fmt.Println(c.ServiceName)
	// fmt.Println(c.GrpcAddr)
	// fmt.Println(c.GrpcPort)

	// port := flag.String("p", "3000", "port to serve on")
	// directory := flag.String("d", ".", "the directory of static file to host")
	// flag.Parse()

	// http.Handle("/", http.FileServer(http.Dir(*directory)))

	// log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	// log.Fatal(http.ListenAndServe(":"+*port, nil))

	// port := flag.String("p", "3000", "port to serve on")
	// directory := flag.String("d", ".", "./static/openapiv2/auth/auth.swagger.json")
	// flag.Parse()

	// http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir(*directory))))

	// log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	// log.Fatal(http.ListenAndServe(":3000", nil))

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
		fmt.Println("Error in creating scheduler")
	}

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.DurationJob(
			time.Second * 5,
		),
		gocron.NewTask(
			func() {
				fmt.Println("hello abc")
			},
		),
	)
	if err != nil {
		// handle error
		fmt.Println("Error in adding job to scheduler")
	}

	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	c := make(chan byte)
	<-c

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
		fmt.Println("Error in shutting down scheduler")
	}
}
