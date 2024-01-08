package utils

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"time"
)

// 将byte转换成字符串16进制的字符串
func EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func uint32ToBin(num uint32) []byte {
	uintByte := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(uintByte, num)
	return uintByte
}

func uint32ToHexString(num uint32) string {
	return EncodeToString(uint32ToBin(num))
}

func timeStamp() uint32 {
	return uint32(time.Now().Unix())
}

func myrand() uint32 {
	return uint32((rand.Int31()))
}
