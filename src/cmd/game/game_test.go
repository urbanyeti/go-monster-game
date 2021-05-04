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
	assert.Equal(bo.Next().GetName(), "Hero Tim")
	assert.Len(bo.Active, 4)
	assert.Len(bo.InActive, 1)
	assert.Equal(bo.Next().GetName(), "Goblin A")
	assert.Len(bo.Active, 3)
	assert.Len(bo.InActive, 2)
	assert.Equal(bo.Next().GetName(), "Goblin B")
	assert.Len(bo.Active, 2)
	assert.Len(bo.InActive, 3)
	assert.Equal(bo.Next().GetName(), "Troll")
	assert.Len(bo.Active, 1)
	assert.Len(bo.InActive, 4)
	assert.Equal(bo.Next().GetName(), "Giant")
	assert.Len(bo.Active, 0)
	assert.Len(bo.InActive, 5)
	assert.Equal(bo.Next().GetName(), "Hero Tim")
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

func TestBattleOrder_RemoveAttacker(t *testing.T) {
	assert := assert.New(t)
	bo := BattleOrder{}
	bo.AddDefaultAttackers()
	bo.Build()
	bo.RemoveAttacker(bo.Minions[1])
	assert.Equal(
		bo.Attackers.ListNames(),
		[]string{"Hero Tim", "Goblin A", "Troll", "Giant"},
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
		Stats: &Stats{spd: 6},
		Attacks: &Attacks{
			Attack{Name: "Scratch", Dmg: 5, Element: Physical},
			Attack{Name: "Grenade", Dmg: 5, Element: Fire},
		},
		Passives: &Passives{},
	})
	bo.Add(Minion{
		Name:  "Goblin B",
		Stats: &Stats{spd: 5},
		Attacks: &Attacks{
			Attack{Name: "Bite", Dmg: 10, Element: Physical},
			Attack{Name: "Scratch", Dmg: 5, Element: Physical},
		},
		Passives: &Passives{},
	})
	bo.Add(Hero{
		Name:  "Hero Tim",
		Stats: &Stats{spd: 7},
		Attacks: &Attacks{
			Attack{Name: "Sword Attack", Dmg: 20, Element: Physical},
			Attack{Name: "Longbow Shot", Dmg: 10, Element: Physical},
		},
		Passives: &Passives{
			Passive{Name: "Shield Defense"},
		},
	})
	bo.Add(Minion{
		Name:  "Troll",
		Stats: &Stats{spd: 2},
		Attacks: &Attacks{
			Attack{Name: "Club Attack", Dmg: 10, Element: Physical},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
	bo.Add(Minion{
		Name:  "Giant",
		Stats: &Stats{spd: 1},
		Attacks: &Attacks{
			Attack{Name: "Stomp", Dmg: 15, Element: Nature},
			Attack{Name: "Club Attack", Dmg: 10, Element: Physical},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
}
