package validate

import (
	"regexp"
	"unicode"

	"github.com/ogen-go/errors"
)

// String validator.
type String struct {
	MinLength    int
	MinLengthSet bool
	MaxLength    int
	MaxLengthSet bool
	Email        bool
	Regex        *regexp.Regexp
	Hostname     bool
}

func (t *String) SetMaxLength(v int) {
	t.MaxLengthSet = true
	t.MaxLength = v
}

func (t *String) SetMinLength(v int) {
	t.MinLengthSet = true
	t.MinLength = v
}

func (t String) Set() bool {
	return t.MaxLengthSet || t.MinLengthSet || t.Email || t.Regex != nil || t.Hostname
}

func (t String) checkHostname(v string) error {
	if v == "" {
		return errors.New("blank")
	}
	for i, r := range v {
		if !unicode.IsPrint(r) {
			return errors.New("not printable character")
		}
		if unicode.IsSpace(r) {
			return errors.New("space character")
		}
		if i > 255 {
			return errors.New("too long")
		}
		if r == '.' {
			continue
		}
		if !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '-' || r >= 'A' && r <= 'Z') {
			return errors.New("invalid character")
		}
	}
	return nil
}

func (t String) checkEmail(v string) error {
	// Pretty basic validation, but should work for most cases and is not
	// too strict to break things.
	//
	// Still better than obscure regex or std `mail.ParseAddress`.
	var (
		gotAt bool
		last  rune
	)
	for i, r := range v {
		if !unicode.IsPrint(r) {
			return errors.New("not printable character")
		}
		if unicode.IsSpace(r) {
			return errors.New("space character")
		}

		last = r
		if r != '@' {
			continue
		}
		if gotAt {
			return errors.New(`got @ multiple times`)
		}
		if i == 0 {
			return errors.New(`got @ at start`)
		}
		gotAt = true
	}
	if last == '@' {
		return errors.New("@ at end")
	}
	if !gotAt {
		return errors.New(`no @`)
	}
	return nil
}

func (t String) Validate(v string) error {
	if err := (Array{
		MinLength:    t.MinLength,
		MinLengthSet: t.MaxLengthSet,
		MaxLength:    t.MaxLength,
		MaxLengthSet: t.MaxLengthSet,
	}).ValidateLength(len([]rune(v))); err != nil {
		return err
	}
	if t.Email {
		if err := t.checkEmail(v); err != nil {
			return err
		}
	}
	if t.Hostname {
		if err := t.checkHostname(v); err != nil {
			return err
		}
	}
	if t.Regex != nil {
		if !t.Regex.MatchString(v) {
			return errors.New("no regex match")
		}
	}
	return nil
}
