package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pixelbender/go-traceroute/traceroute"
)

type TraceBind struct {
	Domain string `json:"domain"`
}

type TracePayload struct {
	Status  string  `json:"status"`
	IP      string  `json:"query"`
	Country string  `json:"country_name"`
	City    string  `json:"city"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

func (d *TraceBind) Bind(r *http.Request) error {
	if len(d.Domain) == 0 {
		return errors.New("you should provide a valid domain")
	}

	return nil
}

func main() {
	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", GetFile)
	r.With(DestAddressMiddleware).Get("/trace", GetTraces)

	log.Println("Server Started!")

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	fileName := "index.html"

	t, err := template.ParseFiles(fileName)
	if err != nil {
		render.Status(r, 500)
	}

	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		render.Status(r, 500)
	}
}

func GetTraces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("get trace")

	ctx := r.Context()

	destAddress := ctx.Value(DestKey).(string)
	ip_list, err := net.LookupIP(destAddress)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("\"error\": \"domain address could not be found\""))
		return
	}

	ips := []net.IP{}

	for _, ip := range ip_list {

		if ipv4 := ip.To4(); ipv4 != nil {

			fmt.Println("IPv4: ", ip)
			ips = append(ips, ip)
		}
	}

	destIP := ips[0].String()
	locations := []TracePayload{}

	hops, err := traceroute.Trace(net.ParseIP(destIP))
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range hops {
		for _, n := range h.Nodes {
			log.Printf("%d. %v %v", h.Distance, n.IP, n.RTT)
			result, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%v", n.IP.String()))

			if err != nil {
				log.Println(err)
				return
			}

			body, err := io.ReadAll(result.Body) // response body is []byte
			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println(string(body))

			//log.Println(results)

			location := &TracePayload{}
			json.Unmarshal(body, location)

			locations = append(locations, *location)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
	//log.Println(locations)
}
