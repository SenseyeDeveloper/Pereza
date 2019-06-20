package core

import (
	"strings"
)

type FieldTags struct {
	Name      string
	Omit      bool
	OmitEmpty bool
	AsString  bool
}

// https://golang.org/pkg/encoding/json/
func ParseFieldTags(json string) FieldTags {
	switch json {
	case "":
		return FieldTags{}
	case "-":
		return FieldTags{
			Omit: true,
		}
	case "-,":
		return FieldTags{
			Name: "-",
		}
	}

	splitIndex := strings.IndexByte(json, ',')
	if splitIndex == -1 {
		return FieldTags{
			Name: json,
		}
	}

	settings := strings.Split(json, ",")

	result := FieldTags{}
	name := settings[0]
	if name == "-" {
		result.Omit = true
	} else {
		result.Name = name
	}

	length := len(settings)
	for i := 1; i < length; i++ {
		switch settings[i] {
		case "omitempty":
			result.OmitEmpty = true
		case "string":
			result.AsString = true
		}
	}

	return result
}

func EasyParseFieldTags(json string) FieldTags {
	result := FieldTags{}

	for i, s := range strings.Split(json, ",") {
		switch {
		case i == 0 && s == "-":
			result.OmitEmpty = true
		case i == 0:
			result.Name = s
		case s == "omitempty":
			result.OmitEmpty = true
		case s == "string":
			result.AsString = true
		}
	}

	return result
}

/**
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`

// The "string" option signals that a field is stored as JSON inside a JSON-encoded string.
// It applies only to fields of string, floating point, integer, or boolean types.
// This extra level of encoding is sometimes used when communicating with JavaScript programs
Int64String int64 `json:",string"`
*/
