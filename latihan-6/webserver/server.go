package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Id       string
	Name     string
	Age      int
	Division string
}

var users = []User{
	{
		Id:       "user-1",
		Name:     "User 1",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-2",
		Name:     "User 2",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-3",
		Name:     "User 3",
		Age:      20,
		Division: "IT",
	},
	{
		Id:       "user-4",
		Name:     "User 4",
		Age:      20,
		Division: "IT",
	},
}

type HTML struct {
	Title string
	Data  []User
}

func main() {

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/add", createUser)

	port := ":8001"

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func getUsers(res http.ResponseWriter, req *http.Request) {
	traceId := "91aaf539-3b07-42fd-bd05-713e7f5f7c06"
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("x-trace-id", traceId)
	header := req.Header.Get("client-id")

	if req.Method == http.MethodGet {
		log.Println("[INFO]", req.Method, req.URL.Path, "trace id :", traceId)
		if header == "API" {

			json.NewEncoder(res).Encode(users)
			return
		} else {
			res.Header().Set("Content-Type", "text/html")
			tpl, err := template.ParseFiles("./index.html")
			if err != nil {
				log.Println("[ERROR]", req.Method, req.URL.Path, "trace id :", traceId, "error :", err.Error())
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			d := HTML{
				Title: "halaman users",
				Data:  users,
			}

			tpl.Execute(res, d)
			return
		}
	}

	errMsg := fmt.Sprintf("method %s not allowed !", req.Method)
	log.Println("[ERROR]", req.Method, req.URL.Path, "trace id :", traceId, "error :", errMsg)
	res.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(res).Encode(map[string]string{
		"error": "Ga boleh oi",
	})
}

func createUser(res http.ResponseWriter, req *http.Request) {
	traceId := "91aaf539-3b07-42fd-bd05-713e7f5f7c06"
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("x-trace-id", traceId)

	if req.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(map[string]string{
				"error": "bad request",
			})
			log.Println("[ERROR]", req.Method, req.URL.Path, "trace id :", traceId, "error :", err.Error())
			return
		}

		user.Id = fmt.Sprintf("user-%d", len(users)+1)

		users = append(users, user)
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "created success",
		})
		return
	}

	errMsg := fmt.Sprintf("method %s not allowed !", req.Method)
	log.Println("[ERROR]", req.Method, req.URL.Path, "trace id :", traceId, "error :", errMsg)
	res.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(res).Encode(map[string]string{
		"error": "Ga boleh oi",
	})
}
