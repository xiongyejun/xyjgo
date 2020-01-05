package mp4

import "io"

// Sample Table Box (stbl - mandatory)
//
// Contained in : Media Information Box (minf)
//
// Status: partially decoded (anything other than stsd, stts, stsc, stss, stsz, stco, ctts is ignored)
//
// The table contains all information relevant to data samples (times, chunks, sizes, ...)
/*
sample table
存储媒体数据的单位是samples。一个sample是一系列按时间顺序排列的数据的一个element。Samples存储在media中的chunk内，可以有不同的durations。Chunk存储一个或者多个samples，是数据存取的基本单位，可以有不同的长度，一个chunk内的每个sample也可以有不同的长度。例如如下图，chunk 2和3不同的长度，chunk 2内的sample5和6的长度一样，但是sample 4和5，6的长度不同
*/
type StblBox struct {
	Stsd *StsdBox // sample description
	Stts *SttsBox // time to sample
	Stss *StssBox // sync sample box
	Stsc *StscBox // sample to chunk box
	Stsz *StszBox // sample size
	Stco *StcoBox // chunk offset box
	Ctts *CttsBox // composition time to sample box
}

func DecodeStbl(r io.Reader) (Box, error) {
	l, err := DecodeContainer(r)
	if err != nil {
		return nil, err
	}
	s := &StblBox{}
	for _, b := range l {
		switch b.Type() {
		case "stsd":
			s.Stsd = b.(*StsdBox)
		case "stts":
			s.Stts = b.(*SttsBox)
		case "stsc":
			s.Stsc = b.(*StscBox)
		case "stss":
			s.Stss = b.(*StssBox)
		case "stsz":
			s.Stsz = b.(*StszBox)
		case "stco":
			s.Stco = b.(*StcoBox)
		case "ctts":
			s.Ctts = b.(*CttsBox)
		}
	}
	return s, nil
}

func (b *StblBox) Type() string {
	return "stbl"
}

func (b *StblBox) Size() int {
	sz := b.Stsd.Size()
	if b.Stts != nil {
		sz += b.Stts.Size()
	}
	if b.Stss != nil {
		sz += b.Stss.Size()
	}
	if b.Stsc != nil {
		sz += b.Stsc.Size()
	}
	if b.Stsz != nil {
		sz += b.Stsz.Size()
	}
	if b.Stco != nil {
		sz += b.Stco.Size()
	}
	if b.Ctts != nil {
		sz += b.Ctts.Size()
	}
	return sz + BoxHeaderSize
}

func (b *StblBox) Dump() {
	if b.Stsc != nil {
		b.Stsc.Dump()
	}
	if b.Stts != nil {
		b.Stts.Dump()
	}
	if b.Stsz != nil {
		b.Stsz.Dump()
	}
	if b.Stss != nil {
		b.Stss.Dump()
	}
	if b.Stco != nil {
		b.Stco.Dump()
	}
}

func (b *StblBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	err = b.Stsd.Encode(w)
	if err != nil {
		return err
	}
	err = b.Stts.Encode(w)
	if err != nil {
		return err
	}
	if b.Stss != nil {
		err = b.Stss.Encode(w)
		if err != nil {
			return err
		}
	}
	err = b.Stsc.Encode(w)
	if err != nil {
		return err
	}
	err = b.Stsz.Encode(w)
	if err != nil {
		return err
	}
	err = b.Stco.Encode(w)
	if err != nil {
		return err
	}
	if b.Ctts != nil {
		return b.Ctts.Encode(w)
	}
	return nil
}
