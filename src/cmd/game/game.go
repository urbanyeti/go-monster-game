package main

import (
	"fmt"
)

func main() {
}

type Attacks []Attack

type Attack struct {
	Name string
	Desc string
}

type Passives []Passive

type Passive struct {
	Name string
	Desc string
}

type Stats struct {
	spd int
}

func (s *Stats) Spd() int {
	return s.spd
}

func (s *Stats) SetSpd(val int) {
	s.spd = val
}

type Minion struct {
	Name string
	*Attacks
	*Passives
	*Stats
}

func (m Minion) GetName() string {
	return m.Name
}

func (m Minion) String() string {
	return fmt.Sprintf("%v (Minion) - %v", m.Name, m.spd)
}

type Hero struct {
	Name string
	*Attacks
	*Passives
	*Stats
}

func (h Hero) GetName() string {
	return h.Name
}

func (h Hero) String() string {
	return fmt.Sprintf("%v (Hero) - %v", h.Name, h.spd)
}
