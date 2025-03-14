// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s AnyOfTest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.SizeLimit.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sizeLimit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s AnyOfTestSizeLimit) Validate() error {
	switch s.Type {
	case IntAnyOfTestSizeLimit:
		return nil // no validation needed
	case StringAnyOfTestSizeLimit:
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$"],
		}).Validate(string(s.String)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s ArrayTest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Required == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "required",
			Error: err,
		})
	}
	if err := func() error {
		if s.NullableOptional.Set {
			if err := func() error {
				if s.NullableOptional.Value == nil {
					return errors.New("nil is invalid value")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "nullable_optional",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Data) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        true,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Email)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     true,
			Regex:        nil,
		}).Validate(string(s.Hostname)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hostname",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^\\d-\\d$"],
		}).Validate(string(s.Format)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if err := func() error {
		if s.NullableEnum.Set {
			if err := func() error {
				if err := s.NullableEnum.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "nullable_enum",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s DefaultTest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Enum.Set {
			if err := func() error {
				if err := s.Enum.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "enum",
			Error: err,
		})
	}
	if err := func() error {
		if s.Email.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        true,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(s.Email.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if s.Hostname.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     true,
					Regex:        nil,
				}).Validate(string(s.Hostname.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hostname",
			Error: err,
		})
	}
	if err := func() error {
		if s.Format.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^\\d-\\d$"],
				}).Validate(string(s.Format.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s DefaultTestEnum) Validate() error {
	switch s {
	case "big":
		return nil
	case "smol":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s Hash) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    64,
			MinLengthSet: true,
			MaxLength:    64,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Hex)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s MapWithProperties) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.SubMap.Set {
			if err := func() error {
				if err := s.SubMap.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sub_map",
			Error: err,
		})
	}
	if err := func() error {
		if s.MapValidation.Set {
			if err := func() error {
				if err := s.MapValidation.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "map_validation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s NullableEnums) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.OnlyNullable.Value.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "only_nullable",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.OnlyNullValue.Value.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "only_null_value",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Both.Value.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "both",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s NullableEnumsBoth) Validate() error {
	switch s {
	case "asc":
		return nil
	case "desc":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s NullableEnumsOnlyNullValue) Validate() error {
	switch s {
	case "asc":
		return nil
	case "desc":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s NullableEnumsOnlyNullable) Validate() error {
	switch s {
	case "asc":
		return nil
	case "desc":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s OneOfBooleanSumNullables) Validate() error {
	switch s.Type {
	case BoolOneOfBooleanSumNullables:
		return nil // no validation needed
	case OneOfNullablesOneOfBooleanSumNullables:
		if err := s.OneOfNullables.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s OneOfBugs) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.OneOfMinusUUIDMinusIntMinusEnum.Set {
			if err := func() error {
				if err := s.OneOfMinusUUIDMinusIntMinusEnum.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "oneOf-uuid-int-enum",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s OneOfNullables) Validate() error {
	switch s.Type {
	case NullOneOfNullables:
		return nil // no validation needed
	case StringOneOfNullables:
		return nil // no validation needed
	case IntOneOfNullables:
		return nil // no validation needed
	case StringArrayOneOfNullables:
		if s.StringArray == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s OneOfUUIDAndIntEnum) Validate() error {
	switch s.Type {
	case UUIDOneOfUUIDAndIntEnum:
		return nil // no validation needed
	case OneOfUUIDAndIntEnum1OneOfUUIDAndIntEnum:
		if err := s.OneOfUUIDAndIntEnum1.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s OneOfUUIDAndIntEnum1) Validate() error {
	switch s {
	case 0:
		return nil
	case 1:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s OneOfWithNullable) Validate() error {
	switch s.Type {
	case NullOneOfWithNullable:
		return nil // no validation needed
	case StringOneOfWithNullable:
		return nil // no validation needed
	case IntOneOfWithNullable:
		return nil // no validation needed
	case StringArrayOneOfWithNullable:
		if s.StringArray == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s Pet) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Primary == nil {
			return nil // optional
		}
		if err := func() error {
			if err := s.Primary.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "pointer")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "primary",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           0,
			MaxSet:        true,
			Max:           100000,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.ID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    4,
			MinLengthSet: true,
			MaxLength:    24,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Name)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if err := func() error {
		if s.Type.Set {
			if err := func() error {
				if err := s.Type.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Kind.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "kind",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Array{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    255,
			MaxLengthSet: true,
		}).ValidateLength(len(s.Friends)); err != nil {
			return errors.Wrap(err, "array")
		}
		var failures []validate.FieldError
		for i, elem := range s.Friends {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "friends",
			Error: err,
		})
	}
	if err := func() error {
		if s.Next.Set {
			if err := func() error {
				if err := s.Next.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "next",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestInteger1.Set {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        false,
					Min:           0,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: true,
					MultipleOf:    10,
				}).Validate(int64(s.TestInteger1.Value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testInteger1",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestFloat1.Set {
			if err := func() error {
				if err := (validate.Float{
					MinSet:        true,
					Min:           15,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: true,
					MultipleOf:    ratMap["10"],
				}).Validate(float64(s.TestFloat1.Value)); err != nil {
					return errors.Wrap(err, "float")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testFloat1",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.TestArray1 {
			if err := func() error {
				if elem == nil {
					return errors.New("nil is invalid value")
				}
				if err := (validate.Array{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    16,
					MaxLengthSet: true,
				}).ValidateLength(len(elem)); err != nil {
					return errors.Wrap(err, "array")
				}
				var failures []validate.FieldError
				for i, elem := range elem {
					if err := func() error {
						if err := (validate.String{
							MinLength:    0,
							MinLengthSet: false,
							MaxLength:    255,
							MaxLengthSet: true,
							Email:        false,
							Hostname:     false,
							Regex:        nil,
						}).Validate(string(elem)); err != nil {
							return errors.Wrap(err, "string")
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testArray1",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestArray2.Set {
			if err := func() error {
				if err := s.TestArray2.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testArray2",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestMap.Set {
			if err := func() error {
				if err := s.TestMap.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testMap",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestMapWithProps.Set {
			if err := func() error {
				if err := s.TestMapWithProps.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testMapWithProps",
			Error: err,
		})
	}
	if err := func() error {
		if s.TestAnyOf.Set {
			if err := func() error {
				if err := s.TestAnyOf.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testAnyOf",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetKind) Validate() error {
	switch s {
	case "big":
		return nil
	case "smol":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetName) Validate() error {
	if err := (validate.String{
		MinLength:    6,
		MinLengthSet: true,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        nil,
	}).Validate(string(s)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}
func (s PetType) Validate() error {
	switch s {
	case "fifa":
		return nil
	case "fofa":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s RecursiveArray) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s StringMap) Validate() error {
	var failures []validate.FieldError
	for key, elem := range s {
		if err := func() error {
			if err := (validate.String{
				MinLength:    1,
				MinLengthSet: true,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        nil,
			}).Validate(string(elem)); err != nil {
				return errors.Wrap(err, "string")
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  key,
				Error: err,
			})
		}
	}

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s StringStringMap) Validate() error {
	var failures []validate.FieldError
	for key, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  key,
				Error: err,
			})
		}
	}

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TestFloatValidation) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{
			MinSet:        true,
			Min:           1.5,
			MaxSet:        true,
			Max:           2,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    nil,
		}).Validate(float64(s.Minmax)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "minmax",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{
			MinSet:        false,
			Min:           0,
			MaxSet:        false,
			Max:           0,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: true,
			MultipleOf:    ratMap["5"],
		}).Validate(float64(s.MultipleOf)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "multipleOf",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TestNullableOneofsApplicationJSONCreated) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s TestNullableOneofsApplicationJSONOK) Validate() error {
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}
func (s ValidationStringMap) Validate() error {
	var failures []validate.FieldError
	for key, elem := range s {
		if err := func() error {
			if err := (validate.String{
				MinLength:    1,
				MinLengthSet: true,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        nil,
			}).Validate(string(elem)); err != nil {
				return errors.Wrap(err, "string")
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  key,
				Error: err,
			})
		}
	}

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
