// 原代码：https://github.com/ThankYouMotion/mp4

package mp4

import (
	"io"
	"os"
)

// A MPEG-4 media
//
// A MPEG-4 media contains three main boxes :
//
//   ftyp : the file type box
//   moov : the movie box (meta-data)
//   mdat : the media data (chunks and samples)
//
// Other boxes can also be present (pdin, moof, mfra, free, ...), but are not decoded.
type MP4 struct {
	Ftyp  *FtypBox
	Moov  *MoovBox
	Mdat  *MdatBox
	boxes []Box
}

// Decode decodes a media from a Reader
func Decode(r *os.File) (m *MP4, err error) {
	v := &MP4{
		boxes: []Box{},
	}

	var fi os.FileInfo
	if fi, err = r.Stat(); err != nil {
		return nil, err
	}

	flen := fi.Size()
	var tmp int64 = 0

	for tmp < flen {
		h, err := DecodeHeader(r)
		if err != nil {
			return nil, err
		}
		tmp += int64(h.Size)

		box, err := DecodeBox(h, r)
		if err != nil {
			return nil, err
		}
		v.boxes = append(v.boxes, box)
		switch h.Type {
		case "ftyp":
			v.Ftyp = box.(*FtypBox)
		case "moov":
			v.Moov = box.(*MoovBox)
			return v, nil
		case "mdat":
			v.Mdat = box.(*MdatBox)
			v.Mdat.ContentSize = h.Size - BoxHeaderSize

			r.Seek(tmp, 0)
		}
	}
	return v, nil
}

// Dump displays some information about a media
func (m *MP4) Dump() {
	m.Ftyp.Dump()
	m.Moov.Dump()
}

// Boxes lists the top-level boxes from a media
func (m *MP4) Boxes() []Box {
	return m.boxes
}

// Encode encodes a media to a Writer
func (m *MP4) Encode(w io.Writer) error {
	err := m.Ftyp.Encode(w)
	if err != nil {
		return err
	}
	err = m.Moov.Encode(w)
	if err != nil {
		return err
	}
	for _, b := range m.boxes {
		if b.Type() != "ftyp" && b.Type() != "moov" {
			err = b.Encode(w)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
