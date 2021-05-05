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
	SelectAttack(int) *Attack
	SelectRandomAttack() *Attack
}

type BattleOrder struct {
	Attackers Attackers
	Active    Attackers
	InActive  Attackers
	Minions   Attackers
	Heroes    Attackers
	Dead      Attackers
	Lookup    OrderLookup
}

type Attackers []Attacker

type OrderLookup map[Attacker][]int

func (bo *BattleOrder) Add(a Attacker) {
	bo.Attackers = append(bo.Attackers, a)
}

func (bo *BattleOrder) Build() {
	bo.Lookup = OrderLookup{}
	bo.InActive = Attackers{}
	bo.Heroes = Attackers{}
	bo.Minions = Attackers{}
	bo.Dead = Attackers{}

	bo.Attackers.Order()

	bo.Active = make(Attackers, len(bo.Attackers))
	for i, a := range bo.Attackers {
		bo.Lookup[a] = []int{i, i, -1, -1, -1, -1}
		bo.Active[i] = a

		switch a.(type) {
		case Minion:
			bo.Lookup[a][3] = len(bo.Minions)
			bo.Minions = append(bo.Minions, a)
		case Hero:
			bo.Lookup[a][4] = len(bo.Heroes)
			bo.Heroes = append(bo.Heroes, a)
		}
	}

}

func (bo *BattleOrder) Next() Attacker {
	if len(bo.Active) == 0 {
		bo.Attackers.Order()

		bo.Active = make(Attackers, len(bo.Attackers))
		for i, a := range bo.Attackers {
			bo.Lookup[a][0], bo.Lookup[a][1], bo.Lookup[a][2] = i, i, -1
			bo.Active[i] = a
		}
		bo.InActive = Attackers{}
	}

	a := bo.Active[0]
	bo.Lookup[a][1] = -1
	bo.Active = bo.Active[1:]
	bo.InActive = append(bo.InActive, a)
	bo.Lookup[a][2] = len(bo.InActive)
	return a
}

func (bo *BattleOrder) Attack(source *Attacker, target *Attacker, a *Attack) {
	if a == nil {
		a = (*source).SelectRandomAttack()
	}
	for _, d := range a.Damages {
		totalDmg := d.Dmg + (*source).Pow()
		totalDmg -= GetResistance(totalDmg, d.Element, *target)
		(*target).SetHP((*target).HP() - totalDmg)
	}

	if (*target).HP() <= 0 {
		bo.KillAttacker(*target)
	}
}

func (bo *BattleOrder) KillAttacker(a Attacker) {
	bo.Attackers.RemoveAttacker(a, bo.Lookup[a][0])
	bo.Active.RemoveAttacker(a, bo.Lookup[a][1])
	bo.InActive.RemoveAttacker(a, bo.Lookup[a][2])
	switch (a).(type) {
	case Minion:
		bo.Minions.RemoveAttacker(a, bo.Lookup[a][3])
	case Hero:
		bo.Heroes.RemoveAttacker(a, bo.Lookup[a][4])
	}
	bo.Lookup[a][5] = len(bo.Dead)
	bo.Dead = append(bo.Dead, a)
}

func (al *Attackers) RemoveAttacker(a Attacker, pos int) {
	if pos == -1 {
		return
	}
	for j := pos; j < len(*al)-1; j++ {
		(*al)[j] = (*al)[j+1]
	}
	*al = (*al)[:len(*al)-1]
	return
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
