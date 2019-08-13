// ICMP Internet Control Message Protocol
// Internet 控制报文协议
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
)

//RFC792定义的echo数据包结构：

//    0                   1                   2                   3
//    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |     Type      |     Code      |          Checksum             |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |           Identifier          |        Sequence Number        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |     Data ...
//   +-+-+-+-+-
//---------------------
//作者：爱神CODE
//来源：CSDN
//原文：https://blog.csdn.net/gophers/article/details/21481447

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func main() {
	var icmp *ICMP = new(ICMP)
	laddr := net.IPAddr{IP: net.ParseIP("192.168.1.5")}
	raddr := net.IPAddr{IP: net.ParseIP("192.168.1.5")}

	conn, err := net.DialIP("ip4:icmp", &laddr, &raddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	icmp.Type = 8

	buf := bytes.Buffer{}
	if err = binary.Write(&buf, binary.BigEndian, icmp); err != nil {
		fmt.Println(err)
		return
	}
	icmp.GetCheckSum(buf.Bytes())
	fmt.Println("icmp", icmp)

	buf.Reset()
	if err = binary.Write(&buf, binary.BigEndian, icmp); err != nil {
		fmt.Println(err)
		return
	}

	if _, err = conn.Write(buf.Bytes()); err != nil {
		fmt.Println(err)
		return
	}

	var b []byte
	if b, err = ioutil.ReadAll(conn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(123)

	fmt.Println(conn.RemoteAddr(), string(b))
	fmt.Println("ok")
}

func (me *ICMP) GetCheckSum(data []byte) {
	length := len(data)

	var sum uint32
	var i int = 0
	for ; length > 1; i += 2 {
		sum += uint32(data[i])<<8 + uint32(data[i+1])
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[i])
	}
	sum += (sum >> 16)
	me.Checksum = uint16(^sum)
}
