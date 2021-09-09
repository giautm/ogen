// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

func Register(r chi.Router, s Server) {
	r.Method("GET", "/foobar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, err := ParseFooBarGetParameters(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := s.FooBarGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(b))
	}))
	r.Method("POST", "/foobar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(Pet)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// Requets body is optional.
			if !errors.Is(err, io.EOF) {
				// Unexpected error.
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			req = nil
		}

		resp, err := s.FooBarPost(r.Context(), req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(b))
	}))
	r.Method("GET", "/pet", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, err := ParsePetGetParameters(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := s.PetGet(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(b))
	}))
	r.Method("POST", "/pet", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(Pet)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// Requets body is optional.
			if !errors.Is(err, io.EOF) {
				// Unexpected error.
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			req = nil
		}

		resp, err := s.PetCreate(r.Context(), req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(b))
	}))
	r.Method("GET", "/pet/{name}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, err := ParsePetGetByNameParameters(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := s.PetGetByName(r.Context(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(b))
	}))
}
