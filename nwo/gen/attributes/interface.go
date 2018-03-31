package attributes

import (
	"fmt"
)

//Rated a raated aspect of a character
type Rated interface {
	Level() Level
}

//Named an aspect of teh character that has a name
type Named interface {
	Name() string
}

type Stringer interface {
	String() string
}

//Attribute a character attribute
type Attribute interface {
	Named
	Rated
	Stringer
}

type attr struct {
	name  string
	level Level
}

func (a *attr) Name() string {
	return a.name
}

func (a *attr) Level() Level {
	return a.level
}

func (a *attr) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.level.String())
}

//ArrayOne Attribute Generation Choice One
var ArrayOne = []Level{D4, D6, D6, D8, D10}

//ArrayTwo Attribute Generation Choice Two
var ArrayTwo = []Level{D4, D4, D6, D10, D12}

//AttributeNames List Attribute Names
func AttributeNames() []string {
	return []string{"Nerve", "Suave", "Pulse", "Intellect", "Reflex"}
}

//RandomAtttributes Generate Random Attributes
func RandomAtttributes(generator func() Level) []Attribute {
	names := AttributeNames()
	atts := make([]Attribute, len(names))
	for i, name := range names {
		atts[i] = &attr{name: name, level: generator()}
	}
	return atts
}

//Level a level rating [d4,d12+d4]
type Level int

const (
	_ Level = iota * 2
	_
	//D4 d4
	D4
	//D6 d6
	D6
	//D8 d8
	D8
	//D10 d10
	D10
	//D12 d12
	D12
	//D12D4 d12+d4
	D12D4
)

//String human readable level
func (l Level) String() string {
	switch l {
	case D4:
		return "d4"
	case D6:
		return "d6"
	case D8:
		return "d8"
	case D10:
		return "d10"
	case D12:
		return "d12"
	case D12D4:
		return "d12+d4"
	}
	return ""
}

//HalfRoundedUp calculate a new level: l2 = ceil(l/2)
func (l Level) HalfRoundedUp() Level {
	return ToLevelRoundUp((int(l) / 2))
}

//ToLevelRoundUp convert a calculated value to a Level
func ToLevelRoundUp(i int) Level {
	var l Level
	if i <= int(D4) {
		l = D4
	} else if i <= int(D6) {
		l = D6
	} else if i <= int(D8) {
		l = D8
	} else if i <= int(D10) {
		l = D10
	} else if i <= int(D12) {
		l = D12
	} else {
		l = D12D4
	}
	return l
}
