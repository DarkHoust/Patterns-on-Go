package main

import "fmt"

type WeaponBehaviour interface {
	Inspect()
}

type BigWeapon struct{}

func (bgw BigWeapon) Inspect() {
	fmt.Println("I love big weapon, Sir!")
}

type RangedWeapon struct {
	AmountOfBullets int
}

func (rngw RangedWeapon) Inspect() {
	if rngw.AmountOfBullets == 0 {
		fmt.Println("I need more bullets!")
	} else {
		fmt.Println("Bdu Bdu Bdu Bdu!")
	}
}

type MeleeWeapon struct{}

func (mlw MeleeWeapon) Inspect() {
	fmt.Println("Thanks for the weapon, Sir!")
}

type Weapon struct {
	Behavior WeaponBehaviour
}

func WeaponInfo(w Weapon) {
	w.Behavior.Inspect()
}

func main() {
	toyPistol := Weapon{Behavior: RangedWeapon{AmountOfBullets: 0}}
	toyRifle := Weapon{Behavior: RangedWeapon{AmountOfBullets: 10}}
	knife := Weapon{Behavior: MeleeWeapon{}}
	toyHammer := Weapon{Behavior: BigWeapon{}}

	WeaponInfo(toyPistol)
	WeaponInfo(toyRifle)
	WeaponInfo(toyHammer)
	WeaponInfo(knife)
}
