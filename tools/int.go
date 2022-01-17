package tools

import "encoding/binary"

func Uint64ToByteUseBig(i uint64) []byte {
	s := make([]byte, 8)
	binary.BigEndian.PutUint64(s, i)
	return s
}
func Uint64ToByteUseLittle(i uint64) []byte {
	s := make([]byte, 8)
	binary.LittleEndian.PutUint64(s, i)
	return s
}

func Uint32ToByteUseLittle(i uint32) []byte {
	s := make([]byte, 4)
	binary.LittleEndian.PutUint32(s, i)
	return s
}

func Uint32ToByteUseBig(i uint32) []byte {
	s := make([]byte, 4)
	binary.BigEndian.PutUint32(s, i)
	return s
}
