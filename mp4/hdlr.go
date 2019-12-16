package mp4

import (
	"encoding/binary"
	"io"
	"io/ioutil"
)

// Handler Reference Box (hdlr - mandatory)
//
// Contained in: Media Box (mdia) or Meta Box (meta)
//
// Status: decoded
//
// This box describes the type of data contained in the trak.
//
// HandlerType can be : "vide" (video track), "soun" (audio track), "hint" (hint track), "meta" (timed Metadata track), "auxv" (auxiliary video track).
// https://blog.csdn.net/yue_huang/article/details/72812109
/*
Handler reference atom 定义了描述此媒体数据的media handler component，类型是'hdlr'。在过去，handler reference atom也可以用来数据引用，但是现在，已经不允许这样使用了。一个media atom内的handler atom解释了媒体流的播放过程。例如，一个视频handler处理一个video track.
*/
type HdlrBox struct {
	Version                    byte
	Flags                      [3]byte
	ComponetType               uint32
	ComponetSubType            uint32
	ComponetSubTypeManufactuer uint32
	ComponetFlags              uint32
	ComponetFlagsMask          uint32
	ComponetName               []byte
}

func DecodeHdlr(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &HdlrBox{
		Version:                    data[0],
		Flags:                      [3]byte{data[1], data[2], data[3]},
		ComponetType:               binary.BigEndian.Uint32(data[4:8]),
		ComponetSubType:            binary.BigEndian.Uint32(data[8:12]),
		ComponetSubTypeManufactuer: binary.BigEndian.Uint32(data[12:16]),
		ComponetFlags:              binary.BigEndian.Uint32(data[16:20]),
		ComponetFlagsMask:          binary.BigEndian.Uint32(data[20:24]),
		ComponetName:               data[24:],
	}, nil
}

func (b *HdlrBox) Type() string {
	return "hdlr"
}

func (b *HdlrBox) Size() int {
	return BoxHeaderSize + 24 + len(b.ComponetName)
}

func (b *HdlrBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	buf := makebuf(b)
	buf[0] = b.Version
	buf[1], buf[2], buf[3] = b.Flags[0], b.Flags[1], b.Flags[2]
	binary.BigEndian.PutUint32(buf[4:], b.ComponetType)
	binary.BigEndian.PutUint32(buf[8:], b.ComponetSubType)
	binary.BigEndian.PutUint32(buf[12:], b.ComponetSubTypeManufactuer)
	binary.BigEndian.PutUint32(buf[16:], b.ComponetFlags)
	binary.BigEndian.PutUint32(buf[20:], b.ComponetFlagsMask)
	copy(buf[24:], b.ComponetName)

	_, err = w.Write(buf)
	return err
}
