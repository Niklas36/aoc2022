package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
)

//go:embed input.txt
var input []byte

type Column []byte

func (s *Column) PushN(b ...byte) {
	*s = append(*s, b...)
}

func (s *Column) PopN(n int) Column {
	off := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return off
}

type Warehouse []Column

func (w *Warehouse) ApplyV9000(m Move) {
	w.expand(m.DstCol)
	for i := 0; i < m.Quantity; i++ {
		cargo := (*w)[m.SrcCol].PopN(1)
		(*w)[m.DstCol].PushN(cargo...)
	}
}

func (w *Warehouse) ApplyV9001(m Move) {
	w.expand(m.DstCol)
	cargo := (*w)[m.SrcCol].PopN(m.Quantity)
	(*w)[m.DstCol].PushN(cargo...)
}

func (w *Warehouse) expand(col int) {
	required := col - len(*w)
	for i := 0; i < required; i++ {
		*w = append(*w, Column{})
	}
}

func (w *Warehouse) String() string {
	tmp := ""
	for _, col := range *w {
		tmp += string(col.PopN(1))
	}

	return tmp
}

type Move struct {
	Quantity int
	SrcCol   int
	DstCol   int
}

func NewWarehouse() Warehouse {
	return Warehouse{
		{'W', 'M', 'L', 'F'},
		{'B', 'Z', 'V', 'M', 'F'},
		{'H', 'V', 'R', 'S', 'L', 'Q'},
		{'F', 'S', 'V', 'Q', 'P', 'M', 'T', 'J'},
		{'L', 'S', 'W'},
		{'F', 'V', 'P', 'M', 'R', 'J', 'W'},
		{'J', 'Q', 'C', 'P', 'N', 'R', 'F'},
		{'V', 'H', 'P', 'S', 'Z', 'W', 'R', 'B'},
		{'B', 'M', 'J', 'C', 'G', 'H', 'Z', 'W'},
	}
}

func parseMoves() []Move {
	var moves []Move
	r := bytes.NewReader(input)
	for {
		var m Move
		if _, err := fmt.Fscanf(r, "move %d from %d to %d\n", &m.Quantity, &m.SrcCol, &m.DstCol); err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		m.SrcCol--
		m.DstCol--
		moves = append(moves, m)
	}

	return moves
}

func part1() {
	wh := NewWarehouse()
	for _, move := range parseMoves() {
		wh.ApplyV9000(move)
	}

	fmt.Printf("part1: %s\n", wh.String())
}

func part2() {
	wh := NewWarehouse()
	for _, move := range parseMoves() {
		wh.ApplyV9001(move)
	}

	fmt.Printf("part2: %s\n", wh.String())
}

func main() {

	part1()
	part2()

}
