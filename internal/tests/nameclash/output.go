// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package nameclash

import (
	"encoding/json"

	errors0 "github.com/fjl/gencodec/internal/clasherrors"
	json0 "github.com/fjl/gencodec/internal/clashjson"
)

func (y Y) MarshalJSON() ([]byte, error) {
	type Y struct {
		Foo    json0.Foo
		Foo2   json0.Foo
		Bar    errors0.Foo
		Gazonk YJSON
		Over   enc
	}
	var enc0 Y
	enc0.Foo = y.Foo
	enc0.Foo2 = y.Foo2
	enc0.Bar = y.Bar
	enc0.Gazonk = y.Gazonk
	enc0.Over = enc(y.Over)
	return json.Marshal(&enc0)
}

func (y *Y) UnmarshalJSON(input []byte) error {
	type Y struct {
		Foo    *json0.Foo
		Foo2   *json0.Foo
		Bar    *errors0.Foo
		Gazonk *YJSON
		Over   *enc
	}
	var dec Y
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Foo != nil {
		y.Foo = *dec.Foo
	}
	if dec.Foo2 != nil {
		y.Foo2 = *dec.Foo2
	}
	if dec.Bar != nil {
		y.Bar = *dec.Bar
	}
	if dec.Gazonk != nil {
		y.Gazonk = *dec.Gazonk
	}
	if dec.Over != nil {
		y.Over = int(*dec.Over)
	}
	return nil
}

func (y Y) MarshalYAML() (interface{}, error) {
	type Y struct {
		Foo    json0.Foo
		Foo2   json0.Foo
		Bar    errors0.Foo
		Gazonk YJSON
		Over   enc
	}
	var enc0 Y
	enc0.Foo = y.Foo
	enc0.Foo2 = y.Foo2
	enc0.Bar = y.Bar
	enc0.Gazonk = y.Gazonk
	enc0.Over = enc(y.Over)
	return &enc0, nil
}

func (y *Y) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Y struct {
		Foo    *json0.Foo
		Foo2   *json0.Foo
		Bar    *errors0.Foo
		Gazonk *YJSON
		Over   *enc
	}
	var dec Y
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Foo != nil {
		y.Foo = *dec.Foo
	}
	if dec.Foo2 != nil {
		y.Foo2 = *dec.Foo2
	}
	if dec.Bar != nil {
		y.Bar = *dec.Bar
	}
	if dec.Gazonk != nil {
		y.Gazonk = *dec.Gazonk
	}
	if dec.Over != nil {
		y.Over = int(*dec.Over)
	}
	return nil
}

func (y Y) MarshalTOML() (interface{}, error) {
	type Y struct {
		Foo    json0.Foo
		Foo2   json0.Foo
		Bar    errors0.Foo
		Gazonk YJSON
		Over   enc
	}
	var enc0 Y
	enc0.Foo = y.Foo
	enc0.Foo2 = y.Foo2
	enc0.Bar = y.Bar
	enc0.Gazonk = y.Gazonk
	enc0.Over = enc(y.Over)
	return &enc0, nil
}

func (y *Y) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Y struct {
		Foo    *json0.Foo
		Foo2   *json0.Foo
		Bar    *errors0.Foo
		Gazonk *YJSON
		Over   *enc
	}
	var dec Y
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Foo != nil {
		y.Foo = *dec.Foo
	}
	if dec.Foo2 != nil {
		y.Foo2 = *dec.Foo2
	}
	if dec.Bar != nil {
		y.Bar = *dec.Bar
	}
	if dec.Gazonk != nil {
		y.Gazonk = *dec.Gazonk
	}
	if dec.Over != nil {
		y.Over = int(*dec.Over)
	}
	return nil
}
