package main

import "math/rand"

type Attacker interface {
	GetName() string
	Spd() int
	SetSpd(int)
}

type BattleOrder struct {
	Attackers Attackers
	Active    Attackers
	InActive  Attackers
}

type Attackers []Attacker

func (b *BattleOrder) Build() {
	b.Attackers.Order()
	b.Active = make(Attackers, len(b.Attackers))
	copy(b.Active, b.Attackers)
	b.InActive = Attackers{}
}

func (b *BattleOrder) Next() Attacker {
	if len(b.Active) == 0 {
		b.Active = make(Attackers, len(b.Attackers))
		copy(b.Active, b.Attackers)
		b.InActive = Attackers{}
	}

	a := b.Active[0]
	b.Active = b.Active[1:]
	b.InActive = append(b.InActive, a)
	return a
}

// Order performs a quick sort of descending speed for attackers
func (a Attackers) Order() {
	if len(a) <= 1 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	a[right], a[pivot] = a[pivot], a[right]

	for i := range a {
		if a[i].Spd() > a[right].Spd() {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	a[:left].Order()
	a[left+1:].Order()
}
