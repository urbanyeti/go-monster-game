package main

import (
	"fmt"
)

func main() {

	bo := BattleOrder{}
	bo.Attackers = append(bo.Attackers, Minion{&Stats{spd: 5}})
	bo.Attackers = append(bo.Attackers, Minion{&Stats{spd: 1}})
	bo.Attackers = append(bo.Attackers, Hero{&Stats{spd: 3}})
	bo.Attackers = append(bo.Attackers, Minion{&Stats{spd: 4}})
	bo.Attackers = append(bo.Attackers, Minion{&Stats{spd: 2}})
	bo.Attackers = append(bo.Attackers, Hero{&Stats{spd: 9}})
	bo.Attackers = append(bo.Attackers, Hero{&Stats{spd: 7}})

	bo.Build()
	fmt.Println(bo.Attackers)
	fmt.Println(bo.Active)
	a := bo.Next()

	switch v := a.(type) {
	case Minion:
		fmt.Printf("It's a minion! %v\n", v)
	case Hero:
		fmt.Printf("It's a hero! %v\n", v)
	default:
		fmt.Println("unknown")
	}

	a.SetSpd(1)
	fmt.Println(a)
	fmt.Println(bo.Attackers)
	fmt.Println(bo.Active)
	fmt.Println(bo.InActive)
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
	*Stats
}

func (m Minion) String() string {
	return fmt.Sprintf("Minion - %v", m.spd)
}

type Hero struct {
	*Stats
}

func (h Hero) String() string {
	return fmt.Sprintf("Hero - %v", h.spd)
}
