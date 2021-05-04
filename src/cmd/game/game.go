package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var bo BattleOrder
	bo = append(bo, Minion{Stats{spd: 5}})
	bo = append(bo, Minion{Stats{spd: 1}})
	bo = append(bo, Hero{Stats{spd: 3}})
	bo = append(bo, Minion{Stats{spd: 4}})
	bo = append(bo, Minion{Stats{spd: 2}})
	bo = append(bo, Hero{Stats{spd: 9}})
	bo = append(bo, Hero{Stats{spd: 7}})

	bo.Order()
	fmt.Println(bo)
}

type Stats struct {
	spd int
}

func (s Stats) Spd() int {
	return s.spd
}

type Minion struct {
	Stats
}

func (m Minion) String() string {
	return fmt.Sprintf("Minion - %v", m.spd)
}

type Hero struct {
	Stats
}

func (h Hero) String() string {
	return fmt.Sprintf("Hero - %v", h.spd)
}

type Attacker interface {
	Spd() int
}

type BattleOrder []Attacker

// Order performs a quick sort of descending speed for attackers
func (bo BattleOrder) Order() {
	if len(bo) <= 1 {
		return
	}

	left, right := 0, len(bo)-1
	pivot := rand.Int() % len(bo)

	bo[right], bo[pivot] = bo[pivot], bo[right]

	for i := range bo {
		if bo[i].Spd() > bo[right].Spd() {
			bo[i], bo[left] = bo[left], bo[i]
			left++
		}
	}

	bo[left], bo[right] = bo[right], bo[left]
	bo[:left].Order()
	bo[left+1:].Order()
}
