package main

import (
	"Hyrule_Castle/base_game/gamefunc"
	"fmt"
	"os"
	"strings"
)

type Perso struct {
	Name  string
	Hp    int
	Str   int
	MaxHp int
}
type Boss struct {
	Name  string
	Hp    int
	Str   int
	MaxHp int
}
type Enemy struct {
	Name  string
	Hp    int
	Str   int
	MaxHp int
}

func main() {

	Perso := Perso{"Link", 60, 15, 60}

	//////////////////////////////////

	Boss := Boss{"Ganon", 150, 20, 150}

	//////////////////////////////////

	Enemies := Enemy{"Bokoblin", 30, 5, 30}

	///////////////////////////////////

	var Choice string
	var counter int

	Heal_Point := Perso.Hp / 2

	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("|WELCOME IN THE HYRULE CASTLE !")
	fmt.Println("|Your Player is", Perso.Name, ". You have", Perso.Hp, "Hp and", Perso.Str, "Str.")
	fmt.Println("|You Have to fight 9", Enemies.Name, "in 9 different floor to be able to fight with the FINAL BOSS	")
	fmt.Println("|The Final Boss will be", Boss.Name, ", he have", Boss.Hp, "Hp and", Boss.Str, "Str.")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	Floor := 1
	fmt.Println("You are at floor", Floor)
	fmt.Println("An", Enemies.Name, "is in front of you ?")

	for Floor == 1 {
		if mathfunc.IsEven(counter) {
			fmt.Println("What are you doing ?")
			fmt.Println("1. Attack		2.Heal")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")
			if Choice == "1" {
				fmt.Println("You attack", Enemies.Name)
				Enemies.Hp -= Perso.Str
				fmt.Println("the enemies have now", Enemies.Hp, "Hp")
				fmt.Println("-------------------------------------------------------------------------------------------------")

			}
			if Choice == "2" {
				if Perso.Hp == Perso.MaxHp {
					fmt.Println("You have the maximum of Hp")
					fmt.Println("-------------------------------------------------------------------------------------------------")
				} else {
					fmt.Println("You heal yourself")
					Perso.Hp += Heal_Point
					if Perso.Hp > Perso.MaxHp {
						Perso.Hp = Perso.MaxHp
					}
					fmt.Println("You have now", Perso.Hp, "Hp.")
					fmt.Println("-------------------------------------------------------------------------------------------------")
				}

			}
		} else {
			fmt.Println(Enemies.Name, "attack you")
			Perso.Hp -= Enemies.Str
			fmt.Println("You have now", Perso.Hp, "Hp.")
			fmt.Println("-------------------------------------------------------------------------------------------------")

		}
		counter++

		if Enemies.Hp <= 0 {
			fmt.Println("You have fought the enemies")
			fmt.Println("You can now pass to the next floor, do you want to continue ?(Y/n)")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")

			if strings.ToUpper(Choice) == "Y" {
				Floor++
				counter = 0
				Enemies.Hp = Enemies.MaxHp
			} else if strings.ToLower(Choice) == "n" {
				os.Exit(3)
			}
		}
		if Perso.Hp <= 0 {
			fmt.Println("You die")

			os.Exit(3)

		}
	}

	for Floor < 10 {
		if mathfunc.IsEven(counter) {
			fmt.Println("An", Enemies.Name, "is in front of you ")
			fmt.Println("What are you doing ?")
			fmt.Println("1. Attack		2.Heal")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")
			if Choice == "1" {
				fmt.Println("You attack", Enemies.Name)
				Enemies.Hp -= Perso.Str
				fmt.Println("the enemies have now", Enemies.Hp, "Hp")
				fmt.Println("-------------------------------------------------------------------------------------------------")

			}
			if Choice == "2" {
				if Perso.Hp == Perso.MaxHp {
					fmt.Println("You have the maximum of Hp")
					fmt.Println("-------------------------------------------------------------------------------------------------")

				} else {
					fmt.Println("You heal yourself")
					Perso.Hp += Heal_Point
					if Perso.Hp > Perso.MaxHp {
						Perso.Hp = Perso.MaxHp
					}
					fmt.Println("You have now", Perso.Hp, "Hp.")
					fmt.Println("-------------------------------------------------------------------------------------------------")

				}

			}
		} else {
			fmt.Println(Enemies.Name, "attack you")
			Perso.Hp -= Enemies.Str
			fmt.Println("You have now", Perso.Hp, "Hp.")
			fmt.Println("-------------------------------------------------------------------------------------------------")

		}
		counter++

		if Enemies.Hp <= 0 {
			fmt.Println("You have fought the enemies")
			fmt.Println("You can now pass to the next floor, do you want to continue ?(Y/n)")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")

			if strings.ToUpper(Choice) == "Y" {
				Floor++
				counter = 0
				fmt.Println("You are at floor", Floor)
				fmt.Println("-------------------------------------------------------------------------------------------------")

				Enemies.Hp = Enemies.MaxHp
			} else if strings.ToLower(Choice) == "n" {
				os.Exit(3)

			}
		}

		if Perso.Hp <= 0 {
			fmt.Println("You die")
			os.Exit(3)

		}

	}
	fmt.Println("Welcome to the FINAL BOSS floor !")
	fmt.Println("-------------------------------------------------------------------------------------------------")

	for Floor == 10 {
		if mathfunc.IsEven(counter) {
			fmt.Println("1. Attack		2.Heal")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")

			if Choice == "1" {
				fmt.Println("You attack", Boss.Name)
				Boss.Hp -= Perso.Str
				fmt.Println("the enemies have now", Boss.Hp, "Hp")
				fmt.Println("-------------------------------------------------------------------------------------------------")

			}
			if Choice == "2" {
				if Perso.Hp == Perso.MaxHp {
					fmt.Println("You have the maximum of Hp")
					fmt.Println("-------------------------------------------------------------------------------------------------")

				} else {
					fmt.Println("You heal yourself")
					Perso.Hp += Heal_Point
					if Perso.Hp > Perso.MaxHp {
						Perso.Hp = Perso.MaxHp
					}
					fmt.Println("You have now", Perso.Hp, "Hp.")
					fmt.Println("-------------------------------------------------------------------------------------------------")

				}

			}
		} else {
			fmt.Println(Boss.Name, "attack you")
			Perso.Hp -= Boss.Str
			fmt.Println("You have now", Perso.Hp, "Hp.")
			fmt.Println("-------------------------------------------------------------------------------------------------")

		}
		counter++

		if Boss.Hp <= 0 {
			fmt.Println("You have fought the Bosss")
			fmt.Println("Do you want to retry ?(Y/n)")
			fmt.Scanln(&Choice)
			fmt.Println("-------------------------------------------------------------------------------------------------")

			if strings.ToUpper(Choice) == "Y" {
				Floor = 1
				counter = 0
			} else if strings.ToLower(Choice) == "n" {
				os.Exit(3)
			}
		}
		if Perso.Hp <= 0 {
			fmt.Println("You die")
			os.Exit(3)

		}
	}
}
