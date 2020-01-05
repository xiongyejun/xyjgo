// 原代码：https://github.com/ThankYouMotion/mp4

package mp4

import (
	"fmt"
	"errors"
	"io"
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
func Decode(r io.Reader) (m *MP4, err error) {
	v := &MP4{
		boxes: []Box{},
	}

	var mdatSeek int64 = 0
	var flag byte = 0
	for flag != 0b00000111 {
		h, err := DecodeHeader(r)
		fmt.Printf("type=%s, size=%d\n", h.Type,h.Size)
		if err != nil {
			return nil, err
		}

		box, err := DecodeBox(h, r)
		if err != nil {
			return nil, err
		}

		v.boxes = append(v.boxes, box)
		switch h.Type {
		case "ftyp":
			v.Ftyp = box.(*FtypBox)
			flag = flag | 0b00000001
		case "moov":
			v.Moov = box.(*MoovBox)
			flag = flag | 0b00000010
		case "mdat":
			v.Mdat = box.(*MdatBox)
			v.Mdat.ContentSize = h.Size - BoxHeaderSize
			flag = flag | 0b00000100
			
			if rs, ok := r.(io.Seeker); ok {				
				rs.Seek(int64(v.Mdat.ContentSize), 1)
			} else {
				return nil, errors.New("Seeker err")
			}
		}
			// 记录mdat之前的box
			if flag & 0b00000100 != 0b00000100 {
				mdatSeek += int64(box.Size())
			}
	}
	
	mdatSeek += int64(BoxHeaderSize)
	fmt.Printf("mdatSeek = %x\n",mdatSeek)
	if rs, ok := v.Mdat.Reader().(io.Seeker); ok {				
		if _,err := rs.Seek(mdatSeek, io.SeekStart); err != nil {
			return nil,err
		}
	}
	
//	有的为什么还是没有跳过Mdat的Header
//	var b []byte = make([]byte, 8) 
//	v.Mdat.Reader().Read(b)
//	fmt.Printf("% x\n", b)
//	fmt.Printf("% s\n", b)
			
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
