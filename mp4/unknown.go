package mp4

import (
	"io"
	"io/ioutil"
)

type TrefBox struct {
	notDecoded []byte
}

func DecodeTref(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &TrefBox{data}, nil
}

func (b *TrefBox) Type() string {
	return "Tref"
}

func (b *TrefBox) Size() int {
	return BoxHeaderSize + len(b.notDecoded)
}

func (b *TrefBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	_, err = w.Write(b.notDecoded)
	return err
}
