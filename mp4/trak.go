package mp4

import "io"

// Track Box (tkhd - mandatory)
//
// Contained in : Movie Box (moov)
//
// A media file can contain one or more tracks.
/*
例如，《让子弹飞》的正版DVD，
1）有一条视频轨用于电影画面。
2）至少有两条音频轨分别提供了普通话与四川话版，实际上为了营造更加逼真的现场效果，为了配合多声道家庭影院该影片还独有一条音效轨。
3）多条字幕轨，简体中文，繁体中文，英文……。
从中我们可以理解为什么trak box可以有多个：每个track都是独立的，具有自我特征与属性的，因此需要各自描述互不干涉。

有很多种类的track，其中有三个最重要，video track包含了视频sample；audio track包含了audiosample；hint track稍有不同，它描述了一个流媒体服务器如何把文件中的媒体数据组成符合流媒体协议的数据包。 如果文件只是本地播放，可以忽略hint track，他们只与流媒体有关系。

Track atoms 的atom类型是'trak'. Track atom要求必须有一个track header atom ('tkhd')和一个media atom ('mdia')。其他的track clipping atom ('clip')，track matte atom ('matt')，edit atom ('edts')，track reference atom ('tref')，track load settings atom ('load')，a track input map atom ('imap')以及user data atom ('udta')都是可选的。
*/
type TrakBox struct {
	Tkhd *TkhdBox
	Mdia *MdiaBox
	Edts *EdtsBox
}

func DecodeTrak(r io.Reader) (Box, error) {
	l, err := DecodeContainer(r)
	if err != nil {
		return nil, err
	}
	t := &TrakBox{}
	for _, b := range l {
		switch b.Type() {
		case "tkhd":
			t.Tkhd = b.(*TkhdBox)
		case "mdia":
			t.Mdia = b.(*MdiaBox)
		case "edts":
			t.Edts = b.(*EdtsBox)
		default:
			return nil, ErrBadFormat
		}
	}
	return t, nil
}

func (b *TrakBox) Type() string {
	return "trak"
}

func (b *TrakBox) Size() int {
	sz := b.Tkhd.Size()
	sz += b.Mdia.Size()
	if b.Edts != nil {
		sz += b.Edts.Size()
	}
	return sz + BoxHeaderSize
}

func (b *TrakBox) Dump() {
	b.Tkhd.Dump()
	if b.Edts != nil {
		b.Edts.Dump()
	}
	b.Mdia.Dump()
}

func (b *TrakBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	err = b.Tkhd.Encode(w)
	if err != nil {
		return err
	}
	if b.Edts != nil {
		err = b.Edts.Encode(w)
		if err != nil {
			return err
		}
	}
	return b.Mdia.Encode(w)
}
