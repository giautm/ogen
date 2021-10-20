package ast

type Method struct {
	Name        string
	OperationID string
	PathParts   []PathPart
	HTTPMethod  string
	Parameters  []*Parameter

	RequestType Type
	RequestBody *RequestBody

	ResponseType Type
	Responses    *MethodResponse
}

type PathPart struct {
	Raw   string
	Param *Parameter
}

func (m *Method) RawPath() (path string) {
	for _, part := range m.PathParts {
		if part.Raw != "" {
			path += part.Raw
			continue
		}
		path += "{" + part.Param.SourceName + "}"
	}
	return
}

func (m *Method) PathParams() []*Parameter   { return m.getParams(LocationPath) }
func (m *Method) QueryParams() []*Parameter  { return m.getParams(LocationQuery) }
func (m *Method) CookieParams() []*Parameter { return m.getParams(LocationCookie) }
func (m *Method) HeaderParams() []*Parameter { return m.getParams(LocationHeader) }

func (m *Method) getParams(locatedIn ParameterLocation) []*Parameter {
	var params []*Parameter
	for _, p := range m.Parameters {
		if p.In == locatedIn {
			params = append(params, p)
		}
	}
	return params
}

func (m *Method) IsRequestIface() bool {
	_, ok := m.RequestType.(*Interface)
	return ok
}

func (m *Method) IsResponseIface() bool {
	_, ok := m.ResponseType.(*Interface)
	return ok
}
