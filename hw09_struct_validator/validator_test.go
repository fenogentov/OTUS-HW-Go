package hw09structvalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		//		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	TestStringStruct struct {
		StrLen           string   `validate:"len:10"`
		SliceString      []string `validate:"len:10"`
		StringRegexp     string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		StringNumbersLen string   `validate:"regexp:\\d+|len:7"`
		StringList       string   `validate:"in:foo,bar,tst"`
		StringMinMax     string   `validate:"min:5|max:10"`
	}

	TestIntStruct struct {
		IntMinMax int   `validate:"min:10|max:100"`
		IntList   int   `validate:"in:100,200,300"`
		SliceInt  []int `validate:"min:100|max:1000"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			Response{
				Code: 404,
				Body: "qasdfgh",
			},
			nil,
		},
		{
			App{
				Version: "qazxs",
			},
			nil,
		},
		{
			Token{
				Header:    nil,
				Payload:   []byte("stdgd"),
				Signature: []byte{2, 3},
			},
			nil,
		},
		{
			TestStringStruct{
				StrLen:           "qwertyuiop",
				SliceString:      []string{"qwertyuiop", "asdfghjklq", "zxcvbnmasd"},
				StringRegexp:     "qwerty@asdfg.com",
				StringNumbersLen: "1234567",
				StringList:       "tst",
				StringMinMax:     "qazxswe",
			},
			nil,
		},
		{
			TestIntStruct{
				IntMinMax: 50,
				IntList:   200,
				SliceInt:  []int{200, 300, 400},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run("validate without errors string fields structure", func(t *testing.T) {
			tt := tt
			//			t.Parallel()
			err := Validate(tt.in)
			require.NoError(t, err)
			_ = tt
		})
	}
}

func TestValidateError(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
		err         string
	}{
		{
			TestStringStruct{
				StrLen:           "qwertyuio",
				SliceString:      []string{"qwertyuio", "asdfghjklq", "zxxcvbnmasd"},
				StringRegexp:     "qwertyasdfg.com",
				StringNumbersLen: "123456",
				StringList:       "sts",
				StringMinMax:     "qwe",
			},
			ValidationErrors{
				ValidationError{
					Field: "StrLen",
					Err:   fmt.Errorf("error string length (len:10)"),
				},
				{
					Field: "SliceString",
					Err:   fmt.Errorf("error string length (len:10)"),
				},
				{
					Field: "SliceString",
					Err:   fmt.Errorf("error string length (len:10)"),
				},
				{
					Field: "StringRegexp",
					Err:   fmt.Errorf("error regular expression (^\\w+@\\w+\\.\\w+$)"),
				},
				{
					Field: "StringNumbersLen",
					Err:   fmt.Errorf("error string length (\\d+|len:7)"),
				},
				{
					Field: "StringList",
					Err:   fmt.Errorf("error not in the list (foo,bar,tst)"),
				},
				{
					Field: "StringMinMax",
					Err:   fmt.Errorf("error maximum string length (min:5|max:10)"),
				},
			},
			`0. field StrLen : error string length (len:10)
1. field SliceString : error string length (len:10)
2. field SliceString : error string length (len:10)
3. field StringRegexp : error regular expression (^\w+@\w+\.\w+$)
4. field StringNumbersLen : error string length (\d+|len:7)
5. field StringList : error not in the list (foo,bar,tst)
6. field StringMinMax : error maximum string length (min:5|max:10)
`,
		},
		{
			TestIntStruct{
				IntMinMax: 5,
				IntList:   50,
				SliceInt:  []int{50, 300, 400},
			},
			ValidationErrors{
				ValidationError{
					Field: "IntMinMax",
					Err:   fmt.Errorf("error minimum value (min:10|max:100)"),
				},
				{
					Field: "IntList",
					Err:   fmt.Errorf("error not in the list"),
				},
				{
					Field: "SliceInt",
					Err:   fmt.Errorf("error minimum value (min:100|max:1000)"),
				},
			},
			`0. field IntMinMax : error minimum value (min:10|max:100)
1. field IntList : error not in the list
2. field SliceInt : error minimum value (min:100|max:1000)
`,
		},
	}

	for _, tt := range tests {
		t.Run("validate with errors string fields structure", func(t *testing.T) {
			tt := tt
			//			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.err, err.Error())
			_ = tt
		})
	}

	t.Run("invalid input type", func(t *testing.T) {
		err := Validate(1)
		require.Error(t, err, ErrNoStruct)
	})
}
