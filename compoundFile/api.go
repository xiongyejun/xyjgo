package compoundFile

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

func (me *CompoundFile) ModifyStream(streamPath string, newByte []byte) (err error) {
	var stream *cfStream
	if stream, err = me.getStreamItem(streamPath); err != nil {
		return
	}
	oldByte := stream.stream.Bytes()

	// b2保持大小不变，方便复制到filebyte
	b2 := make([]byte, len(oldByte))
	copy(b2, newByte)

	arr := strings.Split(streamPath, `\`)
	streamName := arr[len(arr)-1]

	var streamIndex int32
	var ok bool
	if streamIndex, ok = me.cfs.dic[streamName]; !ok {
		return errors.New("不存在的stream：" + streamName)
	}

	// 修改文件byte
	for i, v := range me.cfs.arrStream[streamIndex].address {
		bStart := int32(i) * me.cfs.arrStream[streamIndex].step
		bEnd := bStart + me.cfs.arrStream[streamIndex].step
		copy(me.SrcByte[v:], b2[bStart:bEnd])
	}

	// 修改dir中的Stream_size
	// b中实际仅有me.cfs.arrDir[streamIndex].Stream_size的大小，但是为了上面循环方便按照step复制，在这里来扣除多余的
	iSub := int32(len(oldByte)) - me.cfs.arrDir[streamIndex].Stream_size
	var iLen int32 = int32(len(newByte)) - iSub
	// int32转byte
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, &iLen)
	// 内存中的数据也修改下
	me.cfs.arrDir[streamIndex].Stream_size = iLen
	// fileByte的下标
	indexStreamSize := me.cfs.arrDirAddr[streamIndex] + DIR_SIZE - 8 // -8是因为在倒数第2个，减2个int32
	copy(me.SrcByte[indexStreamSize:], buf.Bytes())

	return
}
