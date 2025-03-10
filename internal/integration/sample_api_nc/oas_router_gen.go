// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"strings"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [5]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
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
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "defaultTest"
				if l := len("defaultTest"); len(elem) >= l && elem[0:l] == "defaultTest" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleDefaultTestRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			case 'e': // Prefix: "error"
				if l := len("error"); len(elem) >= l && elem[0:l] == "error" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleErrorGetRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleFoobarGetRequest([0]string{}, w, r)
					case "POST":
						s.handleFoobarPostRequest([0]string{}, w, r)
					case "PUT":
						s.handleFoobarPutRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST,PUT")
					}

					return
				}
			case 'n': // Prefix: "n"
				if l := len("n"); len(elem) >= l && elem[0:l] == "n" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "ame/"
					if l := len("ame/"); len(elem) >= l && elem[0:l] == "ame/" {
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
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleDataGetFormatRequest([5]string{
												args[0],
												args[1],
												args[2],
												args[3],
												args[4],
											}, w, r)
										default:
											s.notAllowed(w, r, "GET")
										}

										return
									}
								}
							}
						}
					}
				case 'o': // Prefix: "oAdditionalPropertiesTest"
					if l := len("oAdditionalPropertiesTest"); len(elem) >= l && elem[0:l] == "oAdditionalPropertiesTest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleNoAdditionalPropertiesTestRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'u': // Prefix: "ullableDefaultResponse"
					if l := len("ullableDefaultResponse"); len(elem) >= l && elem[0:l] == "ullableDefaultResponse" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleNullableDefaultResponseRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'o': // Prefix: "oneofBug"
				if l := len("oneofBug"); len(elem) >= l && elem[0:l] == "oneofBug" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleOneofBugRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "atternRecursiveMap"
					if l := len("atternRecursiveMap"); len(elem) >= l && elem[0:l] == "atternRecursiveMap" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handlePatternRecursiveMapGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'e': // Prefix: "et"
					if l := len("et"); len(elem) >= l && elem[0:l] == "et" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handlePetGetRequest([0]string{}, w, r)
						case "POST":
							s.handlePetCreateRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

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
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "avatar"
							if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handlePetGetAvatarByIDRequest([0]string{}, w, r)
								case "POST":
									s.handlePetUploadAvatarByIDRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET,POST")
								}

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
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handlePetFriendsNamesByIDRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

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
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handlePetNameByIDRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'u': // Prefix: "updateName"
							if l := len("updateName"); len(elem) >= l && elem[0:l] == "updateName" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch r.Method {
								case "POST":
									s.handlePetUpdateNamePostRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

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
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handlePetUpdateNameAliasPostRequest([0]string{}, w, r)
									default:
										s.notAllowed(w, r, "POST")
									}

									return
								}
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
							switch r.Method {
							case "GET":
								s.handlePetGetByNameRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

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
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handlePetGetAvatarByNameRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'r': // Prefix: "recursive"
				if l := len("recursive"); len(elem) >= l && elem[0:l] == "recursive" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'A': // Prefix: "Array"
					if l := len("Array"); len(elem) >= l && elem[0:l] == "Array" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRecursiveArrayGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'M': // Prefix: "Map"
					if l := len("Map"); len(elem) >= l && elem[0:l] == "Map" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRecursiveMapGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 's': // Prefix: "s"
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "ecurityTest"
					if l := len("ecurityTest"); len(elem) >= l && elem[0:l] == "ecurityTest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleSecurityTestRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 't': // Prefix: "tringIntMap"
					if l := len("tringIntMap"); len(elem) >= l && elem[0:l] == "tringIntMap" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleStringIntMapGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/header"
					if l := len("/header"); len(elem) >= l && elem[0:l] == "/header" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetHeaderRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'C': // Prefix: "ContentParameter"
					if l := len("ContentParameter"); len(elem) >= l && elem[0:l] == "ContentParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleTestContentParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'F': // Prefix: "FloatValidation"
					if l := len("FloatValidation"); len(elem) >= l && elem[0:l] == "FloatValidation" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleTestFloatValidationRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'N': // Prefix: "NullableOneofs"
					if l := len("NullableOneofs"); len(elem) >= l && elem[0:l] == "NullableOneofs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleTestNullableOneofsRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'O': // Prefix: "ObjectQueryParameter"
					if l := len("ObjectQueryParameter"); len(elem) >= l && elem[0:l] == "ObjectQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleTestObjectQueryParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [5]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
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
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
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
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "defaultTest"
				if l := len("defaultTest"); len(elem) >= l && elem[0:l] == "defaultTest" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: DefaultTest
						r.name = "DefaultTest"
						r.operationID = "defaultTest"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'e': // Prefix: "error"
				if l := len("error"); len(elem) >= l && elem[0:l] == "error" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: ErrorGet
						r.name = "ErrorGet"
						r.operationID = "errorGet"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'f': // Prefix: "foobar"
				if l := len("foobar"); len(elem) >= l && elem[0:l] == "foobar" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: FoobarGet
						r.name = "FoobarGet"
						r.operationID = "foobarGet"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						// Leaf: FoobarPost
						r.name = "FoobarPost"
						r.operationID = "foobarPost"
						r.args = args
						r.count = 0
						return r, true
					case "PUT":
						// Leaf: FoobarPut
						r.name = "FoobarPut"
						r.operationID = ""
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'n': // Prefix: "n"
				if l := len("n"); len(elem) >= l && elem[0:l] == "n" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "ame/"
					if l := len("ame/"); len(elem) >= l && elem[0:l] == "ame/" {
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
										switch method {
										case "GET":
											// Leaf: DataGetFormat
											r.name = "DataGetFormat"
											r.operationID = "dataGetFormat"
											r.args = args
											r.count = 5
											return r, true
										default:
											return
										}
									}
								}
							}
						}
					}
				case 'o': // Prefix: "oAdditionalPropertiesTest"
					if l := len("oAdditionalPropertiesTest"); len(elem) >= l && elem[0:l] == "oAdditionalPropertiesTest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: NoAdditionalPropertiesTest
							r.name = "NoAdditionalPropertiesTest"
							r.operationID = "noAdditionalPropertiesTest"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'u': // Prefix: "ullableDefaultResponse"
					if l := len("ullableDefaultResponse"); len(elem) >= l && elem[0:l] == "ullableDefaultResponse" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: NullableDefaultResponse
							r.name = "NullableDefaultResponse"
							r.operationID = "nullableDefaultResponse"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'o': // Prefix: "oneofBug"
				if l := len("oneofBug"); len(elem) >= l && elem[0:l] == "oneofBug" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: OneofBug
						r.name = "OneofBug"
						r.operationID = "oneofBug"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "atternRecursiveMap"
					if l := len("atternRecursiveMap"); len(elem) >= l && elem[0:l] == "atternRecursiveMap" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: PatternRecursiveMapGet
							r.name = "PatternRecursiveMapGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'e': // Prefix: "et"
					if l := len("et"); len(elem) >= l && elem[0:l] == "et" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "PetGet"
							r.operationID = "petGet"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "PetCreate"
							r.operationID = "petCreate"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "avatar"
							if l := len("avatar"); len(elem) >= l && elem[0:l] == "avatar" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: PetGetAvatarByID
									r.name = "PetGetAvatarByID"
									r.operationID = "petGetAvatarByID"
									r.args = args
									r.count = 0
									return r, true
								case "POST":
									// Leaf: PetUploadAvatarByID
									r.name = "PetUploadAvatarByID"
									r.operationID = "petUploadAvatarByID"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
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
								switch method {
								case "GET":
									// Leaf: PetFriendsNamesByID
									r.name = "PetFriendsNamesByID"
									r.operationID = "petFriendsNamesByID"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
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
								switch method {
								case "GET":
									// Leaf: PetNameByID
									r.name = "PetNameByID"
									r.operationID = "petNameByID"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'u': // Prefix: "updateName"
							if l := len("updateName"); len(elem) >= l && elem[0:l] == "updateName" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									r.name = "PetUpdateNamePost"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case 'A': // Prefix: "Alias"
								if l := len("Alias"); len(elem) >= l && elem[0:l] == "Alias" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										// Leaf: PetUpdateNameAliasPost
										r.name = "PetUpdateNameAliasPost"
										r.operationID = ""
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
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
							switch method {
							case "GET":
								r.name = "PetGetByName"
								r.operationID = "petGetByName"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/avatar"
							if l := len("/avatar"); len(elem) >= l && elem[0:l] == "/avatar" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: PetGetAvatarByName
									r.name = "PetGetAvatarByName"
									r.operationID = "petGetAvatarByName"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'r': // Prefix: "recursive"
				if l := len("recursive"); len(elem) >= l && elem[0:l] == "recursive" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'A': // Prefix: "Array"
					if l := len("Array"); len(elem) >= l && elem[0:l] == "Array" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: RecursiveArrayGet
							r.name = "RecursiveArrayGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'M': // Prefix: "Map"
					if l := len("Map"); len(elem) >= l && elem[0:l] == "Map" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: RecursiveMapGet
							r.name = "RecursiveMapGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 's': // Prefix: "s"
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "ecurityTest"
					if l := len("ecurityTest"); len(elem) >= l && elem[0:l] == "ecurityTest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: SecurityTest
							r.name = "SecurityTest"
							r.operationID = "securityTest"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 't': // Prefix: "tringIntMap"
					if l := len("tringIntMap"); len(elem) >= l && elem[0:l] == "tringIntMap" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: StringIntMapGet
							r.name = "StringIntMapGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 't': // Prefix: "test"
				if l := len("test"); len(elem) >= l && elem[0:l] == "test" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/header"
					if l := len("/header"); len(elem) >= l && elem[0:l] == "/header" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetHeader
							r.name = "GetHeader"
							r.operationID = "getHeader"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'C': // Prefix: "ContentParameter"
					if l := len("ContentParameter"); len(elem) >= l && elem[0:l] == "ContentParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: TestContentParameter
							r.name = "TestContentParameter"
							r.operationID = "testContentParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'F': // Prefix: "FloatValidation"
					if l := len("FloatValidation"); len(elem) >= l && elem[0:l] == "FloatValidation" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: TestFloatValidation
							r.name = "TestFloatValidation"
							r.operationID = "testFloatValidation"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'N': // Prefix: "NullableOneofs"
					if l := len("NullableOneofs"); len(elem) >= l && elem[0:l] == "NullableOneofs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: TestNullableOneofs
							r.name = "TestNullableOneofs"
							r.operationID = "testNullableOneofs"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'O': // Prefix: "ObjectQueryParameter"
					if l := len("ObjectQueryParameter"); len(elem) >= l && elem[0:l] == "ObjectQueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: TestObjectQueryParameter
							r.name = "TestObjectQueryParameter"
							r.operationID = "testObjectQueryParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
