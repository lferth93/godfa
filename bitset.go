package main

import (
	"fmt"
	"strings"
)

type bitset struct {
	data []uint8
	size int
}

func newBitset(n int) *bitset {
	return &bitset{data: make([]uint8, (n+7)/8), size: n}
}

func (b *bitset) clone() *bitset {
	var bs bitset
	bs.size = b.size
	bs.data = make([]uint8, len(b.data))
	copy(bs.data, b.data)
	return &bs
}

func (b *bitset) set(n int) {
	if 0 <= n && n < b.size {
		b.data[n/8] |= (0x80 >> (n % 8))
	}
}

func (b *bitset) unset(n int) {
	if 0 <= n && n < b.size {
		b.data[n/8] &= ^(0x80 >> (n % 8))
	}
}

func (b *bitset) isSet(n int) bool {
	if 0 <= n && n < b.size {
		return b.data[n/8]&(0x80>>(n%8)) > 0
	}
	return false
}

func (b *bitset) union(b2 *bitset) *bitset {
	if b2 != nil && b.size == b2.size {
		for i := range b.data {
			b.data[i] |= b2.data[i]
		}
	}
	return b
}

func (b *bitset) intersection(b2 *bitset) *bitset {
	if b2 != nil && b.size == b2.size {
		for i := range b.data {
			b.data[i] &= b2.data[i]
		}
	}
	return b
}

func (b *bitset) isEmpty() bool {
	for _, bt := range b.data {
		if bt > 0 {
			return false
		}
	}
	return true
}

func (b *bitset) String() string {
	var sb strings.Builder
	for i := 0; i < b.size/8; i++ {
		sb.WriteString(fmt.Sprintf("%08b", b.data[i]))
	}
	for n, m := b.size%8, uint8(0x80); n > 0; n-- {
		if b.data[len(b.data)-1]&m > 0 {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
		m >>= 1
	}
	return sb.String()
}

func (b *bitset) array() []int {
	var a []int
	for i, m := 0, uint8(0x80); i < b.size; i++ {
		if b.data[i/8]&m > 0 {
			a = append(a, i)
		}
		m >>= 1
		if m == 0 {
			m = 0x80
		}
	}
	return a
}
