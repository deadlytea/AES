package aes

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type AES struct {
  // Nb should be const 4
	Nb     int
	Nk     int
	Nr     int
	cipher int
	state  []byte
	// Need a word type here
	w []byte
}

func (a *AES) EncryptBlock(in []byte) (out []byte) {

	// TODO: Change this to copy slice
	state := in

	a.AddRoundKey(0)

	for round := 1; round < a.Nr; round++ {
		a.SubBytes()
		a.ShiftRows()
		a.MixColumns()
		a.AddRoundKey(round)
	}

	a.SubBytes()
	a.ShiftRows()
	a.AddRoundKey(a.Nr)

	// Need to initialize?
	out = state
	return
}

func (a *AES) SubBytes() {

}

func (a *AES) MixColumns() {

}

func (a *AES) RotWord() {

}

func (a *AES) SubWord() {

}

func (a *AES) AddRoundKey(r int) {

}

func (a *AES) InvSubBytes() {

}

func (a *AES) InvShiftRows() {

}

func (a *AES) InvMixColumns() {

}
