package utils

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func ReadChar(pkt []byte, idx *int) string {
	ans := string(pkt[*idx])
	*idx++
	return ans
}

func ReadByte(pkt []byte, idx *int) uint8 {
	ans := pkt[*idx]
	*idx++
	return ans
}

func WriteByte(pkt []byte, num uint8) []byte {
	return append(pkt, num)
}

func ReadBytes(pkt []byte, idx *int, n int) []byte {
	ans := pkt[*idx : *idx+n]
	*idx += n
	return ans
}

func WriteBytes(pkt, arr []byte) []byte {
	return append(pkt, arr...)
}

func ReadString(pkt []byte, idx *int, n int) string {
	return string(ReadBytes(pkt, idx, n))
}

func WriteString(pkt []byte, str string) []byte {
	arr := []byte(str)
	return append(pkt, arr...)
}

func ReadHexEncoding(pkt []byte, idx *int, n int) string {
	ans := hex.EncodeToString(pkt[*idx : *idx+n])
	*idx += n
	return strings.ToUpper(ans)
}

func WriteHexEncoding(pkt []byte, src string) ([]byte, error) {
	dst, err := hex.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return append(pkt, dst...), nil
}

/*
 *
 *
 * BigEndian
 *
 *
 */

func ReadWordBigEndian(pkt []byte, idx *int) uint16 {
	ans := binary.BigEndian.Uint16(pkt[*idx : *idx+2])
	*idx += 2
	return ans
}

func WriteWordBigEndian(pkt []byte, num uint16) []byte {
	numPkt := make([]byte, 2)
	binary.BigEndian.PutUint16(numPkt, num)
	return append(pkt, numPkt...)
}

// -- -- -- -- -- -- --

func ReadThreeBytesBigEndian(pkt []byte, idx *int) uint32 {
	ans := uint32(pkt[*idx])<<16 | uint32(pkt[*idx+1])<<8 | uint32(pkt[*idx+2])
	*idx += 3
	return ans
}

func WriteThreeBytesBigEndian(pkt []byte, num uint32) []byte {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, num)
	threeBytes := tmp[1:]
	return append(pkt, threeBytes...)
}

// -- -- -- -- -- -- --

func ReadDoubleWordBigEndian(pkt []byte, idx *int) uint32 {
	ans := binary.BigEndian.Uint32(pkt[*idx : *idx+4])
	*idx += 4
	return ans
}

func WriteDoubleWordBigEndian(pkt []byte, num uint32) []byte {
	numPkt := make([]byte, 4)
	binary.BigEndian.PutUint32(numPkt, num)
	return append(pkt, numPkt...)
}

/*
 *
 *
 * LittleEndian
 *
 *
 */

func ReadWordLittleEndian(pkt []byte, idx *int) uint16 {
	ans := binary.LittleEndian.Uint16(pkt[*idx : *idx+2])
	*idx += 2
	return ans
}

func WriteWordLittleEndian(pkt []byte, num uint16) []byte {
	numPkt := make([]byte, 2)
	binary.LittleEndian.PutUint16(numPkt, num)
	return append(pkt, numPkt...)
}

// -- -- -- -- -- -- --

func ReadThreeBytesLittleEndian(pkt []byte, idx *int) uint32 {
	ans := uint32(pkt[*idx]) | uint32(pkt[*idx+1])<<8 | uint32(pkt[*idx+2])<<16
	*idx += 3
	return ans
}

func WriteThreeBytesLittleEndian(pkt []byte, num uint32) []byte {
	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, num)
	threeBytes := tmp[:3]
	return append(pkt, threeBytes...)
}

// -- -- -- -- -- -- --

func ReadDoubleWordLittleEndian(pkt []byte, idx *int) uint32 {
	ans := binary.LittleEndian.Uint32(pkt[*idx : *idx+4])
	*idx += 4
	return ans
}

func WriteDoubleWordLittleEndian(pkt []byte, num uint32) []byte {
	numPkt := make([]byte, 4)
	binary.LittleEndian.PutUint32(numPkt, num)
	return append(pkt, numPkt...)
}

// -- -- -- -- -- -- --

func ReadUint64LittleEndian(pkt []byte, idx *int) uint64 {
	ans := binary.LittleEndian.Uint64(pkt[*idx : *idx+8])
	*idx += 8
	return ans
}
