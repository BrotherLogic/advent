package main

import "fmt"

type MFighter struct {
	hitp, damage, mana, armour, spent int
}


type Effect struct {
	name string
	armour int
	damage int
	mana int
	tick int
	active bool
}

type Spell struct {
	name string
	hitp int
	deal int
	cost int
	effect Effect
}

func EffectActive(eff Effect, effects []Effect) bool {
	for i := 0 ; i < len(effects) ; i++ {
		if effects[i].tick > 1 && effects[i].name == eff.name {
			return true
		}
	}
	return false
}

func ApplyEffect(boss MFighter, me MFighter, eff Effect) (MFighter,MFighter,Effect) {
	if eff.tick > 0 {
		multiplier := 2
		if !eff.active || eff.tick == -1 || eff.tick == 1{
			multiplier = 1
			eff.active = true
		}

		eff.tick = eff.tick-multiplier
		
		if eff.name == "Poison" {
			boss.hitp -= eff.damage*multiplier
		} else if eff.name == "Recharge" {
			me.mana += eff.mana*multiplier
		} else if eff.name == "Shield" {
			if eff.tick == 5 {
				me.armour += eff.armour
			} else if eff.tick <= 0 {
				me.armour -= eff.armour
			}
		}
	}

	return boss,me,eff
}

func MagicFight(boss MFighter, me MFighter, spells []Spell, neffects []Effect, round int, hard bool) int {
	//My Turn
	best := 99999999
	for i := 0 ; i < len(spells) ; i++ {
		effects := make([]Effect,len(neffects))
		copy(effects,neffects)
		if me.mana >= spells[i].cost && !EffectActive(spells[i].effect,effects){
			tboss := boss
			tme := me

			if hard {
				tme.hitp -= 1
			}
			
			tboss.hitp -= spells[i].deal
			tme.hitp += spells[i].hitp
			tme.mana -= spells[i].cost
			tme.spent += spells[i].cost

			if spells[i].effect.name != "" {
				effects = append(effects,spells[i].effect)
			}

			for j := 0; j < len(effects) ; j++ {
				if (effects[j].tick > 0) {
					var neff Effect
					tboss,tme,neff = ApplyEffect(tboss,tme,effects[j])
					effects[j] = neff
				}
			}

			if tboss.hitp <= 0 {
				return tme.spent
			}

			tme.hitp -= tboss.damage - tme.armour

			if tme.hitp > 0 {
				best = Min(best, MagicFight(tboss,tme,spells,effects,round+1, hard))
			}
		}
	}

	return best
}

func daytwentytwo() {
	boss := MFighter{hitp:71,damage:10}
	me := MFighter{hitp:50,mana:500}
	spells := make([]Spell,0)
	effects := make([]Effect,0)

	magicmissile := Spell{cost:53, name:"Magic Missile", deal:4}
	spells = append(spells,magicmissile)
	
	drain := Spell{cost:73, name:"Drain", deal:2, hitp:2}
	spells = append(spells,drain)

	seffect := Effect{name:"Shield",tick:6,armour:7}
	shield := Spell{cost:113, name:"Shield", effect:seffect}
	spells = append(spells,shield)

	peffect := Effect{name:"Poison",tick:6,damage:3}
	poison:= Spell{cost:173, name:"Poison", effect:peffect}
	spells = append(spells,poison)
	
	reffect := Effect{name:"Recharge",tick:5,mana:101}
	recharge := Spell{cost:229, name:"Recharge", effect:reffect}
	spells = append(spells,recharge)

	fmt.Printf("Result = %v\n",MagicFight(boss,me,spells,effects,1, false))
	fmt.Printf("Result = %v\n",MagicFight(boss,me,spells,effects,1, true))
}
