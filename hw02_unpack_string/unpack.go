package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	finalResult := new(strings.Builder)
	stringByRunes := []rune(packedString)
	startIndex := 0
	endIndex := 0
	for i := 0; i < len(stringByRunes); i++ {
		switch isDigit := unicode.IsDigit(stringByRunes[i]); {
		case !isDigit:
			{
				endIndex = i
				blockResult, err := unpackSmallerBlock(stringByRunes[startIndex:endIndex])
				if err != nil {
					return "", err
				}
				finalResult.WriteString(blockResult)
				startIndex = i
				if i == len(stringByRunes)-1 {
					finalResult.WriteString(string(stringByRunes[startIndex:]))
				}
			}
		case isDigit:
			{
				if i == len(stringByRunes)-1 {
					blockResult, err := unpackSmallerBlock(stringByRunes[startIndex:])
					if err != nil {
						return "", ErrInvalidString
					}
					finalResult.WriteString(blockResult)
				}
				continue
			}
		}
	}
	return (*finalResult).String(), nil
}

func unpackSmallerBlock(block []rune) (string, error) {
	switch length := len(block); {
	case length == 0:
		{
			return "", nil
		}
	case length == 1:
		{
			if unicode.IsDigit(block[0]) {
				return "", ErrInvalidString
			}
			return string(block[0]), nil
		}
	case length == 2:
		{
			if !unicode.IsDigit(block[0]) {
				repeatCount, _ := strconv.Atoi(string(block[1]))
				return strings.Repeat(string(block[0]), repeatCount), nil
			}
			return "", ErrInvalidString
		}
	default:
		{
			return "", ErrInvalidString
		}
	}
}
