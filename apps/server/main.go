package main

import (
	"fmt"
	"log"
	"net/http"

	arri "github.com/modiimedia/arri/languages/go/go-server"
)

func main() {
	app := arri.NewApp(
		http.DefaultServeMux,
		arri.AppOptions[RpcEvent]{
			OnRequest: func(c *RpcEvent) arri.RpcError {
				c.Writer().Header().Set("Access-Control-Allow-Origin", "*")
				return nil
			},
		},
		// initialize an RpcEvent using the incoming response writer and http request
		// you can extend RpcEvent in ./rpc_event.go
		func(w http.ResponseWriter, r *http.Request) (*RpcEvent, arri.RpcError) {
			return &RpcEvent{W: w, R: r}, nil
		},
	)

	// register procedures on the app
	arri.Rpc(&app, SayHello, arri.RpcOptions{
		Method:      http.MethodGet,
		Description: "Katakan hello dari namanu",
	})
	arri.Rpc(&app, SayGoodbye, arri.RpcOptions{
		Method:      http.MethodGet,
		Description: "Katakan Goodbye",
	})
	arri.Rpc(&app, GetAddress, arri.RpcOptions{
		Method:      http.MethodGet,
		Description: "Ambil Alamat Mu",
	})
	arri.Rpc(&app, GetUsers, arri.RpcOptions{
		Method:      http.MethodGet,
		Description: "Paginate Users Data",
	})

	err := app.Run(arri.RunOptions{Port: 4000})
	if err != nil {
		log.Fatal(err)
	}
}

type GreetingParams struct {
	Name string
	Age  uint16
}

type GreetingResponse struct {
	Message string
}

type OptionalParams struct {
	Address *string
}
type PaginationResponse struct {
	Data *[]User
}

type PaginationParams struct {
	Page   uint16
	Limit  uint16
	Offset uint16
}

type User struct {
	Name string
	Age  uint16
}

func SayHello(params GreetingParams, event RpcEvent) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Hello %s, Your age is %d", params.Name, params.Age)}, nil
}

func SayGoodbye(params GreetingParams, event RpcEvent) (GreetingResponse, arri.RpcError) {
	return GreetingResponse{Message: fmt.Sprintf("Goodbye %s", params.Name)}, nil
}

func GetAddress(params OptionalParams, event RpcEvent) (GreetingResponse, arri.RpcError) {
	if params.Address == nil {
		return GreetingResponse{Message: "No address provided"}, nil
	}
	return GreetingResponse{Message: fmt.Sprintf("Your address is %s", *params.Address)}, nil
}

var dummyUsers = []User{
	{
		Name: "John",
		Age:  20,
	},
	{
		Name: "Jane",
		Age:  21,
	},
	{
		Name: "Jack",
		Age:  22,
	},
}

func GetUsers(params PaginationParams, event RpcEvent) (PaginationResponse, arri.RpcError) {
	return PaginationResponse{Data: &dummyUsers}, nil
}
