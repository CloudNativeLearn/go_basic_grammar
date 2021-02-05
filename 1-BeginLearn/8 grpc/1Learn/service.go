package main

import (
	service "./service"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	rpcService:=grpc.NewServer()
	service.RegisterProdServiceServer(rpcService,new(service.ProdService))

	//lis,_:= net.Listen("tcp",":8081")
	//rpcService.Serve(lis)

	mux:= http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcService.ServeHTTP(writer,request)
	})


	httpService:=&http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	httpService.ListenAndServe()




}