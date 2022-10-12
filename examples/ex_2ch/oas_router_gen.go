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
	args := [3]string{}

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
			case 'a': // Prefix: "api/"
				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "captcha/"
					if l := len("captcha/"); len(elem) >= l && elem[0:l] == "captcha/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '2': // Prefix: "2chcaptcha/"
						if l := len("2chcaptcha/"); len(elem) >= l && elem[0:l] == "2chcaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptcha2chcaptchaIDGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 's': // Prefix: "show"
							if l := len("show"); len(elem) >= l && elem[0:l] == "show" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptcha2chcaptchaShowGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					case 'a': // Prefix: "app/id/"
						if l := len("app/id/"); len(elem) >= l && elem[0:l] == "app/id/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "public_key"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAPICaptchaAppIDPublicKeyGetRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'i': // Prefix: "invisible_recaptcha/"
						if l := len("invisible_recaptcha/"); len(elem) >= l && elem[0:l] == "invisible_recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaInvisibleRecaptchaIDGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'm': // Prefix: "mobile"
							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaInvisibleRecaptchaMobileGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					case 'r': // Prefix: "recaptcha/"
						if l := len("recaptcha/"); len(elem) >= l && elem[0:l] == "recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaRecaptchaIDGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'm': // Prefix: "mobile"
							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaRecaptchaMobileGetRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				case 'd': // Prefix: "dislike"
					if l := len("dislike"); len(elem) >= l && elem[0:l] == "dislike" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPIDislikeGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'l': // Prefix: "like"
					if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPILikeGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'm': // Prefix: "mobile/v2/"
					if l := len("mobile/v2/"); len(elem) >= l && elem[0:l] == "mobile/v2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "after/"
						if l := len("after/"); len(elem) >= l && elem[0:l] == "after/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "thread"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[1] = elem[:idx]
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

								// Param: "num"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleAPIMobileV2AfterBoardThreadNumGetRequest([3]string{
											args[0],
											args[1],
											args[2],
										}, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}
							}
						}
					case 'b': // Prefix: "boards"
						if l := len("boards"); len(elem) >= l && elem[0:l] == "boards" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAPIMobileV2BoardsGetRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'i': // Prefix: "info/"
						if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "thread"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPIMobileV2InfoBoardThreadGetRequest([2]string{
										args[0],
										args[1],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					case 'p': // Prefix: "post/"
						if l := len("post/"); len(elem) >= l && elem[0:l] == "post/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "num"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPIMobileV2PostBoardNumGetRequest([2]string{
										args[0],
										args[1],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'u': // Prefix: "user/"
				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
					case 'a': // Prefix: "asslogin"
						if l := len("asslogin"); len(elem) >= l && elem[0:l] == "asslogin" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleUserPassloginPostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'o': // Prefix: "osting"
						if l := len("osting"); len(elem) >= l && elem[0:l] == "osting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleUserPostingPostRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				case 'r': // Prefix: "report"
					if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleUserReportPostRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "POST")
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
	args        [3]string
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
		args = [3]string{}
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
			case 'a': // Prefix: "api/"
				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "captcha/"
					if l := len("captcha/"); len(elem) >= l && elem[0:l] == "captcha/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '2': // Prefix: "2chcaptcha/"
						if l := len("2chcaptcha/"); len(elem) >= l && elem[0:l] == "2chcaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptcha2chcaptchaIDGet
									r.name = "APICaptcha2chcaptchaIDGet"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 's': // Prefix: "show"
							if l := len("show"); len(elem) >= l && elem[0:l] == "show" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptcha2chcaptchaShowGet
									r.name = "APICaptcha2chcaptchaShowGet"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					case 'a': // Prefix: "app/id/"
						if l := len("app/id/"); len(elem) >= l && elem[0:l] == "app/id/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "public_key"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: APICaptchaAppIDPublicKeyGet
								r.name = "APICaptchaAppIDPublicKeyGet"
								r.operationID = ""
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'i': // Prefix: "invisible_recaptcha/"
						if l := len("invisible_recaptcha/"); len(elem) >= l && elem[0:l] == "invisible_recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptchaInvisibleRecaptchaIDGet
									r.name = "APICaptchaInvisibleRecaptchaIDGet"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 'm': // Prefix: "mobile"
							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptchaInvisibleRecaptchaMobileGet
									r.name = "APICaptchaInvisibleRecaptchaMobileGet"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					case 'r': // Prefix: "recaptcha/"
						if l := len("recaptcha/"); len(elem) >= l && elem[0:l] == "recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"
							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptchaRecaptchaIDGet
									r.name = "APICaptchaRecaptchaIDGet"
									r.operationID = ""
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						case 'm': // Prefix: "mobile"
							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APICaptchaRecaptchaMobileGet
									r.name = "APICaptchaRecaptchaMobileGet"
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
				case 'd': // Prefix: "dislike"
					if l := len("dislike"); len(elem) >= l && elem[0:l] == "dislike" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: APIDislikeGet
							r.name = "APIDislikeGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'l': // Prefix: "like"
					if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: APILikeGet
							r.name = "APILikeGet"
							r.operationID = ""
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'm': // Prefix: "mobile/v2/"
					if l := len("mobile/v2/"); len(elem) >= l && elem[0:l] == "mobile/v2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "after/"
						if l := len("after/"); len(elem) >= l && elem[0:l] == "after/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "thread"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[1] = elem[:idx]
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

								// Param: "num"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									switch method {
									case "GET":
										// Leaf: APIMobileV2AfterBoardThreadNumGet
										r.name = "APIMobileV2AfterBoardThreadNumGet"
										r.operationID = ""
										r.args = args
										r.count = 3
										return r, true
									default:
										return
									}
								}
							}
						}
					case 'b': // Prefix: "boards"
						if l := len("boards"); len(elem) >= l && elem[0:l] == "boards" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: APIMobileV2BoardsGet
								r.name = "APIMobileV2BoardsGet"
								r.operationID = ""
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'i': // Prefix: "info/"
						if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "thread"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APIMobileV2InfoBoardThreadGet
									r.name = "APIMobileV2InfoBoardThreadGet"
									r.operationID = ""
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}
						}
					case 'p': // Prefix: "post/"
						if l := len("post/"); len(elem) >= l && elem[0:l] == "post/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
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

							// Param: "num"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: APIMobileV2PostBoardNumGet
									r.name = "APIMobileV2PostBoardNumGet"
									r.operationID = ""
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "user/"
				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
					case 'a': // Prefix: "asslogin"
						if l := len("asslogin"); len(elem) >= l && elem[0:l] == "asslogin" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: UserPassloginPost
								r.name = "UserPassloginPost"
								r.operationID = ""
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'o': // Prefix: "osting"
						if l := len("osting"); len(elem) >= l && elem[0:l] == "osting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: UserPostingPost
								r.name = "UserPostingPost"
								r.operationID = ""
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'r': // Prefix: "report"
					if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: UserReportPost
							r.name = "UserReportPost"
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
		}
	}
	return r, false
}
