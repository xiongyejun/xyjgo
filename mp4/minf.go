package mp4

import "io"

// Media Information Box (minf - mandatory)
//
// Contained in : Media Box (mdia)
//
// Status: partially decoded (hmhd - hint tracks - and nmhd - null media - are ignored)
/*
Media information atoms的类型是'minf'，存储了解释该track的媒体数据的handler-specific的信息。media handler用这些信息将媒体时间映射到媒体数据，并进行处理。它是一个容器atom，包含其他的子atom。
这些信息是与媒体定义的数据类型特别对应的，而且media information atoms 的格式和内容也是与解释此媒体数据流的media handler 密切相关的。其他的media handler不知道如何解释这些信息。
*/
type MinfBox struct {
	Vmhd *VmhdBox
	Smhd *SmhdBox
	Stbl *StblBox
	Dinf *DinfBox
	Hdlr *HdlrBox
}

func DecodeMinf(r io.Reader) (Box, error) {
	l, err := DecodeContainer(r)
	if err != nil {
		return nil, err
	}
	m := &MinfBox{}
	for _, b := range l {
		switch b.Type() {
		case "vmhd":
			m.Vmhd = b.(*VmhdBox)
		case "smhd":
			m.Smhd = b.(*SmhdBox)
		case "stbl":
			m.Stbl = b.(*StblBox)
		case "dinf":
			m.Dinf = b.(*DinfBox)
		case "hdlr":
			m.Hdlr = b.(*HdlrBox)
		}
	}
	return m, nil
}

func (b *MinfBox) Type() string {
	return "minf"
}

func (b *MinfBox) Size() int {
	sz := 0
	if b.Vmhd != nil {
		sz += b.Vmhd.Size()
	}
	if b.Smhd != nil {
		sz += b.Smhd.Size()
	}
	sz += b.Stbl.Size()
	if b.Dinf != nil {
		sz += b.Dinf.Size()
	}
	if b.Hdlr != nil {
		sz += b.Hdlr.Size()
	}
	return sz + BoxHeaderSize
}

func (b *MinfBox) Dump() {
	b.Stbl.Dump()
}

func (b *MinfBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	if b.Vmhd != nil {
		err = b.Vmhd.Encode(w)
		if err != nil {
			return err
		}
	}
	if b.Smhd != nil {
		err = b.Smhd.Encode(w)
		if err != nil {
			return err
		}
	}
	err = b.Dinf.Encode(w)
	if err != nil {
		return err
	}
	err = b.Stbl.Encode(w)
	if err != nil {
		return err
	}
	if b.Hdlr != nil {
		return b.Hdlr.Encode(w)
	}
	return nil
}
