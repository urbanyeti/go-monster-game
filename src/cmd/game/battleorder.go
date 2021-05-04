package main

import (
	"math"
	"math/rand"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Attacker interface {
	GetName() string
	HP() int
	SetHP(int)
	Pow() int
	SetPow(int)
	Spd() int
	SetSpd(int)
	Fire() int
	SetFire(int)
	Ice() int
	SetIce(int)
	Armor() int
	SetArmor(int)
	Nature() int
	SetNature(int)
	Air() int
	SetAir(int)
	SelectAttack() Attack
}

type BattleOrder struct {
	Attackers Attackers
	Active    Attackers
	InActive  Attackers
	Heroes    Attackers
	Minions   Attackers
}

type Attackers []Attacker

func (bo *BattleOrder) Add(a Attacker) {
	bo.Attackers = append(bo.Attackers, a)
	switch x := a.(type) {
	case Minion:
		bo.Minions = append(bo.Minions, x)
	case Hero:
		bo.Heroes = append(bo.Heroes, x)
	}
}

func (b *BattleOrder) Build() {
	b.Attackers.Order()
	b.Active = make(Attackers, len(b.Attackers))
	copy(b.Active, b.Attackers)
	b.InActive = Attackers{}
}

func (bo *BattleOrder) Next() Attacker {
	if len(bo.Active) == 0 {
		bo.Active = make(Attackers, len(bo.Attackers))
		copy(bo.Active, bo.Attackers)
		bo.InActive = Attackers{}
	}

	a := bo.Active[0]
	bo.Active = bo.Active[1:]
	bo.InActive = append(bo.InActive, a)
	return a
}

func (bo *BattleOrder) Attack(source Attacker, target Attacker) {
	a := source.SelectAttack()
	totalDmg := a.Dmg + source.Pow()
	totalDmg -= GetResistance(totalDmg, a.Element, target)
	target.SetHP(target.HP() - totalDmg)
}

func GetResistance(dmg int, element Element, target Attacker) int {
	res := dmg
	switch element {
	case None:
		return dmg
	case Fire:
		res = int(math.Round(float64(res) * (float64(target.Fire()) / float64(100))))
	case Ice:
		res = int(math.Round(float64(res) * (float64(target.Ice())) / float64(100)))
	case Physical:
		res = int(math.Round(float64(res) * (float64(target.Armor()) / float64(100))))
	case Nature:
		res = int(math.Round(float64(res) * (float64(target.Nature()) / float64(100))))
	case Air:
		res = int(math.Round(float64(res) * (float64(target.Air())) / float64(100)))
	}
	return res
}

func (bo *BattleOrder) SelectTargetFor(a Attacker) Attacker {
	switch a.(type) {
	case Minion:
		return bo.Heroes.MostHP()
	case Hero:
		return bo.Minions[rand.Int()%len(bo.Minions)]
	}
	return Minion{}
}

func (al Attackers) MostHP() Attacker {
	maxHP := al[0]
	for _, a := range al {
		if a.HP() > maxHP.HP() {
			maxHP = a
		}
	}
	return maxHP
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
