package api

import (
	"color-kwing/logger"
	"fmt"
	"strconv"
	"strings"
)

var Logg = logger.Get()

type conversions interface {
	HexToRgb() string
	HexToHsl() string
	HexToHsv() string
}

type hexNumberInidividual struct {
	R [2]Hex
	G [2]Hex
	B [2]Hex
}

type RgbNumbersInidividual struct {
	R uint8
	G uint8
	B uint8
}

const hash rune = '#'

// type HexCode struct {
// 	hash   rune
// 	number hexNumberInidividual
// }

type hexNumbersString [6]Hex
type hexCodeString []rune

func (h hexCodeString) toHex(start, end uint8) [2]Hex {
	// aba := [2]Hex{(rune(h[start])), rune((h[end]))}

	var a, b uint64

	Logg.Info().
		Any("start", rune(h[start])-'0').
		Any("end", rune(h[end])-'0').
		Bool("String", false).Msg("toHex Old")

	Logg.Info().
		Str("start", string(rune(h[start]))).
		Str("end", string(rune(h[end]))).
		Bool("String", true).Msg("toHex New")

	stringToInt := func(startOrEnd uint64, isStart bool) (toReplace uint64) {
		ifstatement := func(hex Hex) uint64 {
			if isStart {
				return uint64(_A * 16)
			} else {
				return uint64(hex)
			}
		}
		switch alphabeticRune := string(rune(h[startOrEnd])); alphabeticRune {
		case "A":
			toReplace = ifstatement(_A)
		case "B":
			toReplace = ifstatement(_B)
		case "C":
			toReplace = ifstatement(_C)
		case "D":
			toReplace = ifstatement(_D)
		case "E":
			toReplace = ifstatement(_E)
		case "F":
			toReplace = ifstatement(_F)
		default:
			var err error
			toReplace, err = strconv.ParseUint(string(rune(h[startOrEnd])), 10, 8)
			if err != nil {
				fmt.Errorf("%v", err)
			}
		}
		return toReplace
	}

	a = stringToInt(uint64(start), true)
	b = stringToInt(uint64(end), false)

	aba := [2]Hex{uint8(a), uint8(b)}

	Logg.Info().Any("aba", aba).Msg("toHex ending")

	return aba
}

func StringToHex(s string) (hexNumberInidividual, error) {
	seperated := strings.Split(s, "#")
	var runes hexCodeString = []rune(seperated[1])

	if len(runes) > 6 || len(runes) < 5 {
		return hexNumberInidividual{}, fmt.Errorf("Not a valid Hex Code: %v")
	}

	shibi := hexNumberInidividual{
		R: runes.toHex(0, 1),
		G: runes.toHex(2, 3),
		B: runes.toHex(4, 5),
	}

	return shibi, nil
}

func (h *hexNumberInidividual) RunesToIntArray() RgbNumbersInidividual {
	var r, g, b uint8
	for i, val := range h.R {
		if i == 0 {
			r = r + uint8(val*16)
			fmt.Printf("r: %v\n", r)
		} else {
			r = r + uint8(val)
		}
	}
	for _, val := range h.G {
		g = g + uint8(val*16)
	}
	for _, val := range h.B {
		b = b + uint8(val*16)
	}

	hibi := RgbNumbersInidividual{
		R: r,
		G: g,
		B: b,
	}

	fmt.Printf("hibi: %v\n", hibi)

	return hibi
}

func NewRgbFromHex(hexCode string) (hexNumberInidividual, error) {
	justNumbers, err := StringToHex(hexCode)

	fmt.Printf("justNumbers: %v\n", justNumbers)

	if err != nil {
		fmt.Errorf("NewHex could not be created: %v", err)
	}
	return justNumbers, nil
}

const (
	_0 = Hex(0x00)
	_1 = Hex(0x01)
	_2 = Hex(0x02)
	_3 = Hex(0x03)
	_4 = Hex(0x04)
	_5 = Hex(0x05)
	_6 = Hex(0x06)
	_7 = Hex(0x07)
	_8 = Hex(0x08)
	_9 = Hex(0x09)
	_A = Hex(0x0A)
	_B = Hex(0x0B)
	_C = Hex(0x0C)
	_D = Hex(0x0D)
	_E = Hex(0x0E)
	_F = Hex(0x0F)
)

type Hex = uint8

// const (
// 	_0 Hex = iota
// 	_1
// 	_2
// 	_3
// 	_4
// 	_5
// 	_6
// 	_7
// 	_8
// 	_9
// 	_A
// 	_B
// 	_C
// 	_D
// 	_E
// 	_F
// )
