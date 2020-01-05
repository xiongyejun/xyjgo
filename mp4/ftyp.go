package mp4

import (
	"fmt"
	"io"
	"io/ioutil"
)

// File Type Box (ftyp - mandatory)
//
// Status: decoded
type FtypBox struct {
	MajorBrand       []byte
	MinorVersion     []byte
	CompatibleBrands []byte
}

func DecodeFtyp(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	b := &FtypBox{
		MajorBrand:       data[0:4],
		MinorVersion:     data[4:8],
		CompatibleBrands: data[8:],
	}
	//	if len(data) > 8 {
	//		for i := 8; i < len(data); i += 4 {
	//			b.CompatibleBrands = append(b.CompatibleBrands, string(data[i:i+4]))
	//		}
	//	}
	return b, nil
}

func (b *FtypBox) Type() string {
	return "ftyp"
}

func (b *FtypBox) Size() int {
	return BoxHeaderSize + 8 + len(b.CompatibleBrands)
}

func (b *FtypBox) Dump() {
	fmt.Printf("File Type: %s\n", b.MajorBrand)
}

func (b *FtypBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	buf := makebuf(b)
	//	strtobuf(buf, b.MajorBrand, 4)
	copy(buf[0:], b.MajorBrand)
	copy(buf[4:], b.MinorVersion)
	copy(buf[8:], b.CompatibleBrands)
	//	for i, c := range b.CompatibleBrands {
	//		strtobuf(buf[8+i*4:], c, 4)
	//	}
	_, err = w.Write(buf)
	return err
}
