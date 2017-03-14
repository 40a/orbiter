package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gianarb/orbiter/core"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetRouter(core core.Core, eventChannel chan *logrus.Entry) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/handle/{autoscaler_name}/{service_name}", Handle(core.Autoscalers)).Methods("POST")
	r.HandleFunc("/health", Health()).Methods("GET")
	r.HandleFunc("/events", Events(eventChannel)).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())
	return r
}
