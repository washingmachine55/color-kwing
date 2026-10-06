package api

import (
	"color-kwing/logger"
	"fmt"
	"strconv"
	"strings"
)

var Logg = logger.Get()

type Hex = uint8
type Red = uint8
type Green = uint8
type Blue = uint8

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

type HexToX interface {
	// HexToRgb() string
	// HexToHsl() string
	// HexToHsv() string
	// toHex(start, end uint8) [2]Hex
	// RunesToIntArray() RgbNumbersIndividual
	StringToHex(s string) (HexNumberIndividual, error)
	NewRgbFromHex(hexCode string) (RgbNumbersIndividual, error)
}

type HexNumberIndividual struct {
	R [2]Hex
	G [2]Hex
	B [2]Hex
}

type RgbNumbersIndividual struct {
	R Red
	G Green
	B Blue
}

type NewHexString struct {
	S string
}

const hash rune = '#'

type hexNumbersString [6]Hex
type HexCodeString []rune

func (h HexCodeString) toHex(start, end uint8) [2]Hex {
	var a, b uint64

	Logg.Debug().
		Any("start", rune(h[start])-'0').
		Any("end", rune(h[end])-'0').
		Bool("String", false).Msg("toHex Old")

	Logg.Debug().
		Str("start", string(rune(h[start]))).
		Str("end", string(rune(h[end]))).
		Bool("String", true).Msg("toHex New")

	stringToInt := func(startOrEnd uint64, isStart bool) (toReplace uint64) {
		convertAndMultiply := func(hex Hex) uint64 {
			if isStart {
				return uint64(hex * 16)
			} else {
				return uint64(hex)
			}
		}
		switch alphabeticRune := string(rune(h[startOrEnd])); alphabeticRune {
		case "0":
			toReplace = convertAndMultiply(_0)
		case "1":
			toReplace = convertAndMultiply(_1)
		case "2":
			toReplace = convertAndMultiply(_2)
		case "3":
			toReplace = convertAndMultiply(_3)
		case "4":
			toReplace = convertAndMultiply(_4)
		case "5":
			toReplace = convertAndMultiply(_5)
		case "6":
			toReplace = convertAndMultiply(_6)
		case "7":
			toReplace = convertAndMultiply(_7)
		case "8":
			toReplace = convertAndMultiply(_8)
		case "9":
			toReplace = convertAndMultiply(_9)
		case "A":
			toReplace = convertAndMultiply(_A)
		case "B":
			toReplace = convertAndMultiply(_B)
		case "C":
			toReplace = convertAndMultiply(_C)
		case "D":
			toReplace = convertAndMultiply(_D)
		case "E":
			toReplace = convertAndMultiply(_E)
		case "F":
			toReplace = convertAndMultiply(_F)
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

	Logg.Debug().Any("aba", aba).Msg("toHex ending")

	return aba
}
func (h *HexNumberIndividual) RunesToIntArray() RgbNumbersIndividual {
	// Logg.Debug().Any("R", h.R).Msg("Start of RunesToIntArray")
	// Logg.Debug().Any("G", h.G).Msg("Start of RunesToIntArray")
	// Logg.Debug().Any("B", h.B).Msg("Start of RunesToIntArray")

	var r, g, b uint8
	for _, val := range h.R {
		r = r + val
	}
	for _, val := range h.G {
		g = g + val
	}
	for _, val := range h.B {
		b = b + val
	}

	return RgbNumbersIndividual{
		R: r,
		G: g,
		B: b,
	}
}

func StringToHex(s string) (HexNumberIndividual, error) {
	separated := strings.Split(s, "#")
	var runes HexCodeString = []rune(separated[1])

	if len(runes) > 6 || len(runes) < 5 {
		return HexNumberIndividual{}, fmt.Errorf("Not a valid Hex Code: %v")
	}

	return HexNumberIndividual{
		R: runes.toHex(0, 1),
		G: runes.toHex(2, 3),
		B: runes.toHex(4, 5),
	}, nil
}

func NewRgbFromHex(hexCode string) (RgbNumbersIndividual, error) {
	justNumbers, err := StringToHex(hexCode)
	if err != nil {
		fmt.Errorf("NewHex could not be created: %v", err)
	}

	return justNumbers.RunesToIntArray(), nil
}
