package main

import "fmt"

type Fighter struct {
	hitp, damage, armour int
}

func Fight(one Fighter, two Fighter) int {
	onehp := one.hitp
	twohp := two.hitp

	onein := true

	for true {
		if onein {
			twohp -= Max(1,one.damage-two.armour)
			if twohp <= 0 {
				return 1
			}
		} else {
			onehp -= Max(1,two.damage-one.armour)
			if onehp <= 0 {
				return 2
			}
		}
		onein = !onein
	}

	return -1
}

func daytwentyone() {
	boss := Fighter{109,8,2}

	best := 99999
	best2 := 0

	weapons := [5]int{8,10,25,40,74}	
	armours := [5]int{13,31,53,75,102}
	ringsd := [3]int{25,50,100}
	ringsa := [3]int{20,40,80}

	for weapon := 0 ; weapon < len(weapons) ; weapon++ {
		for armour := -1 ; armour < len(armours) ; armour++ {
			me := Fighter{100,weapon+4,armour+1}
			
			gold := weapons[weapon]
			if armour >= 0 {
				gold += armours[armour]
			}
			if Fight(me,boss) == 1 {
				best = Min(gold,best)
			} else {
				best2 = Max(gold,best2)
			}
			
			for ringd := -1; ringd < len(ringsd) ; ringd++ {
				for ringa := -1; ringa < len(ringsa) ; ringa++ {
					me := Fighter{100,weapon+4+ringd+1,armour+1+ringa+1}
					gold := weapons[weapon]
					if armour >= 0 {
						gold += armours[armour]
					}
					if ringd >= 0 {
						gold += ringsd[ringd]
					}
					if ringa >= 0 {
						gold += ringsa[ringa]
					}
					if Fight(me,boss) == 1 {
						best = Min(gold,best)
					} else{
						best2 = Max(gold,best2)
					}
				}
			}
		}
	}

	fmt.Printf("Fight 1 %v\n", best)
	fmt.Printf("Fight 2 %v\n", best2)
}
