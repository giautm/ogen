// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

// WriteJSON implements json.Marshaler.
func (s HelloWorld) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes HelloWorld json value to io.Writer.
func (s HelloWorld) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads HelloWorld json value from io.Reader.
func (s *HelloWorld) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads HelloWorld from json stream.
func (s *HelloWorld) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s WorldObject) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("id")
	j.WriteInt64(s.ID)
	more.More()
	j.WriteObjectField("randomNumber")
	j.WriteInt64(s.RandomNumber)
	j.WriteObjectEnd()
}

// WriteJSONTo writes WorldObject json value to io.Writer.
func (s WorldObject) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads WorldObject json value from io.Reader.
func (s *WorldObject) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads WorldObject from json stream.
func (s *WorldObject) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "id":
			s.ID = i.ReadInt64()
			return i.Error == nil
		case "randomNumber":
			s.RandomNumber = i.ReadInt64()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
