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
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/pets"
			if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleListPetRequest([0]string{}, w, r)
				case "POST":
					s.handleCreatePetRequest([0]string{}, w, r)
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

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch r.Method {
					case "DELETE":
						s.handleDeletePetRequest([1]string{
							args[0],
						}, w, r)
					case "GET":
						s.handleReadPetRequest([1]string{
							args[0],
						}, w, r)
					case "PATCH":
						s.handleUpdatePetRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "DELETE,GET,PATCH")
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
					case 'c': // Prefix: "categories"
						if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListPetCategoriesRequest([1]string{
									args[0],
								}, w, r)
							case "POST":
								s.handleCreatePetCategoriesRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET,POST")
							}

							return
						}
					case 'f': // Prefix: "friends"
						if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListPetFriendsRequest([1]string{
									args[0],
								}, w, r)
							case "POST":
								s.handleCreatePetFriendsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET,POST")
							}

							return
						}
					case 'o': // Prefix: "owner"
						if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "DELETE":
								s.handleDeletePetOwnerRequest([1]string{
									args[0],
								}, w, r)
							case "GET":
								s.handleReadPetOwnerRequest([1]string{
									args[0],
								}, w, r)
							case "POST":
								s.handleCreatePetOwnerRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET,POST")
							}

							return
						}
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
	args        [1]string
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
		args = [1]string{}
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
		case '/': // Prefix: "/pets"
			if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = "ListPet"
					r.operationID = "listPet"
					r.args = args
					r.count = 0
					return r, true
				case "POST":
					r.name = "CreatePet"
					r.operationID = "createPet"
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

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch method {
					case "DELETE":
						r.name = "DeletePet"
						r.operationID = "deletePet"
						r.args = args
						r.count = 1
						return r, true
					case "GET":
						r.name = "ReadPet"
						r.operationID = "readPet"
						r.args = args
						r.count = 1
						return r, true
					case "PATCH":
						r.name = "UpdatePet"
						r.operationID = "updatePet"
						r.args = args
						r.count = 1
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
					case 'c': // Prefix: "categories"
						if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListPetCategories
								r.name = "ListPetCategories"
								r.operationID = "listPetCategories"
								r.args = args
								r.count = 1
								return r, true
							case "POST":
								// Leaf: CreatePetCategories
								r.name = "CreatePetCategories"
								r.operationID = "createPetCategories"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'f': // Prefix: "friends"
						if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListPetFriends
								r.name = "ListPetFriends"
								r.operationID = "listPetFriends"
								r.args = args
								r.count = 1
								return r, true
							case "POST":
								// Leaf: CreatePetFriends
								r.name = "CreatePetFriends"
								r.operationID = "createPetFriends"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'o': // Prefix: "owner"
						if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								// Leaf: DeletePetOwner
								r.name = "DeletePetOwner"
								r.operationID = "deletePetOwner"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								// Leaf: ReadPetOwner
								r.name = "ReadPetOwner"
								r.operationID = "readPetOwner"
								r.args = args
								r.count = 1
								return r, true
							case "POST":
								// Leaf: CreatePetOwner
								r.name = "CreatePetOwner"
								r.operationID = "createPetOwner"
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
		}
	}
	return r, false
}
