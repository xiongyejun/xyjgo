package mp4

import (
	"io"
	"io/ioutil"
)

type UnknownBox struct {
	notDecoded []byte
}

func DecodeUnknown(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &UnknownBox{data}, nil
}

func (b *UnknownBox) Type() string {
	return "Unknown"
}

func (b *UnknownBox) Size() int {
	return BoxHeaderSize + len(b.notDecoded)
}

func (b *UnknownBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	_, err = w.Write(b.notDecoded)
	return err
}
