package crypto

import (
	"io"

	"github.com/foxxorcat/DriverCore/tools"
)

type cryptoReader struct {
	crypto func(dst []byte, src []byte)
	reader io.Reader

	offset int

	isInit bool
	data   []byte // 加密后数据
	buf    []byte
}

// Stream
type cryptoStream struct {
	cryptoReader
	bufSize int // 块大小
}

func (a *cryptoStream) Read(p []byte) (n int, err error) {
	if a.offset >= len(a.data) {
		if !a.isInit {
			a.data, a.buf, a.isInit = make([]byte, a.bufSize), make([]byte, a.bufSize), true
		}

		n, err = a.reader.Read(a.buf)
		a.crypto(a.data[:n], a.buf[:n])
		a.offset, a.data = 0, a.data[:n]
	}
	n = copy(p, a.data[a.offset:])
	a.offset += n
	return
}

// Block
type cryptoBlockEncrypt struct {
	cryptoReader
	blockSize int
	readFull  bool
}

func (a *cryptoBlockEncrypt) Read(p []byte) (n int, err error) {
	if a.offset >= len(a.data) {
		if !a.isInit {
			a.data, a.buf, a.isInit = make([]byte, a.blockSize), make([]byte, a.blockSize), true
		}
		if a.readFull {
			return 0, io.EOF
		}

		n, err = io.ReadAtLeast(a.reader, a.buf, a.blockSize)
		if (err == io.EOF || err == io.ErrUnexpectedEOF) && !a.readFull {
			a.buf, a.readFull, n = tools.PKCS7Padding(a.buf[:n], a.blockSize), true, a.blockSize
		}

		a.crypto(a.data[:n], a.buf[:n])
		a.offset, a.data = 0, a.data[:n]
	}
	n = copy(p, a.data[a.offset:])
	a.offset += n
	return
}

type cryptoBlockDecrypt struct {
	cryptoReader
	blockSize int
}

func (a *cryptoBlockDecrypt) Read(p []byte) (n int, err error) {
	if a.offset >= len(a.data) {
		if !a.isInit {
			a.data, a.buf, a.isInit = make([]byte, a.blockSize), make([]byte, a.blockSize), true
			if n, err = io.ReadAtLeast(a.reader, a.buf[:1], 1); err != nil {
				return
			}
		}

		if _, err = io.ReadAtLeast(a.reader, a.buf[1:], a.blockSize-1); err != nil {
			return 0, err
		}

		a.crypto(a.data, a.buf)
		a.offset = 0

		// 再次读取判断是否到尾部
		if n, err = io.ReadAtLeast(a.reader, a.buf[:1], 1); err != nil {
			if err != io.EOF {
				return
			}
			a.data = tools.PKCS7UnPadding(a.data)
		}
	}
	n = copy(p, a.data[a.offset:])
	a.offset += n
	return
}
