package main

import (
	"github.com/ONSdigital/dp-preview-spike/observation"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"net/http"
	"os"
)

var connection1 bolt.Conn
var connection2 bolt.Conn
var connection3 bolt.Conn
var connection4 bolt.Conn
var connection5 bolt.Conn

func main() {
	log.Namespace = "preview"
	log.Info("starting preview service", nil)
	var err error
	connection1, err = bolt.NewDriver().OpenNeo("bolt://localhost:7687")
	connection2, err = bolt.NewDriver().OpenNeo("bolt://localhost:7687")
	connection3, err = bolt.NewDriver().OpenNeo("bolt://localhost:7687")
	connection4, err = bolt.NewDriver().OpenNeo("bolt://localhost:7687")
	connection5, err = bolt.NewDriver().OpenNeo("bolt://localhost:7687")
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	defer connection1.Close()
	defer connection2.Close()
	defer connection3.Close()
	defer connection4.Close()
	defer connection5.Close()

	router := mux.NewRouter()
	router.Path("/preview").HandlerFunc(previewHandler)
	router.Path("/preview1").HandlerFunc(previewHandler1)
	router.Path("/preview2").HandlerFunc(previewHandler2)
	router.Path("/preview3").HandlerFunc(previewHandler3)
	router.Path("/preview4").HandlerFunc(previewHandler4)

	log.Debug("starting http server", log.Data{"bind_addr": ":9091"})
	srv := server.New(":9091", router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func previewHandler(w http.ResponseWriter, req *http.Request) {
	geo := []*observation.DimensionOption{{Option: "K02000001"}}
	//cpi := []*observation.DimensionOption{{Option: "cpi1dim1G10100"}}
	dims := []*observation.DimensionFilter{{Name: "geography", Options: geo}}
	filter := observation.Filter{InstanceID: "edaf09a3-b76f-4600-9ccb-e4693c831c91", DimensionFilters: dims}
	s := observation.NewStore(connection1)
	rows, err := s.GetCSVRows(&filter, "20")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var csv string
	for {
		data, err := rows.Read()
		if err != nil {
			break
		}
		csv += data
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(csv))
}

func previewHandler1(w http.ResponseWriter, req *http.Request) {
	geo := []*observation.DimensionOption{{Option: "K02000001"}}
	cpi := []*observation.DimensionOption{{Option: "cpi1dim1G10100"}}
	dims := []*observation.DimensionFilter{{Name: "geography", Options: geo}, {Name: "aggregate", Options: cpi}}
	filter := observation.Filter{InstanceID: "edaf09a3-b76f-4600-9ccb-e4693c831c91", DimensionFilters: dims}
	s := observation.NewStore(connection2)
	rows, err := s.GetCSVRows(&filter, "15")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var csv string
	for {
		data, err := rows.Read()
		if err != nil {
			break
		}
		csv += data
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(csv))
}

func previewHandler2(w http.ResponseWriter, req *http.Request) {
	geo := []*observation.DimensionOption{{Option: "K02000001"}}
	cpi := []*observation.DimensionOption{{Option: "cpi1dim1G20200"}}
	time := []*observation.DimensionOption{{Option: "Aug-16"}}
	dims := []*observation.DimensionFilter{{Name: "geography", Options: geo}, {Name: "aggregate", Options: cpi}, {Name: "time", Options: time}}
	filter := observation.Filter{InstanceID: "edaf09a3-b76f-4600-9ccb-e4693c831c91", DimensionFilters: dims}

	s := observation.NewStore(connection3)
	rows, err := s.GetCSVRows(&filter, "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var csv string
	for {
		data, err := rows.Read()
		if err != nil {
			break
		}
		csv += data
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(csv))
}

func previewHandler3(w http.ResponseWriter, req *http.Request) {
	geo := []*observation.DimensionOption{{Option: "K02000001"}}
	cpi := []*observation.DimensionOption{{Option: "cpi1dim1G60100"}}
	dims := []*observation.DimensionFilter{{Name: "geography", Options: geo}, {Name: "aggregate", Options: cpi}}
	filter := observation.Filter{InstanceID: "edaf09a3-b76f-4600-9ccb-e4693c831c91", DimensionFilters: dims}

	s := observation.NewStore(connection4)
	rows, err := s.GetCSVRows(&filter, "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var csv string
	for {
		data, err := rows.Read()
		if err != nil {
			break
		}
		csv += data
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(csv))
}

func previewHandler4(w http.ResponseWriter, req *http.Request) {
	geo := []*observation.DimensionOption{{Option: "K02000001"}}
	cpi := []*observation.DimensionOption{{Option: "cpi1dim1G40300"}}
	dims := []*observation.DimensionFilter{{Name: "geography", Options: geo}, {Name: "aggregate", Options: cpi}}
	filter := observation.Filter{InstanceID: "edaf09a3-b76f-4600-9ccb-e4693c831c91", DimensionFilters: dims}

	s := observation.NewStore(connection5)
	rows, err := s.GetCSVRows(&filter, "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var csv string
	for {
		data, err := rows.Read()
		if err != nil {
			break
		}
		csv += data
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(csv))
}
