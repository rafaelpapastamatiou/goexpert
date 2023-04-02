package math

type Math struct {
	A int
	B int
	c int
}

func (m Math) Add() int {
	return m.A + m.B + m.c
}

type math2 struct {
	A int
	B int
	c int
}

func NewMath2(a, b int) *math2 {
	return &math2{A: a, B: b, c: 0}
}
