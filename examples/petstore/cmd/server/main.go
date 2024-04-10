package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	petstorev3 "petstore/gen/petstore/v3"
	petstorev3connect "petstore/gen/petstore/v3/petstorev3connect"
	"petstore/gen/petstore/v3/petstorev3oas2connect"

	"connectrpc.com/connect"
)

type PetstoreServer struct {
	petstorev3connect.UnimplementedPetServiceHandler
}

func (s *PetstoreServer) FindPet(ctx context.Context, req *connect.Request[petstorev3.FindPetRequest]) (*connect.Response[petstorev3.FindPetResponse], error) {
	res := connect.NewResponse(&petstorev3.FindPetResponse{
		Pet: &petstorev3.Pet{
			Id:   1,
			Name: "pochi",
			Category: &petstorev3.Category{
				Id:   2,
				Name: "dog",
			},
			PhotoUrls: []string{},
			Tags:      []*petstorev3.Tag{},
			Status:    "healthy",
		},
	})

	return res, nil
}

func (s *PetstoreServer) AddPet(ctx context.Context, req *connect.Request[petstorev3.AddPetRequest]) (*connect.Response[petstorev3.AddPetResponse], error) {
	res := connect.NewResponse(&petstorev3.AddPetResponse{
		Pet: &petstorev3.Pet{
			Id:   0,
			Name: "tama",
			Category: &petstorev3.Category{
				Id:   1,
				Name: "cat",
			},
			PhotoUrls: []string{},
			Tags:      []*petstorev3.Tag{},
			Status:    "hungry",
		},
	})

	return res, nil
}

func main() {
	server := &PetstoreServer{}
	mux := http.NewServeMux()
	petstorev3oas2connect.RegisterPetServiceEndpoints(mux, server)

	port := 8080
	log.Printf("start to listen on %d", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux); err != nil {
		log.Printf("[ERROR] failed to listen: %s", err)
	}
}
