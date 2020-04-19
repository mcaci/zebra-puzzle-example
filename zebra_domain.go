package zebra

import (
	"strconv"
)

type domainer interface {
	toDomainVal() uint8
}

type nationality uint8

const (
	norwegian nationality = iota + 1
	spaniard
	englishman
	ukrainian
	japanese
)

func (n nationality) toDomainVal() uint8 { return uint8(n) }
func (n nationality) String() string {
	return []string{"-", "Norwegian", "Spaniard", "Englishman", "Ukrainian", "Japanese"}[n]
}

func ownerOf(item domainer, assignment ...uint8) string {
	index, _ := indexOf(item, assignment...)
	return nationality(assignment[index%5]).String()
}

type color uint8

const (
	green color = iota + 1
	ivory
	red
	yellow
	blue
)

func (c color) toDomainVal() uint8 { return uint8(c) }
func (c color) String() string     { return []string{"-", "green", "ivory", "red", "yellow", "blue"}[c] }

type drink uint8

const (
	water drink = iota + 1
	coffee
	milk
	tea
	orangeJuice
)

func (d drink) toDomainVal() uint8 { return uint8(d) }
func (d drink) String() string {
	return []string{"-", "water", "coffee", "milk", "tea", "orange juice"}[d]
}

type smoke uint8

const (
	oldGold smoke = iota + 1
	chesterfields
	kools
	parliaments
	luckyStrike
)

func (s smoke) toDomainVal() uint8 { return uint8(s) }
func (s smoke) String() string {
	return []string{"Old Gold", "Chesterfields", "Kools", "Parliaments", "Lucky Strike"}[s]
}

type pets uint8

const (
	snails pets = iota + 1
	dog
	fox
	horse
	zebra
)

func (p pets) toDomainVal() uint8 { return uint8(p) }
func (p pets) String() string     { return []string{"-", "snails", "dog", "fox", "horse", "zebra"}[p] }

type number uint8

const (
	house1 number = iota + 1
	house2
	house3
	house4
	house5
)

func (n number) toDomainVal() uint8 { return uint8(n) }
func (n number) String() string {
	if n == 0 {
		return "-"
	}
	return strconv.Itoa(int(n))
}

func category(d domainer) int {
	switch d.(type) {
	case nationality:
		return 0
	case color:
		return 1
	case drink:
		return 2
	case smoke:
		return 3
	case pets:
		return 4
	case number:
		return 5
	default:
		return -1
	}
}
