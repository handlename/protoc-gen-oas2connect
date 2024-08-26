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
	"go.akshayshah.org/connectproto"
	"google.golang.org/protobuf/encoding/protojson"
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
	pet := &petstorev3.Pet{
		Id:        req.Msg.GetId(),
		Name:      req.Msg.GetName(),
		PhotoUrls: []string{},
		Tags:      []*petstorev3.Tag{},
	}

	if req.Msg.Category != nil {
		pet.Category = &petstorev3.Category{
			Id:   req.Msg.GetCategory().GetId(),
			Name: req.Msg.GetCategory().GetName(),
		}
	}

	if req.Msg.PhotoUrls != nil {
		for _, p := range req.Msg.GetPhotoUrls() {
			pet.PhotoUrls = append(pet.PhotoUrls, p)
		}
	}

	if req.Msg.Tags != nil {
		for _, t := range req.Msg.GetTags() {
			pet.Tags = append(pet.Tags, &petstorev3.Tag{
				Id:   t.GetId(),
				Name: t.GetName(),
			})
		}
	}

	if req.Msg.Status != nil {
		pet.Status = req.Msg.GetStatus()
	}

	res := connect.NewResponse(&petstorev3.AddPetResponse{
		Pet: pet,
	})

	return res, nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	server := &PetstoreServer{}
	mux := http.NewServeMux()
	opt := connectproto.WithJSON(
		protojson.MarshalOptions{EmitDefaultValues: true},
		protojson.UnmarshalOptions{},
	)
	petstorev3oas2connect.RegisterPetServiceEndpoints(mux, server, loggingMiddleware, opt)

	port := 8080
	log.Printf("start to listen on %d", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux); err != nil {
		log.Printf("[ERROR] failed to listen: %s", err)
	}
}
