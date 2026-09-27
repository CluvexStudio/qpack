package qpack

import "io"

func (d *Decoder) DecodeFull(p []byte) ([]HeaderField, error) {
	var fields []HeaderField
	next := d.Decode(p)
	for {
		hf, err := next()
		if err == io.EOF {
			return fields, nil
		}
		if err != nil {
			return nil, err
		}
		fields = append(fields, hf)
	}
}
