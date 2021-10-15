// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/uri"
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
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = math.Mod
)

func (s *Pet) validate() error {
	// Validate 'Friends' field.
	if s.Friends != nil {
		if len(*s.Friends) > 255 {
			return fmt.Errorf("field 'Friends': array must contain at most 255 items")
		}
		for i, item := range *s.Friends {
			if err := func() error {
				if err := item.validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return fmt.Errorf("field 'Friends': item %d: %w", i, err)
			}
		}
	}

	// Validate 'ID' field.
	if s.ID < 0 {
		return fmt.Errorf("field 'ID': value must be greater than or equal to 0")
	}
	if s.ID > 100000 {
		return fmt.Errorf("field 'ID': value must be less than or equal to 100000")
	}

	// Validate 'Name' field.
	if len([]rune(s.Name)) < 4 {
		return fmt.Errorf("field 'Name': string must contain at least 4 characters")
	}
	if len([]rune(s.Name)) > 24 {
		return fmt.Errorf("field 'Name': string must contain at most 24 characters")
	}

	// Validate 'TestArray1' field.
	if s.TestArray1 != nil {
		for i, item := range *s.TestArray1 {
			if err := func() error {
				if len(item) > 16 {
					return fmt.Errorf("array must contain at most 16 items")
				}
				for i, item := range item {
					if err := func() error {
						if len([]rune(item)) > 255 {
							return fmt.Errorf("string must contain at most 255 characters")
						}
						return nil
					}(); err != nil {
						return fmt.Errorf("item %d: %w", i, err)
					}
				}
				return nil
			}(); err != nil {
				return fmt.Errorf("field 'TestArray1': item %d: %w", i, err)
			}
		}
	}

	// Validate 'TestFloat1' field.
	if s.TestFloat1 != nil {
		if *s.TestFloat1 < 15 {
			return fmt.Errorf("field 'TestFloat1': value must be greater than or equal to 15")
		}
		if math.Mod(float64(*s.TestFloat1), 10) != 0 {
			return fmt.Errorf("field 'TestFloat1': value must be multiple of 10")
		}
	}

	// Validate 'TestInteger1' field.
	if s.TestInteger1 != nil {
		if *s.TestInteger1%10 != 0 {
			return fmt.Errorf("field 'TestInteger1': value must be multiple of 10")
		}
	}

	// Validate 'Type' field.
	if s.Type != nil {
		switch *s.Type {
		case "fifa":
		case "fofa":
		default:
			return fmt.Errorf("field 'Type': invalid enum value: %v", *s.Type)
		}
	}

	return nil
}
