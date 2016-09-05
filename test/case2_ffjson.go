// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: case2.go
// DO NOT EDIT!

package test

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *RecursiveStruct) MarshalFFJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *RecursiveStruct) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"IntField":`)
	fflib.FormatBits2(buf, uint64(mj.IntField), 10, mj.IntField < 0)
	buf.WriteString(`,"StrField":`)
	fflib.WriteJsonString(buf, string(mj.StrField))
	if mj.RecursiveField != nil {
		buf.WriteString(`,"RecursiveField":`)

		{

			err = mj.RecursiveField.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"RecursiveField":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_RecursiveStructbase = iota
	ffj_t_RecursiveStructno_such_key

	ffj_t_RecursiveStruct_IntField

	ffj_t_RecursiveStruct_StrField

	ffj_t_RecursiveStruct_RecursiveField
)

var ffj_key_RecursiveStruct_IntField = []byte("IntField")

var ffj_key_RecursiveStruct_StrField = []byte("StrField")

var ffj_key_RecursiveStruct_RecursiveField = []byte("RecursiveField")

func (uj *RecursiveStruct) UnmarshalFFJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *RecursiveStruct) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_RecursiveStructbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_RecursiveStructno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'I':

					if bytes.Equal(ffj_key_RecursiveStruct_IntField, kn) {
						currentKey = ffj_t_RecursiveStruct_IntField
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'R':

					if bytes.Equal(ffj_key_RecursiveStruct_RecursiveField, kn) {
						currentKey = ffj_t_RecursiveStruct_RecursiveField
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'S':

					if bytes.Equal(ffj_key_RecursiveStruct_StrField, kn) {
						currentKey = ffj_t_RecursiveStruct_StrField
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_RecursiveStruct_RecursiveField, kn) {
					currentKey = ffj_t_RecursiveStruct_RecursiveField
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_RecursiveStruct_StrField, kn) {
					currentKey = ffj_t_RecursiveStruct_StrField
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_RecursiveStruct_IntField, kn) {
					currentKey = ffj_t_RecursiveStruct_IntField
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_RecursiveStructno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_RecursiveStruct_IntField:
					goto handle_IntField

				case ffj_t_RecursiveStruct_StrField:
					goto handle_StrField

				case ffj_t_RecursiveStruct_RecursiveField:
					goto handle_RecursiveField

				case ffj_t_RecursiveStructno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_IntField:

	/* handler: uj.IntField type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.IntField = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_StrField:

	/* handler: uj.StrField type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.StrField = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RecursiveField:

	/* handler: uj.RecursiveField type=test.RecursiveStruct kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.RecursiveField = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.RecursiveField == nil {
			uj.RecursiveField = new(RecursiveStruct)
		}

		err = uj.RecursiveField.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}