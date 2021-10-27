// Code generated by ogen, DO NOT EDIT.

package api

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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL *url.URL
	http      HTTPClient
}

func NewClient(serverURL string) *Client {
	u, err := url.Parse(serverURL)
	if err != nil {
		panic(err) // TODO: fix
	}
	return &Client{
		serverURL: u,
		http: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Client) FoobarGet(ctx context.Context, params FoobarGetParams) (res FoobarGetRes, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/foobar"

	q := u.Query()
	{
		// Encode "inlinedParam" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.InlinedParam
		param := e.EncodeInt64(v)
		q.Set("inlinedParam", param)
	}
	{
		// Encode "skip" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Skip
		param := e.EncodeInt32(v)
		q.Set("skip", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarGetResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPost(ctx context.Context, req Pet) (res FoobarPostRes, err error) {
	buf, contentType, err := encodeFoobarPostRequest(req)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/foobar"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPostResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPut(ctx context.Context) (res FoobarPutDefStatusCode, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/foobar"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPutResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetCreate(ctx context.Context, req PetCreateReq) (res Pet, err error) {
	buf, contentType, err := encodePetCreateRequest(req)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetCreateResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) (res []string, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/pet/friendNames/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeInt(params.ID)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetFriendsNamesByIDResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGet(ctx context.Context, params PetGetParams) (res PetGetRes, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/pet"

	q := u.Query()
	{
		// Encode "petID" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.PetID
		param := e.EncodeInt64(v)
		q.Set("petID", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	{
		value := conv.UUIDArrayToString(params.XTags)
		for _, v := range value {
			r.Header.Add("x-tags", v)
		}
	}
	{
		value := conv.StringArrayToString(params.XScope)
		for _, v := range value {
			r.Header.Add("x-scope", v)
		}
	}

	{
		value := conv.StringToString(params.Token)
		r.AddCookie(&http.Cookie{
			Name:   "token",
			Value:  value,
			MaxAge: 1337,
		})
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetGetResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGetByName(ctx context.Context, params PetGetByNameParams) (res Pet, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/pet/"
	{
		// Encode "name" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "name",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeString(params.Name)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetGetByNameResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetNameByID(ctx context.Context, params PetNameByIDParams) (res string, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/pet/name/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeInt(params.ID)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetNameByIDResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetUpdateNameAliasPost(ctx context.Context, req PetName) (res PetUpdateNameAliasPostDefStatusCode, err error) {
	buf, contentType, err := encodePetUpdateNameAliasPostRequest(req)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateNameAlias"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNameAliasPostResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetUpdateNamePost(ctx context.Context, req string) (res PetUpdateNamePostDefStatusCode, err error) {
	buf, contentType, err := encodePetUpdateNamePostRequest(req)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateName"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNamePostResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}
