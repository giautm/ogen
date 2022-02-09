// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [5]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleErrorGetRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'e': // Prefix: "error"
				if l := len("error"); len(elem) >= l && elem[0:l] == "error" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: ErrorGet
					s.handleErrorGetRequest([0]string{}, w, r)

					return
				}
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: FoobarGet
					s.handleFoobarGetRequest([0]string{}, w, r)

					return
				}
			case 'n': // Prefix: "name/"
				if l := len("name/"); len(elem) >= l && elem[0:l] == "name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "foo"
					// Match until "1"
					idx := strings.IndexByte(elem, '1')
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '1': // Prefix: "1234"
						if l := len("1234"); len(elem) >= l && elem[0:l] == "1234" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "bar"
						// Match until "-"
						idx := strings.IndexByte(elem, '-')
						if idx < 0 {
							idx = len(elem)
						}
						args[2] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '-': // Prefix: "-"
							if l := len("-"); len(elem) >= l && elem[0:l] == "-" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "baz"
							// Match until "!"
							idx := strings.IndexByte(elem, '!')
							if idx < 0 {
								idx = len(elem)
							}
							args[3] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '!': // Prefix: "!"
								if l := len("!"); len(elem) >= l && elem[0:l] == "!" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "kek"
								// Leaf parameter
								args[4] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: DataGetFormat
									s.handleDataGetFormatRequest([5]string{
										args[0],
										args[1],
										args[2],
										args[3],
										args[4],
									}, w, r)

									return
								}
							}
						}
					}
				}
			case 'p': // Prefix: "pet"
				if l := len("pet"); len(elem) >= l && elem[0:l] == "pet" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePetGetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handlePetGetAvatarByIDRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetGetAvatarByID
							s.handlePetGetAvatarByIDRequest([0]string{}, w, r)

							return
						}
					case 'f': // Prefix: "friendNames/"
						if l := len("friendNames/"); len(elem) >= l && elem[0:l] == "friendNames/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: PetFriendsNamesByID
							s.handlePetFriendsNamesByIDRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					case 'n': // Prefix: "name/"
						if l := len("name/"); len(elem) >= l && elem[0:l] == "name/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: PetNameByID
							s.handlePetNameByIDRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
					// Param: "name"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handlePetGetByNameRequest([1]string{
							args[0],
						}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/avatar"
						if l := len("/avatar"); len(elem) >= l && elem[0:l] == "/avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetGetAvatarByName
							s.handlePetGetAvatarByNameRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
				}
			case 'r': // Prefix: "recursiveMap"
				if l := len("recursiveMap"); len(elem) >= l && elem[0:l] == "recursiveMap" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: RecursiveMapGet
					s.handleRecursiveMapGetRequest([0]string{}, w, r)

					return
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleTestObjectQueryParameterRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/header"
					if l := len("/header"); len(elem) >= l && elem[0:l] == "/header" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: GetHeader
						s.handleGetHeaderRequest([0]string{}, w, r)

						return
					}
				case 'O': // Prefix: "ObjectQueryParameter"
					if l := len("ObjectQueryParameter"); len(elem) >= l && elem[0:l] == "ObjectQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: TestObjectQueryParameter
						s.handleTestObjectQueryParameterRequest([0]string{}, w, r)

						return
					}
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleOneofBugRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: FoobarPost
					s.handleFoobarPostRequest([0]string{}, w, r)

					return
				}
			case 'o': // Prefix: "oneofBug"
				if l := len("oneofBug"); len(elem) >= l && elem[0:l] == "oneofBug" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: OneofBug
					s.handleOneofBugRequest([0]string{}, w, r)

					return
				}
			case 'p': // Prefix: "pet"
				if l := len("pet"); len(elem) >= l && elem[0:l] == "pet" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePetCreateRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handlePetUploadAvatarByIDRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetUploadAvatarByID
							s.handlePetUploadAvatarByIDRequest([0]string{}, w, r)

							return
						}
					case 'u': // Prefix: "updateName"
						if l := len("updateName"); len(elem) >= l && elem[0:l] == "updateName" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handlePetUpdateNamePostRequest([0]string{}, w, r)

							return
						}
						switch elem[0] {
						case 'A': // Prefix: "Alias"
							if l := len("Alias"); len(elem) >= l && elem[0:l] == "Alias" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: PetUpdateNameAliasPost
								s.handlePetUpdateNameAliasPostRequest([0]string{}, w, r)

								return
							}
						}
					}
				}
			}
		}
	case "PUT":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/foobar"
			if l := len("/foobar"); len(elem) >= l && elem[0:l] == "/foobar" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: FoobarPut
				s.handleFoobarPutRequest([0]string{}, w, r)

				return
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [5]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [5]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "ErrorGet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'e': // Prefix: "error"
				if l := len("error"); len(elem) >= l && elem[0:l] == "error" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: ErrorGet
					r.name = "ErrorGet"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: FoobarGet
					r.name = "FoobarGet"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'n': // Prefix: "name/"
				if l := len("name/"); len(elem) >= l && elem[0:l] == "name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "foo"
					// Match until "1"
					idx := strings.IndexByte(elem, '1')
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '1': // Prefix: "1234"
						if l := len("1234"); len(elem) >= l && elem[0:l] == "1234" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "bar"
						// Match until "-"
						idx := strings.IndexByte(elem, '-')
						if idx < 0 {
							idx = len(elem)
						}
						args[2] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '-': // Prefix: "-"
							if l := len("-"); len(elem) >= l && elem[0:l] == "-" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "baz"
							// Match until "!"
							idx := strings.IndexByte(elem, '!')
							if idx < 0 {
								idx = len(elem)
							}
							args[3] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '!': // Prefix: "!"
								if l := len("!"); len(elem) >= l && elem[0:l] == "!" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "kek"
								// Leaf parameter
								args[4] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: DataGetFormat
									r.name = "DataGetFormat"
									r.args = args
									r.count = 5
									return r, true
								}
							}
						}
					}
				}
			case 'p': // Prefix: "pet"
				if l := len("pet"); len(elem) >= l && elem[0:l] == "pet" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "PetGet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "PetGetAvatarByID"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetGetAvatarByID
							r.name = "PetGetAvatarByID"
							r.args = args
							r.count = 0
							return r, true
						}
					case 'f': // Prefix: "friendNames/"
						if l := len("friendNames/"); len(elem) >= l && elem[0:l] == "friendNames/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: PetFriendsNamesByID
							r.name = "PetFriendsNamesByID"
							r.args = args
							r.count = 1
							return r, true
						}
					case 'n': // Prefix: "name/"
						if l := len("name/"); len(elem) >= l && elem[0:l] == "name/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: PetNameByID
							r.name = "PetNameByID"
							r.args = args
							r.count = 1
							return r, true
						}
					}
					// Param: "name"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "PetGetByName"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/avatar"
						if l := len("/avatar"); len(elem) >= l && elem[0:l] == "/avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetGetAvatarByName
							r.name = "PetGetAvatarByName"
							r.args = args
							r.count = 1
							return r, true
						}
					}
				}
			case 'r': // Prefix: "recursiveMap"
				if l := len("recursiveMap"); len(elem) >= l && elem[0:l] == "recursiveMap" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: RecursiveMapGet
					r.name = "RecursiveMapGet"
					r.args = args
					r.count = 0
					return r, true
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "TestObjectQueryParameter"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/header"
					if l := len("/header"); len(elem) >= l && elem[0:l] == "/header" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: GetHeader
						r.name = "GetHeader"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'O': // Prefix: "ObjectQueryParameter"
					if l := len("ObjectQueryParameter"); len(elem) >= l && elem[0:l] == "ObjectQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: TestObjectQueryParameter
						r.name = "TestObjectQueryParameter"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "OneofBug"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: FoobarPost
					r.name = "FoobarPost"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'o': // Prefix: "oneofBug"
				if l := len("oneofBug"); len(elem) >= l && elem[0:l] == "oneofBug" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: OneofBug
					r.name = "OneofBug"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'p': // Prefix: "pet"
				if l := len("pet"); len(elem) >= l && elem[0:l] == "pet" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "PetCreate"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "PetUploadAvatarByID"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: PetUploadAvatarByID
							r.name = "PetUploadAvatarByID"
							r.args = args
							r.count = 0
							return r, true
						}
					case 'u': // Prefix: "updateName"
						if l := len("updateName"); len(elem) >= l && elem[0:l] == "updateName" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							r.name = "PetUpdateNamePost"
							r.args = args
							r.count = 0
							return r, true
						}
						switch elem[0] {
						case 'A': // Prefix: "Alias"
							if l := len("Alias"); len(elem) >= l && elem[0:l] == "Alias" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: PetUpdateNameAliasPost
								r.name = "PetUpdateNameAliasPost"
								r.args = args
								r.count = 0
								return r, true
							}
						}
					}
				}
			}
		}
	case "PUT":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/foobar"
			if l := len("/foobar"); len(elem) >= l && elem[0:l] == "/foobar" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: FoobarPut
				r.name = "FoobarPut"
				r.args = args
				r.count = 0
				return r, true
			}
		}
	}
	return r, false
}
