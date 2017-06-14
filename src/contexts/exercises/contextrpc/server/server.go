package main

import (
	"contexts/exercises/contextrpc"
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *contextrpc.HelloArgs, reply *contextrpc.HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}
func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc/timeout/deadline", s)
	fmt.Println("Listening")
	time.Sleep(20 * time.Second)
	http.ListenAndServe(":1234", nil)
}
