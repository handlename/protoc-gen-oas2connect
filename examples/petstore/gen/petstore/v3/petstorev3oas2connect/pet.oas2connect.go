// Code generated by protoc-gen-oas2connect. DO NOT EDIT.

package petstorev3oas2connect

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	connect "petstore/gen/petstore/v3/petstorev3connect"
	pb "petstore/gen/petstore/v3"
)

func RegisterPetServiceEndpoints(mux *http.ServeMux, svc connect.PetServiceHandler) {
	path, handler := connect.NewPetServiceHandler(svc)
	mux.HandleFunc(NewPetServiceUpdatePetHandler(path, handler))
	mux.HandleFunc(NewPetServiceAddPetHandler(path, handler))
	mux.HandleFunc(NewPetServiceFindPetsByStatusHandler(path, handler))
	mux.HandleFunc(NewPetServiceFindPetsByTagsHandler(path, handler))
	mux.HandleFunc(NewPetServiceFindPetHandler(path, handler))
	mux.HandleFunc(NewPetServiceDeletePetHandler(path, handler))
	log.Printf("registered 6 endpoints of PetService")
}
func NewPetServiceUpdatePetHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "PUT /pet", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.UpdatePetRequest{}
		dec := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := dec.Decode(&pbr); err != nil {
			log.Printf("failed to decode request body: %v", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"UpdatePet", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
func NewPetServiceAddPetHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "POST /pet", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.AddPetRequest{}
		dec := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := dec.Decode(&pbr); err != nil {
			log.Printf("failed to decode request body: %v", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"AddPet", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
func NewPetServiceFindPetsByStatusHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "GET /pet/findByStatus", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.FindPetsByStatusRequest{}

		if v := r.URL.Query().Get("status"); v != "" {
			if vv, err := ToString(v); err != nil {
				log.Printf("failed to convert status=%s: %v", v, err)
				http.Error(w, "invalid status", http.StatusBadRequest)
				return
			} else {
				pbr.Status = &vv
			}
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"FindPetsByStatus", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
func NewPetServiceFindPetsByTagsHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "GET /pet/findByTags", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.FindPetsByTagsRequest{}

		if rawParams, ok := r.URL.Query()["tags"]; ok {
			params := make([]string, len(rawParams))
			for _, rp := range rawParams {
				v, err := ToString(rp)
				if err != nil {
					http.Error(w, "invalid tags", http.StatusBadRequest)
					return
				}
				params = append(params, v)
			}
			pbr.Tags = params
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"FindPetsByTags", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
func NewPetServiceFindPetHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "GET /pet/{pet_id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.FindPetRequest{}

		if v, err := ToInt64(r.PathValue("pet_id")); err != nil {
			log.Printf("failed to convert pet_id=%s: %v", v, err)
			http.Error(w, "invalid pet_id", http.StatusBadRequest)
			return
		} else {
			pbr.PetId = v
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"FindPet", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
func NewPetServiceDeletePetHandler(protoPath string, protoHandler http.Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return "DELETE /pet/{pet_id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbr := pb.DeletePetRequest{}

		if v, err := ToInt64(r.PathValue("pet_id")); err != nil {
			log.Printf("failed to convert pet_id=%s: %v", v, err)
			http.Error(w, "invalid pet_id", http.StatusBadRequest)
			return
		} else {
			pbr.PetId = v
		}

		var cb bytes.Buffer
		json.NewEncoder(&cb).Encode(&pbr)

		cr, err := http.NewRequest(http.MethodPost, protoPath+"DeletePet", &cb)
		if err != nil {
			log.Printf("error on pass request to connect: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cr.Header.Set("content-type", "application/json")

		protoHandler.ServeHTTP(w, cr)
	})
}
