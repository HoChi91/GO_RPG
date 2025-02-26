package main

import (
	parscejson "Hyrule_Castle/mods/ParsceJson"
	rarity "Hyrule_Castle/mods/Rarity"
	mathfunc "Hyrule_Castle/mods/gamefunc"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	Perso := parscejson.ParsePlayer(rarity.Player()) // choose a random player

	Boss := parscejson.ParseBoss(rarity.Boss()) // choose a random Boss

	var enemiesList []parscejson.Enemy //create a list of random enemies
	for i := 0; i < 9; i++ {
		Enemies := parscejson.ParseEnemies(rarity.Enemies())
		enemiesList = append(enemiesList, Enemies)
	}

	var Choice string
	var counter int

	var Reset = "\033[0m"    // for the color in the text
	var Red = "\033[31m"     //
	var Yellow = "\033[33m"  //
	var Magenta = "\033[35m" //
	var Cyan = "\033[36m"    //

	Pv_Max := Perso.Hp
	Heal_Point := Perso.Hp / 2

	fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")      // Introduction of the game
	fmt.Println("| WELCOME TO HYRULE CASTLE !")                                                                                     //
	fmt.Println("| Your adventure begins here. You are a brave warrior named", Perso.Name, "called to rescue Princess Zelda.")      //
	fmt.Println("| An ancient curse has taken hold of the kingdom, and only you can defeat the evil sorcerer.")                     //
	fmt.Println("| You must face 9 guardians, each protecting a floor of the castle, before confronting the final boss", Boss.Name) //
	fmt.Println("| Your journey begins now...")                                                                                     //
	fmt.Println("-------------------------------------------------------------------------------------------------")                //
	fmt.Println("" + Reset)

	Floor := 1

	for Floor <= 9 {
		time.Sleep(1 * time.Second)
		Enemies := enemiesList[Floor-1]
		fmt.Println(Cyan+"You are at floor", Floor)

		//first floor
		if Floor == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println(Cyan + "You step into the castle and feel a dark presence. An ominous guardian appears.")
			fmt.Println("An", Enemies.Name, "is in front of you.")
			fmt.Println(Red+"You have ", Perso.Hp, "Hp.")

		}

		counter = 0

		for {
			if mathfunc.IsEven(counter) { //turn of the player
				time.Sleep(1 * time.Second)
				fmt.Println(Cyan + "What are you doing ?")
				fmt.Println("1. Attack	2. Heal" + Cyan)
				fmt.Scanln(&Choice)
				fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
				if Choice == "1" { //attack option
					time.Sleep(1 * time.Second)
					fmt.Println(Cyan+"You attack", Enemies.Name)
					Enemies.Hp -= Perso.Str
					if Enemies.Hp < 0 {
						Enemies.Hp = 0
					}
					fmt.Println(Red+"The enemies have now", Enemies.Hp, "Hp")
					fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
				} else if Choice == "2" { //heal option
					if Perso.Hp == Pv_Max {
						time.Sleep(1 * time.Second)
						fmt.Println(Red + "You have the maximum of Hp")
						fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
					} else {
						time.Sleep(1 * time.Second)
						fmt.Println(Cyan + "You heal yourself")
						Perso.Hp += Heal_Point
						if Perso.Hp > Pv_Max { //check the Hp dont go over the MaxHp
							Perso.Hp = Pv_Max
						}
						fmt.Println(Red+"You have now", Perso.Hp, "Hp.")
						fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
					}
				} else {
					fmt.Println(Cyan + "You dont do anything")
					fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
				}

			} else { // turn of the enemies
				time.Sleep(1 * time.Second)
				fmt.Println(Cyan+Enemies.Name, "attacks you")
				Perso.Hp -= Enemies.Str
				fmt.Println(Red+"You have now", Perso.Hp, "Hp.")
				fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
			}
			counter++

			if Enemies.Hp <= 0 { // if enemy die
				time.Sleep(1 * time.Second)
				fmt.Println(Cyan + "You have fought the enemies")
				fmt.Println(Red+"You have now", Perso.Hp, "Hp.")
				fmt.Println(Cyan + "You can now pass to the next floor. Do you want to continue? (Y/n)") // choose if continue or not
				fmt.Scanln(&Choice)
				fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
			
				switch strings.ToUpper(Choice) {
				case "Y": // if the player continue
					Floor++
					time.Sleep(1 * time.Second)
					fmt.Println(Cyan + "You ascend to the next floor, ready to face the next guardian. Beware of the dangers ahead!")
				case "N": // if player stop
					os.Exit(3)
				default: // if invalid input
					fmt.Println(Cyan + "You did not enter a valid choice")
					continue
				}
			}
			
			if Perso.Hp <= 0 { // death of player
				fmt.Println(Red + "You die")
				os.Exit(3)
			}
		}
	}

	//Final boss floor
	time.Sleep(1 * time.Second)
	fmt.Println(Yellow + "Welcome to the FINAL BOSS floor !")
	fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")

	time.Sleep(1 * time.Second)
	fmt.Println(Cyan + "As you enter the final chamber, the BOSS stands before you, radiating dark energy.")
	fmt.Println(Cyan + "You know this is the ultimate battle for the fate of Hyrule.")

	counter = 0
	for {
		if mathfunc.IsEven(counter) {
			time.Sleep(1 * time.Second)
			fmt.Println("1. Attack	2. Heal")
			fmt.Scanln(&Choice)
			fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")

			if Choice == "1" { //attack option
				time.Sleep(1 * time.Second)
				fmt.Println(Cyan+"You attack", Boss.Name)
				Boss.Hp -= Perso.Str
				fmt.Println(Red+"The BOSS has now", Boss.Hp, "Hp")
				fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
			} else if Choice == "2" { //heal option
				if Perso.Hp == Pv_Max {
					time.Sleep(1 * time.Second)
					fmt.Println(Red + "You have the maximum of Hp")
					fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
				} else {
					time.Sleep(1 * time.Second)
					fmt.Println(Cyan + "You heal yourself")
					Perso.Hp += Heal_Point
					if Perso.Hp > Pv_Max {
						Perso.Hp = Pv_Max
					}
					fmt.Println(Red+"You have now", Perso.Hp, "Hp.")
					fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
				}
			}
		} else {
			fmt.Println(Cyan+Boss.Name, "attacks you")
			Perso.Hp -= Boss.Str
			fmt.Println(Red+"You have now", Perso.Hp, "Hp.")
			fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")
		}
		counter++

		if Boss.Hp <= 0 { //Win
			fmt.Println(Yellow + "You have fought the FINAL BOSS")

			fmt.Println(Cyan + "Do you want to retry ?(Y/n)") //choose if retry or not
			fmt.Scanln(&Choice)
			fmt.Println(Magenta + "-------------------------------------------------------------------------------------------------")

			if strings.ToUpper(Choice) == "Y" {
				Floor = 1
				counter = 0
			} else if strings.ToLower(Choice) == "n" {
				os.Exit(3)
			}
		}
		if Perso.Hp <= 0 { //Loose
			fmt.Println(Red + "You die")
			os.Exit(3)
		}
	}
}
