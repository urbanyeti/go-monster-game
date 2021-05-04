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
	hp  int
	pow int
	spd int
	def int
	Resistance
}

type Resistance struct {
	fire   int
	ice    int
	armor  int
	nature int
	air    int
}

func (r *Resistance) Fire() int {
	return r.fire
}

func (r *Resistance) SetFire(v int) {
	r.fire = v
}

func (r *Resistance) Ice() int {
	return r.ice
}

func (r *Resistance) SetIce(v int) {
	r.ice = v
}

func (r *Resistance) Armor() int {
	return r.armor
}

func (r *Resistance) SetArmor(v int) {
	r.armor = v
}

func (r *Resistance) Nature() int {
	return r.nature
}

func (r *Resistance) SetNature(v int) {
	r.nature = v
}

func (r *Resistance) Air() int {
	return r.air
}

func (r *Resistance) SetAir(v int) {
	r.air = v
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
