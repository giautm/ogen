// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreatePetResponse(response CreatePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreatePetCategoriesResponse(response CreatePetCategoriesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetCategoriesCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreatePetFriendsResponse(response CreatePetFriendsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetFriendsCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreatePetOwnerResponse(response CreatePetOwnerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetOwnerCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeletePetResponse(response DeletePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeletePetNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeletePetOwnerResponse(response DeletePetOwnerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeletePetOwnerNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListPetResponse(response ListPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListPetCategoriesResponse(response ListPetCategoriesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetCategoriesOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListPetFriendsResponse(response ListPetFriendsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListPetFriendsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeReadPetResponse(response ReadPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeReadPetOwnerResponse(response ReadPetOwnerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetOwnerRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdatePetResponse(response UpdatePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
