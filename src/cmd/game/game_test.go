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
	assert.Len(bo.Active, 5)
	assert.Len(bo.InActive, 0)
	assert.Equal(bo.Next().GetName(), "Hero Tim")
	assert.Len(bo.Active, 4)
	assert.Len(bo.InActive, 1)
}

func (a Attackers) ListNames() []string {
	var output []string
	for _, v := range a {
		output = append(output, v.GetName())
	}

	return output
}

func (bo *BattleOrder) AddDefaultAttackers() {
	bo.Attackers = append(bo.Attackers, Minion{
		Name:  "Goblin A",
		Stats: &Stats{spd: 6},
		Attacks: &Attacks{
			Attack{Name: "Bite"},
			Attack{Name: "Scratch"},
		},
		Passives: &Passives{},
	})
	bo.Attackers = append(bo.Attackers, Minion{
		Name:  "Goblin B",
		Stats: &Stats{spd: 5},
		Attacks: &Attacks{
			Attack{Name: "Bite"},
			Attack{Name: "Scratch"},
		},
		Passives: &Passives{},
	})
	bo.Attackers = append(bo.Attackers, Hero{
		Name:  "Hero Tim",
		Stats: &Stats{spd: 7},
		Attacks: &Attacks{
			Attack{Name: "Sword Attack"},
			Attack{Name: "Longbow Shot"},
		},
		Passives: &Passives{
			Passive{Name: "Shield Defense"},
		},
	})
	bo.Attackers = append(bo.Attackers, Minion{
		Name:  "Troll",
		Stats: &Stats{spd: 2},
		Attacks: &Attacks{
			Attack{Name: "Club Attack"},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
	bo.Attackers = append(bo.Attackers, Minion{
		Name:  "Giant",
		Stats: &Stats{spd: 1},
		Attacks: &Attacks{
			Attack{Name: "Stomp"},
			Attack{Name: "Club Attack"},
		},
		Passives: &Passives{
			Passive{Name: "Hard Skin"},
		},
	})
}
