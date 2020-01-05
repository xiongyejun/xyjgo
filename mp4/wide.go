package mp4

import (
	"io"
	"io/ioutil"
)

type WideBox struct {
	notDecoded []byte
}

func DecodeWide(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &WideBox{data}, nil
}

func (b *WideBox) Type() string {
	return "Wide"
}

func (b *WideBox) Size() int {
	return BoxHeaderSize + len(b.notDecoded)
}

func (b *WideBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	_, err = w.Write(b.notDecoded)
	return err
}
