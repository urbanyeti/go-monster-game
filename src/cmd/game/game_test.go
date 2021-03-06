package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBattleOrder_Order(t *testing.T) {
	assert := assert.New(t)
	bo := BattleOrder{}
	bo.AddDefaultAttackers()
	bo.Build()
	assert.Equal(
		bo.Attackers.ListNames(),
		[]string{"Hero Tim", "Goblin A", "Goblin B", "Troll", "Giant"},
		"The list should be ordered descending by speed.",
	)
}

func TestBattleOrder_Next(t *testing.T) {
	assert := assert.New(t)
	bo := BattleOrder{}
	bo.AddDefaultAttackers()
	bo.Build()
	assert.Len(bo.Active, 5)
	assert.Len(bo.InActive, 0)
	assert.Equal("Hero Tim", bo.Next().GetName())
	assert.Len(bo.Active, 4)
	assert.Len(bo.InActive, 1)
	assert.Equal("Goblin A", bo.Next().GetName())
	assert.Len(bo.Active, 3)
	assert.Len(bo.InActive, 2)
	assert.Equal("Goblin B", bo.Next().GetName())
	assert.Len(bo.Active, 2)
	assert.Len(bo.InActive, 3)
	assert.Equal("Troll", bo.Next().GetName())
	assert.Len(bo.Active, 1)
	assert.Len(bo.InActive, 4)
	assert.Equal("Giant", bo.Next().GetName())
	assert.Len(bo.Active, 0)
	assert.Len(bo.InActive, 5)
	assert.Equal("Hero Tim", bo.Next().GetName())
	assert.Len(bo.Active, 4)
	assert.Len(bo.InActive, 1)
}

func TestBattleOrder_GetResistance(t *testing.T) {
	assert := assert.New(t)
	dmg := 80
	m := Minion{Stats: &Stats{Resistance: Resistance{
		fire:   15,
		ice:    100,
		armor:  -20,
		nature: -200,
		air:    0,
	}}}

	assert.Equal(12, GetResistance(dmg, Fire, m))
	assert.Equal(80, GetResistance(dmg, Ice, m))
	assert.Equal(-16, GetResistance(dmg, Physical, m))
	assert.Equal(-160, GetResistance(dmg, Nature, m))
	assert.Equal(0, GetResistance(dmg, Air, m))
}

func TestBattleOrder_KillAttacker(t *testing.T) {
	assert := assert.New(t)
	bo := BattleOrder{}
	bo.AddDefaultAttackers()
	bo.Build()
	bo.KillAttacker(bo.Minions[1])
	assert.Equal(
		[]string{"Hero Tim", "Goblin A", "Troll", "Giant"},
		bo.Attackers.ListNames(),
	)
	assert.Equal(
		[]string{"Hero Tim", "Goblin A", "Troll", "Giant"},
		bo.Active.ListNames(),
	)
	assert.Equal(
		[]string{"Goblin B"},
		bo.Dead.ListNames(),
	)
}

func TestBattleOrder_Attack(t *testing.T) {
	assert := assert.New(t)
	bo := BattleOrder{}
	bo.AddDefaultAttackers()
	bo.Build()
	source, target := bo.Heroes[0], bo.Minions[0]
	assert.Equal(5, target.HP())
	bo.Attack(&source, &target, source.SelectAttack(0))
	assert.Equal(-9, target.HP())

	assert.Equal(
		[]string{"Hero Tim", "Goblin B", "Troll", "Giant"},
		bo.Attackers.ListNames(),
	)
	assert.Equal(
		[]string{"Hero Tim", "Goblin B", "Troll", "Giant"},
		bo.Active.ListNames(),
	)
	assert.Equal(
		[]string{"Goblin B", "Troll", "Giant"},
		bo.Minions.ListNames(),
	)
	assert.Equal(
		[]string{"Goblin A"},
		bo.Dead.ListNames(),
	)
}

func TestStats_SetHP(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.hp = 10
	assert.Equal(10, s.HP())
	s.SetHP(-100)
	assert.Equal(-100, s.HP())
}

func TestStats_SetPow(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.pow = 10
	assert.Equal(10, s.Pow())
	s.SetPow(-100)
	assert.Equal(-100, s.Pow())
}

func TestStats_SetSpd(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.spd = 10
	assert.Equal(10, s.Spd())
	s.SetSpd(-100)
	assert.Equal(-100, s.Spd())
}

func TestStats_SetFire(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.fire = 10
	assert.Equal(10, s.Fire())
	s.SetFire(-100)
	assert.Equal(-100, s.Fire())
}

func TestStats_SetIce(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.ice = 10
	assert.Equal(10, s.Ice())
	s.SetIce(-100)
	assert.Equal(-100, s.Ice())
}

func TestStats_SetArmor(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.armor = 10
	assert.Equal(10, s.Armor())
	s.SetArmor(-100)
	assert.Equal(-100, s.Armor())
}

func TestStats_SetNature(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.nature = 10
	assert.Equal(10, s.Nature())
	s.SetNature(-100)
	assert.Equal(-100, s.Nature())
}

func TestStats_SetAir(t *testing.T) {
	assert := assert.New(t)
	s := &Stats{}
	s.air = 10
	assert.Equal(10, s.Air())
	s.SetAir(-100)
	assert.Equal(-100, s.Air())
}

func (a Attackers) ListNames() []string {
	var output []string
	for _, v := range a {
		output = append(output, v.GetName())
	}

	return output
}

func (bo *BattleOrder) AddDefaultAttackers() {
	bo.Add(Minion{
		Name:  "Goblin A",
		Stats: &Stats{hp: 5, pow: 1, spd: 6},
		Attacks: &Attacks{
			Attack{Name: "Scratch", Damages: []Damage{{Dmg: 5, Element: Physical}, {Dmg: 1, Element: Nature}}},
			Attack{Name: "Grenade", Damages: []Damage{{Dmg: 5, Element: Fire}}},
		},
		Passives: &Passives{},
	})
	bo.Add(Minion{
		Name:  "Goblin B",
		Stats: &Stats{hp: 7, pow: 2, spd: 5},
		Attacks: &Attacks{
			Attack{Name: "Bite", Damages: []Damage{{Dmg: 10, Element: Physical}}},
			Attack{Name: "Scratch", Damages: []Damage{{Dmg: 5, Element: Physical}, {Dmg: 1, Element: Nature}}},
		},
		Passives: &Passives{},
	})
	bo.Add(Hero{
		Name:  "Hero Tim",
		Stats: &Stats{hp: 30, pow: 2, spd: 7},
		Attacks: &Attacks{
			Attack{Name: "Sword Attack", Damages: []Damage{{Dmg: 12, Element: Physical}}},
			Attack{Name: "Longbow Shot", Damages: []Damage{{Dmg: 8, Element: Physical}, {Dmg: 2, Element: Air}}},
		},
		Passives: &Passives{
			Passive{Name: "Shield Defense"},
		},
	})
	bo.Add(Minion{
		Name:  "Troll",
		Stats: &Stats{hp: 15, pow: 5, spd: 2},
		Attacks: &Attacks{
			Attack{Name: "Club Attack", Damages: []Damage{{Dmg: 10, Element: Physical}}},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
	bo.Add(Minion{
		Name:  "Giant",
		Stats: &Stats{hp: 20, pow: 5, spd: 1},
		Attacks: &Attacks{
			Attack{Name: "Stomp", Damages: []Damage{{Dmg: 15, Element: Nature}}},
			Attack{Name: "Club Attack", Damages: []Damage{{Dmg: 10, Element: Physical}}},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
}
