package json_test

import (
	"bytes"
	"testing"

	"github.com/goccy/go-json"
)

func TestCoverNumber(t *testing.T) {
	type structNumber struct {
		A json.Number `json:"a"`
	}
	type structNumberOmitEmpty struct {
		A json.Number `json:"a,omitempty"`
	}
	type structNumberString struct {
		A json.Number `json:"a,string"`
	}

	type structNumberPtr struct {
		A *json.Number `json:"a"`
	}
	type structNumberPtrOmitEmpty struct {
		A *json.Number `json:"a,omitempty"`
	}
	type structNumberPtrString struct {
		A *json.Number `json:"a,string"`
	}

	tests := []struct {
		name string
		data interface{}
	}{
		{
			name: "Number",
			data: json.Number("10"),
		},
		{
			name: "NumberPtr",
			data: numberptr("10"),
		},
		{
			name: "NumberPtr3",
			data: numberptr3("10"),
		},
		{
			name: "NumberPtrNil",
			data: (*json.Number)(nil),
		},
		{
			name: "NumberPtr3Nil",
			data: (***json.Number)(nil),
		},

		// HeadNumberZero
		{
			name: "HeadNumberZero",
			data: struct {
				A json.Number `json:"a"`
			}{},
		},
		{
			name: "HeadNumberZeroOmitEmpty",
			data: struct {
				A json.Number `json:"a,omitempty"`
			}{},
		},
		{
			name: "HeadNumberZeroString",
			data: struct {
				A json.Number `json:"a,string"`
			}{},
		},

		// HeadNumber
		{
			name: "HeadNumber",
			data: struct {
				A json.Number `json:"a"`
			}{A: "1"},
		},
		{
			name: "HeadNumberOmitEmpty",
			data: struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"},
		},
		{
			name: "HeadNumberString",
			data: struct {
				A json.Number `json:"a,string"`
			}{A: "1"},
		},

		// HeadNumberPtr
		{
			name: "HeadNumberPtr",
			data: struct {
				A *json.Number `json:"a"`
			}{A: numberptr("1")},
		},
		{
			name: "HeadNumberPtrOmitEmpty",
			data: struct {
				A *json.Number `json:"a,omitempty"`
			}{A: numberptr("1")},
		},
		{
			name: "HeadNumberPtrString",
			data: struct {
				A *json.Number `json:"a,string"`
			}{A: numberptr("1")},
		},

		// HeadNumberPtrNil
		{
			name: "HeadNumberPtrNil",
			data: struct {
				A *json.Number `json:"a"`
			}{A: nil},
		},
		{
			name: "HeadNumberPtrNilOmitEmpty",
			data: struct {
				A *json.Number `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "HeadNumberPtrNilString",
			data: struct {
				A *json.Number `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadNumberZero
		{
			name: "PtrHeadNumberZero",
			data: &struct {
				A json.Number `json:"a"`
			}{},
		},
		{
			name: "PtrHeadNumberZeroOmitEmpty",
			data: &struct {
				A json.Number `json:"a,omitempty"`
			}{},
		},
		{
			name: "PtrHeadNumberZeroString",
			data: &struct {
				A json.Number `json:"a,string"`
			}{},
		},

		// PtrHeadNumber
		{
			name: "PtrHeadNumber",
			data: &struct {
				A json.Number `json:"a"`
			}{A: "1"},
		},
		{
			name: "PtrHeadNumberOmitEmpty",
			data: &struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"},
		},
		{
			name: "PtrHeadNumberString",
			data: &struct {
				A json.Number `json:"a,string"`
			}{A: "1"},
		},

		// PtrHeadNumberPtr
		{
			name: "PtrHeadNumberPtr",
			data: &struct {
				A *json.Number `json:"a"`
			}{A: numberptr("1")},
		},
		{
			name: "PtrHeadNumberPtrOmitEmpty",
			data: &struct {
				A *json.Number `json:"a,omitempty"`
			}{A: numberptr("1")},
		},
		{
			name: "PtrHeadNumberPtrString",
			data: &struct {
				A *json.Number `json:"a,string"`
			}{A: numberptr("1")},
		},

		// PtrHeadNumberPtrNil
		{
			name: "PtrHeadNumberPtrNil",
			data: &struct {
				A *json.Number `json:"a"`
			}{A: nil},
		},
		{
			name: "PtrHeadNumberPtrNilOmitEmpty",
			data: &struct {
				A *json.Number `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadNumberPtrNilString",
			data: &struct {
				A *json.Number `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadNumberNil
		{
			name: "PtrHeadNumberNil",
			data: (*struct {
				A *json.Number `json:"a"`
			})(nil),
		},
		{
			name: "PtrHeadNumberNilOmitEmpty",
			data: (*struct {
				A *json.Number `json:"a,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadNumberNilString",
			data: (*struct {
				A *json.Number `json:"a,string"`
			})(nil),
		},

		// HeadNumberZeroMultiFields
		{
			name: "HeadNumberZeroMultiFields",
			data: struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
				C json.Number `json:"c"`
			}{},
		},
		{
			name: "HeadNumberZeroMultiFieldsOmitEmpty",
			data: struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
				C json.Number `json:"c,omitempty"`
			}{},
		},
		{
			name: "HeadNumberZeroMultiFields",
			data: struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
				C json.Number `json:"c,string"`
			}{},
		},

		// HeadNumberMultiFields
		{
			name: "HeadNumberMultiFields",
			data: struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
				C json.Number `json:"c"`
			}{A: "1", B: "2", C: "3"},
		},
		{
			name: "HeadNumberMultiFieldsOmitEmpty",
			data: struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
				C json.Number `json:"c,omitempty"`
			}{A: "1", B: "2", C: "3"},
		},
		{
			name: "HeadNumberMultiFieldsString",
			data: struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
				C json.Number `json:"c,string"`
			}{A: "1", B: "2", C: "3"},
		},

		// HeadNumberPtrMultiFields
		{
			name: "HeadNumberPtrMultiFields",
			data: struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
				C *json.Number `json:"c"`
			}{A: numberptr("1"), B: numberptr("2"), C: numberptr("3")},
		},
		{
			name: "HeadNumberPtrMultiFieldsOmitEmpty",
			data: struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
				C *json.Number `json:"c,omitempty"`
			}{A: numberptr("1"), B: numberptr("2"), C: numberptr("3")},
		},
		{
			name: "HeadNumberPtrMultiFieldsString",
			data: struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
				C *json.Number `json:"c,string"`
			}{A: numberptr("1"), B: numberptr("2"), C: numberptr("3")},
		},

		// HeadNumberPtrNilMultiFields
		{
			name: "HeadNumberPtrNilMultiFields",
			data: struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
				C *json.Number `json:"c"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadNumberPtrNilMultiFieldsOmitEmpty",
			data: struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
				C *json.Number `json:"c,omitempty"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadNumberPtrNilMultiFieldsString",
			data: struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
				C *json.Number `json:"c,string"`
			}{A: nil, B: nil, C: nil},
		},

		// PtrHeadNumberZeroMultiFields
		{
			name: "PtrHeadNumberZeroMultiFields",
			data: &struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
			}{},
		},
		{
			name: "PtrHeadNumberZeroMultiFieldsOmitEmpty",
			data: &struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
			}{},
		},
		{
			name: "PtrHeadNumberZeroMultiFieldsString",
			data: &struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
			}{},
		},

		// PtrHeadNumberMultiFields
		{
			name: "PtrHeadNumberMultiFields",
			data: &struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
			}{A: "1", B: "2"},
		},
		{
			name: "PtrHeadNumberMultiFieldsOmitEmpty",
			data: &struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
			}{A: "1", B: "2"},
		},
		{
			name: "PtrHeadNumberMultiFieldsString",
			data: &struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
			}{A: "1", B: "2"},
		},

		// PtrHeadNumberPtrMultiFields
		{
			name: "PtrHeadNumberPtrMultiFields",
			data: &struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
			}{A: numberptr("1"), B: numberptr("2")},
		},
		{
			name: "PtrHeadNumberPtrMultiFieldsOmitEmpty",
			data: &struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
			}{A: numberptr("1"), B: numberptr("2")},
		},
		{
			name: "PtrHeadNumberPtrMultiFieldsString",
			data: &struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
			}{A: numberptr("1"), B: numberptr("2")},
		},

		// PtrHeadNumberPtrNilMultiFields
		{
			name: "PtrHeadNumberPtrNilMultiFields",
			data: &struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilMultiFieldsOmitEmpty",
			data: &struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilMultiFieldsString",
			data: &struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
			}{A: nil, B: nil},
		},

		// PtrHeadNumberNilMultiFields
		{
			name: "PtrHeadNumberNilMultiFields",
			data: (*struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
			})(nil),
		},
		{
			name: "PtrHeadNumberNilMultiFieldsOmitEmpty",
			data: (*struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadNumberNilMultiFieldsString",
			data: (*struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
			})(nil),
		},

		// HeadNumberZeroNotRoot
		{
			name: "HeadNumberZeroNotRoot",
			data: struct {
				A struct {
					A json.Number `json:"a"`
				}
			}{},
		},
		{
			name: "HeadNumberZeroNotRootOmitEmpty",
			data: struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadNumberZeroNotRootString",
			data: struct {
				A struct {
					A json.Number `json:"a,string"`
				}
			}{},
		},

		// HeadNumberNotRoot
		{
			name: "HeadNumberNotRoot",
			data: struct {
				A struct {
					A json.Number `json:"a"`
				}
			}{A: struct {
				A json.Number `json:"a"`
			}{A: "1"}},
		},
		{
			name: "HeadNumberNotRootOmitEmpty",
			data: struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
			}{A: struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"}},
		},
		{
			name: "HeadNumberNotRootString",
			data: struct {
				A struct {
					A json.Number `json:"a,string"`
				}
			}{A: struct {
				A json.Number `json:"a,string"`
			}{A: "1"}},
		},

		// HeadNumberPtrNotRoot
		{
			name: "HeadNumberPtrNotRoot",
			data: struct {
				A struct {
					A *json.Number `json:"a"`
				}
			}{A: struct {
				A *json.Number `json:"a"`
			}{numberptr("1")}},
		},
		{
			name: "HeadNumberPtrNotRootOmitEmpty",
			data: struct {
				A struct {
					A *json.Number `json:"a,omitempty"`
				}
			}{A: struct {
				A *json.Number `json:"a,omitempty"`
			}{numberptr("1")}},
		},
		{
			name: "HeadNumberPtrNotRootString",
			data: struct {
				A struct {
					A *json.Number `json:"a,string"`
				}
			}{A: struct {
				A *json.Number `json:"a,string"`
			}{numberptr("1")}},
		},

		// HeadNumberPtrNilNotRoot
		{
			name: "HeadNumberPtrNilNotRoot",
			data: struct {
				A struct {
					A *json.Number `json:"a"`
				}
			}{},
		},
		{
			name: "HeadNumberPtrNilNotRootOmitEmpty",
			data: struct {
				A struct {
					A *json.Number `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadNumberPtrNilNotRootString",
			data: struct {
				A struct {
					A *json.Number `json:"a,string"`
				}
			}{},
		},

		// PtrHeadNumberZeroNotRoot
		{
			name: "PtrHeadNumberZeroNotRoot",
			data: struct {
				A *struct {
					A json.Number `json:"a"`
				}
			}{A: new(struct {
				A json.Number `json:"a"`
			})},
		},
		{
			name: "PtrHeadNumberZeroNotRootOmitEmpty",
			data: struct {
				A *struct {
					A json.Number `json:"a,omitempty"`
				}
			}{A: new(struct {
				A json.Number `json:"a,omitempty"`
			})},
		},
		{
			name: "PtrHeadNumberZeroNotRootString",
			data: struct {
				A *struct {
					A json.Number `json:"a,string"`
				}
			}{A: new(struct {
				A json.Number `json:"a,string"`
			})},
		},

		// PtrHeadNumberNotRoot
		{
			name: "PtrHeadNumberNotRoot",
			data: struct {
				A *struct {
					A json.Number `json:"a"`
				}
			}{A: &(struct {
				A json.Number `json:"a"`
			}{A: "1"})},
		},
		{
			name: "PtrHeadNumberNotRootOmitEmpty",
			data: struct {
				A *struct {
					A json.Number `json:"a,omitempty"`
				}
			}{A: &(struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"})},
		},
		{
			name: "PtrHeadNumberNotRootString",
			data: struct {
				A *struct {
					A json.Number `json:"a,string"`
				}
			}{A: &(struct {
				A json.Number `json:"a,string"`
			}{A: "1"})},
		},

		// PtrHeadNumberPtrNotRoot
		{
			name: "PtrHeadNumberPtrNotRoot",
			data: struct {
				A *struct {
					A *json.Number `json:"a"`
				}
			}{A: &(struct {
				A *json.Number `json:"a"`
			}{A: numberptr("1")})},
		},
		{
			name: "PtrHeadNumberPtrNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,omitempty"`
			}{A: numberptr("1")})},
		},
		{
			name: "PtrHeadNumberPtrNotRootString",
			data: struct {
				A *struct {
					A *json.Number `json:"a,string"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,string"`
			}{A: numberptr("1")})},
		},

		// PtrHeadNumberPtrNilNotRoot
		{
			name: "PtrHeadNumberPtrNilNotRoot",
			data: struct {
				A *struct {
					A *json.Number `json:"a"`
				}
			}{A: &(struct {
				A *json.Number `json:"a"`
			}{A: nil})},
		},
		{
			name: "PtrHeadNumberPtrNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,omitempty"`
			}{A: nil})},
		},
		{
			name: "PtrHeadNumberPtrNilNotRootString",
			data: struct {
				A *struct {
					A *json.Number `json:"a,string"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,string"`
			}{A: nil})},
		},

		// PtrHeadNumberNilNotRoot
		{
			name: "PtrHeadNumberNilNotRoot",
			data: struct {
				A *struct {
					A *json.Number `json:"a"`
				}
			}{A: nil},
		},
		{
			name: "PtrHeadNumberNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				} `json:",omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadNumberNilNotRootString",
			data: struct {
				A *struct {
					A *json.Number `json:"a,string"`
				} `json:",string"`
			}{A: nil},
		},

		// HeadNumberZeroMultiFieldsNotRoot
		{
			name: "HeadNumberZeroMultiFieldsNotRoot",
			data: struct {
				A struct {
					A json.Number `json:"a"`
				}
				B struct {
					B json.Number `json:"b"`
				}
			}{},
		},
		{
			name: "HeadNumberZeroMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
				B struct {
					B json.Number `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "HeadNumberZeroMultiFieldsNotRootString",
			data: struct {
				A struct {
					A json.Number `json:"a,string"`
				}
				B struct {
					B json.Number `json:"b,string"`
				}
			}{},
		},

		// HeadNumberMultiFieldsNotRoot
		{
			name: "HeadNumberMultiFieldsNotRoot",
			data: struct {
				A struct {
					A json.Number `json:"a"`
				}
				B struct {
					B json.Number `json:"b"`
				}
			}{A: struct {
				A json.Number `json:"a"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b"`
			}{B: "2"}},
		},
		{
			name: "HeadNumberMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
				B struct {
					B json.Number `json:"b,omitempty"`
				}
			}{A: struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b,omitempty"`
			}{B: "2"}},
		},
		{
			name: "HeadNumberMultiFieldsNotRootString",
			data: struct {
				A struct {
					A json.Number `json:"a,string"`
				}
				B struct {
					B json.Number `json:"b,string"`
				}
			}{A: struct {
				A json.Number `json:"a,string"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b,string"`
			}{B: "2"}},
		},

		// HeadNumberPtrMultiFieldsNotRoot
		{
			name: "HeadNumberPtrMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *json.Number `json:"a"`
				}
				B struct {
					B *json.Number `json:"b"`
				}
			}{A: struct {
				A *json.Number `json:"a"`
			}{A: numberptr("1")}, B: struct {
				B *json.Number `json:"b"`
			}{B: numberptr("2")}},
		},
		{
			name: "HeadNumberPtrMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *json.Number `json:"a,omitempty"`
				}
				B struct {
					B *json.Number `json:"b,omitempty"`
				}
			}{A: struct {
				A *json.Number `json:"a,omitempty"`
			}{A: numberptr("1")}, B: struct {
				B *json.Number `json:"b,omitempty"`
			}{B: numberptr("2")}},
		},
		{
			name: "HeadNumberPtrMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *json.Number `json:"a,string"`
				}
				B struct {
					B *json.Number `json:"b,string"`
				}
			}{A: struct {
				A *json.Number `json:"a,string"`
			}{A: numberptr("1")}, B: struct {
				B *json.Number `json:"b,string"`
			}{B: numberptr("2")}},
		},

		// HeadNumberPtrNilMultiFieldsNotRoot
		{
			name: "HeadNumberPtrNilMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *json.Number `json:"a"`
				}
				B struct {
					B *json.Number `json:"b"`
				}
			}{A: struct {
				A *json.Number `json:"a"`
			}{A: nil}, B: struct {
				B *json.Number `json:"b"`
			}{B: nil}},
		},
		{
			name: "HeadNumberPtrNilMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *json.Number `json:"a,omitempty"`
				}
				B struct {
					B *json.Number `json:"b,omitempty"`
				}
			}{A: struct {
				A *json.Number `json:"a,omitempty"`
			}{A: nil}, B: struct {
				B *json.Number `json:"b,omitempty"`
			}{B: nil}},
		},
		{
			name: "HeadNumberPtrNilMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *json.Number `json:"a,string"`
				}
				B struct {
					B *json.Number `json:"b,string"`
				}
			}{A: struct {
				A *json.Number `json:"a,string"`
			}{A: nil}, B: struct {
				B *json.Number `json:"b,string"`
			}{B: nil}},
		},

		// PtrHeadNumberZeroMultiFieldsNotRoot
		{
			name: "PtrHeadNumberZeroMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A json.Number `json:"a"`
				}
				B struct {
					B json.Number `json:"b"`
				}
			}{},
		},
		{
			name: "PtrHeadNumberZeroMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
				B struct {
					B json.Number `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "PtrHeadNumberZeroMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A json.Number `json:"a,string"`
				}
				B struct {
					B json.Number `json:"b,string"`
				}
			}{},
		},

		// PtrHeadNumberMultiFieldsNotRoot
		{
			name: "PtrHeadNumberMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A json.Number `json:"a"`
				}
				B struct {
					B json.Number `json:"b"`
				}
			}{A: struct {
				A json.Number `json:"a"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b"`
			}{B: "2"}},
		},
		{
			name: "PtrHeadNumberMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A json.Number `json:"a,omitempty"`
				}
				B struct {
					B json.Number `json:"b,omitempty"`
				}
			}{A: struct {
				A json.Number `json:"a,omitempty"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b,omitempty"`
			}{B: "2"}},
		},
		{
			name: "PtrHeadNumberMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A json.Number `json:"a,string"`
				}
				B struct {
					B json.Number `json:"b,string"`
				}
			}{A: struct {
				A json.Number `json:"a,string"`
			}{A: "1"}, B: struct {
				B json.Number `json:"b,string"`
			}{B: "2"}},
		},

		// PtrHeadNumberPtrMultiFieldsNotRoot
		{
			name: "PtrHeadNumberPtrMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *json.Number `json:"a"`
				}
				B *struct {
					B *json.Number `json:"b"`
				}
			}{A: &(struct {
				A *json.Number `json:"a"`
			}{A: numberptr("1")}), B: &(struct {
				B *json.Number `json:"b"`
			}{B: numberptr("2")})},
		},
		{
			name: "PtrHeadNumberPtrMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				}
				B *struct {
					B *json.Number `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,omitempty"`
			}{A: numberptr("1")}), B: &(struct {
				B *json.Number `json:"b,omitempty"`
			}{B: numberptr("2")})},
		},
		{
			name: "PtrHeadNumberPtrMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,string"`
				}
				B *struct {
					B *json.Number `json:"b,string"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,string"`
			}{A: numberptr("1")}), B: &(struct {
				B *json.Number `json:"b,string"`
			}{B: numberptr("2")})},
		},

		// PtrHeadNumberPtrNilMultiFieldsNotRoot
		{
			name: "PtrHeadNumberPtrNilMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *json.Number `json:"a"`
				}
				B *struct {
					B *json.Number `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				} `json:",omitempty"`
				B *struct {
					B *json.Number `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,string"`
				} `json:",string"`
				B *struct {
					B *json.Number `json:"b,string"`
				} `json:",string"`
			}{A: nil, B: nil},
		},

		// PtrHeadNumberNilMultiFieldsNotRoot
		{
			name: "PtrHeadNumberNilMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a"`
				}
				B *struct {
					B *json.Number `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberNilMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
				}
				B *struct {
					B *json.Number `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberNilMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a,string"`
				}
				B *struct {
					B *json.Number `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadNumberDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
				B *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
			}{A: &(struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
			}{A: "1", B: "2"}), B: &(struct {
				A json.Number `json:"a"`
				B json.Number `json:"b"`
			}{A: "3", B: "4"})},
		},
		{
			name: "PtrHeadNumberDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				}
				B *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				}
			}{A: &(struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
			}{A: "1", B: "2"}), B: &(struct {
				A json.Number `json:"a,omitempty"`
				B json.Number `json:"b,omitempty"`
			}{A: "3", B: "4"})},
		},
		{
			name: "PtrHeadNumberDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
				B *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
			}{A: &(struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
			}{A: "1", B: "2"}), B: &(struct {
				A json.Number `json:"a,string"`
				B json.Number `json:"b,string"`
			}{A: "3", B: "4"})},
		},

		// PtrHeadNumberNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
				B *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
				B *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadNumberNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
				B *struct {
					A json.Number `json:"a"`
					B json.Number `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				}
				B *struct {
					A json.Number `json:"a,omitempty"`
					B json.Number `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
				B *struct {
					A json.Number `json:"a,string"`
					B json.Number `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadNumberPtrDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberPtrDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
				B *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
			}{A: &(struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
			}{A: numberptr("1"), B: numberptr("2")}), B: &(struct {
				A *json.Number `json:"a"`
				B *json.Number `json:"b"`
			}{A: numberptr("3"), B: numberptr("4")})},
		},
		{
			name: "PtrHeadNumberPtrDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				}
				B *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
			}{A: numberptr("1"), B: numberptr("2")}), B: &(struct {
				A *json.Number `json:"a,omitempty"`
				B *json.Number `json:"b,omitempty"`
			}{A: numberptr("3"), B: numberptr("4")})},
		},
		{
			name: "PtrHeadNumberPtrDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
				B *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
			}{A: &(struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
			}{A: numberptr("1"), B: numberptr("2")}), B: &(struct {
				A *json.Number `json:"a,string"`
				B *json.Number `json:"b,string"`
			}{A: numberptr("3"), B: numberptr("4")})},
		},

		// PtrHeadNumberPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
				B *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
				B *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadNumberPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
				B *struct {
					A *json.Number `json:"a"`
					B *json.Number `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				}
				B *struct {
					A *json.Number `json:"a,omitempty"`
					B *json.Number `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadNumberPtrNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
				B *struct {
					A *json.Number `json:"a,string"`
					B *json.Number `json:"b,string"`
				}
			})(nil),
		},

		// AnonymousHeadNumber
		{
			name: "AnonymousHeadNumber",
			data: struct {
				structNumber
				B json.Number `json:"b"`
			}{
				structNumber: structNumber{A: "1"},
				B:            "2",
			},
		},
		{
			name: "AnonymousHeadNumberOmitEmpty",
			data: struct {
				structNumberOmitEmpty
				B json.Number `json:"b,omitempty"`
			}{
				structNumberOmitEmpty: structNumberOmitEmpty{A: "1"},
				B:                     "2",
			},
		},
		{
			name: "AnonymousHeadNumberString",
			data: struct {
				structNumberString
				B json.Number `json:"b,string"`
			}{
				structNumberString: structNumberString{A: "1"},
				B:                  "2",
			},
		},

		// PtrAnonymousHeadNumber
		{
			name: "PtrAnonymousHeadNumber",
			data: struct {
				*structNumber
				B json.Number `json:"b"`
			}{
				structNumber: &structNumber{A: "1"},
				B:            "2",
			},
		},
		{
			name: "PtrAnonymousHeadNumberOmitEmpty",
			data: struct {
				*structNumberOmitEmpty
				B json.Number `json:"b,omitempty"`
			}{
				structNumberOmitEmpty: &structNumberOmitEmpty{A: "1"},
				B:                     "2",
			},
		},
		{
			name: "PtrAnonymousHeadNumberString",
			data: struct {
				*structNumberString
				B json.Number `json:"b,string"`
			}{
				structNumberString: &structNumberString{A: "1"},
				B:                  "2",
			},
		},

		// NilPtrAnonymousHeadNumber
		{
			name: "NilPtrAnonymousHeadNumber",
			data: struct {
				*structNumber
				B json.Number `json:"b"`
			}{
				structNumber: nil,
				B:            "2",
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberOmitEmpty",
			data: struct {
				*structNumberOmitEmpty
				B json.Number `json:"b,omitempty"`
			}{
				structNumberOmitEmpty: nil,
				B:                     "2",
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberString",
			data: struct {
				*structNumberString
				B json.Number `json:"b,string"`
			}{
				structNumberString: nil,
				B:                  "2",
			},
		},

		// AnonymousHeadNumberPtr
		{
			name: "AnonymousHeadNumberPtr",
			data: struct {
				structNumberPtr
				B *json.Number `json:"b"`
			}{
				structNumberPtr: structNumberPtr{A: numberptr("1")},
				B:               numberptr("2"),
			},
		},
		{
			name: "AnonymousHeadNumberPtrOmitEmpty",
			data: struct {
				structNumberPtrOmitEmpty
				B *json.Number `json:"b,omitempty"`
			}{
				structNumberPtrOmitEmpty: structNumberPtrOmitEmpty{A: numberptr("1")},
				B:                        numberptr("2"),
			},
		},
		{
			name: "AnonymousHeadNumberPtrString",
			data: struct {
				structNumberPtrString
				B *json.Number `json:"b,string"`
			}{
				structNumberPtrString: structNumberPtrString{A: numberptr("1")},
				B:                     numberptr("2"),
			},
		},

		// AnonymousHeadNumberPtrNil
		{
			name: "AnonymousHeadNumberPtrNil",
			data: struct {
				structNumberPtr
				B *json.Number `json:"b"`
			}{
				structNumberPtr: structNumberPtr{A: nil},
				B:               numberptr("2"),
			},
		},
		{
			name: "AnonymousHeadNumberPtrNilOmitEmpty",
			data: struct {
				structNumberPtrOmitEmpty
				B *json.Number `json:"b,omitempty"`
			}{
				structNumberPtrOmitEmpty: structNumberPtrOmitEmpty{A: nil},
				B:                        numberptr("2"),
			},
		},
		{
			name: "AnonymousHeadNumberPtrNilString",
			data: struct {
				structNumberPtrString
				B *json.Number `json:"b,string"`
			}{
				structNumberPtrString: structNumberPtrString{A: nil},
				B:                     numberptr("2"),
			},
		},

		// PtrAnonymousHeadNumberPtr
		{
			name: "PtrAnonymousHeadNumberPtr",
			data: struct {
				*structNumberPtr
				B *json.Number `json:"b"`
			}{
				structNumberPtr: &structNumberPtr{A: numberptr("1")},
				B:               numberptr("2"),
			},
		},
		{
			name: "PtrAnonymousHeadNumberPtrOmitEmpty",
			data: struct {
				*structNumberPtrOmitEmpty
				B *json.Number `json:"b,omitempty"`
			}{
				structNumberPtrOmitEmpty: &structNumberPtrOmitEmpty{A: numberptr("1")},
				B:                        numberptr("2"),
			},
		},
		{
			name: "PtrAnonymousHeadNumberPtrString",
			data: struct {
				*structNumberPtrString
				B *json.Number `json:"b,string"`
			}{
				structNumberPtrString: &structNumberPtrString{A: numberptr("1")},
				B:                     numberptr("2"),
			},
		},

		// NilPtrAnonymousHeadNumberPtr
		{
			name: "NilPtrAnonymousHeadNumberPtr",
			data: struct {
				*structNumberPtr
				B *json.Number `json:"b"`
			}{
				structNumberPtr: nil,
				B:               numberptr("2"),
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberPtrOmitEmpty",
			data: struct {
				*structNumberPtrOmitEmpty
				B *json.Number `json:"b,omitempty"`
			}{
				structNumberPtrOmitEmpty: nil,
				B:                        numberptr("2"),
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberPtrString",
			data: struct {
				*structNumberPtrString
				B *json.Number `json:"b,string"`
			}{
				structNumberPtrString: nil,
				B:                     numberptr("2"),
			},
		},

		// AnonymousHeadNumberOnly
		{
			name: "AnonymousHeadNumberOnly",
			data: struct {
				structNumber
			}{
				structNumber: structNumber{A: "1"},
			},
		},
		{
			name: "AnonymousHeadNumberOnlyOmitEmpty",
			data: struct {
				structNumberOmitEmpty
			}{
				structNumberOmitEmpty: structNumberOmitEmpty{A: "1"},
			},
		},
		{
			name: "AnonymousHeadNumberOnlyString",
			data: struct {
				structNumberString
			}{
				structNumberString: structNumberString{A: "1"},
			},
		},

		// PtrAnonymousHeadNumberOnly
		{
			name: "PtrAnonymousHeadNumberOnly",
			data: struct {
				*structNumber
			}{
				structNumber: &structNumber{A: "1"},
			},
		},
		{
			name: "PtrAnonymousHeadNumberOnlyOmitEmpty",
			data: struct {
				*structNumberOmitEmpty
			}{
				structNumberOmitEmpty: &structNumberOmitEmpty{A: "1"},
			},
		},
		{
			name: "PtrAnonymousHeadNumberOnlyString",
			data: struct {
				*structNumberString
			}{
				structNumberString: &structNumberString{A: "1"},
			},
		},

		// NilPtrAnonymousHeadNumberOnly
		{
			name: "NilPtrAnonymousHeadNumberOnly",
			data: struct {
				*structNumber
			}{
				structNumber: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberOnlyOmitEmpty",
			data: struct {
				*structNumberOmitEmpty
			}{
				structNumberOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberOnlyString",
			data: struct {
				*structNumberString
			}{
				structNumberString: nil,
			},
		},

		// AnonymousHeadNumberPtrOnly
		{
			name: "AnonymousHeadNumberPtrOnly",
			data: struct {
				structNumberPtr
			}{
				structNumberPtr: structNumberPtr{A: numberptr("1")},
			},
		},
		{
			name: "AnonymousHeadNumberPtrOnlyOmitEmpty",
			data: struct {
				structNumberPtrOmitEmpty
			}{
				structNumberPtrOmitEmpty: structNumberPtrOmitEmpty{A: numberptr("1")},
			},
		},
		{
			name: "AnonymousHeadNumberPtrOnlyString",
			data: struct {
				structNumberPtrString
			}{
				structNumberPtrString: structNumberPtrString{A: numberptr("1")},
			},
		},

		// AnonymousHeadNumberPtrNilOnly
		{
			name: "AnonymousHeadNumberPtrNilOnly",
			data: struct {
				structNumberPtr
			}{
				structNumberPtr: structNumberPtr{A: nil},
			},
		},
		{
			name: "AnonymousHeadNumberPtrNilOnlyOmitEmpty",
			data: struct {
				structNumberPtrOmitEmpty
			}{
				structNumberPtrOmitEmpty: structNumberPtrOmitEmpty{A: nil},
			},
		},
		{
			name: "AnonymousHeadNumberPtrNilOnlyString",
			data: struct {
				structNumberPtrString
			}{
				structNumberPtrString: structNumberPtrString{A: nil},
			},
		},

		// PtrAnonymousHeadNumberPtrOnly
		{
			name: "PtrAnonymousHeadNumberPtrOnly",
			data: struct {
				*structNumberPtr
			}{
				structNumberPtr: &structNumberPtr{A: numberptr("1")},
			},
		},
		{
			name: "PtrAnonymousHeadNumberPtrOnlyOmitEmpty",
			data: struct {
				*structNumberPtrOmitEmpty
			}{
				structNumberPtrOmitEmpty: &structNumberPtrOmitEmpty{A: numberptr("1")},
			},
		},
		{
			name: "PtrAnonymousHeadNumberPtrOnlyString",
			data: struct {
				*structNumberPtrString
			}{
				structNumberPtrString: &structNumberPtrString{A: numberptr("1")},
			},
		},

		// NilPtrAnonymousHeadNumberPtrOnly
		{
			name: "NilPtrAnonymousHeadNumberPtrOnly",
			data: struct {
				*structNumberPtr
			}{
				structNumberPtr: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberPtrOnlyOmitEmpty",
			data: struct {
				*structNumberPtrOmitEmpty
			}{
				structNumberPtrOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadNumberPtrOnlyString",
			data: struct {
				*structNumberPtrString
			}{
				structNumberPtrString: nil,
			},
		},
	}
	for _, test := range tests {
		for _, indent := range []bool{true, false} {
			for _, htmlEscape := range []bool{true, false} {
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				enc.SetEscapeHTML(htmlEscape)
				if indent {
					enc.SetIndent("", "  ")
				}
				if err := enc.Encode(test.data); err != nil {
					t.Fatalf("%s(htmlEscape:%v,indent:%v): %+v: %s", test.name, htmlEscape, indent, test.data, err)
				}
				stdresult := encodeByEncodingJSON(test.data, indent, htmlEscape)
				if buf.String() != stdresult {
					t.Errorf("%s(htmlEscape:%v,indent:%v): doesn't compatible with encoding/json. expected %q but got %q", test.name, htmlEscape, indent, stdresult, buf.String())
				}
			}
		}
	}
}
