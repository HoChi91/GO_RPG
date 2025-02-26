package rarity

import (
	"time"

	"math/rand"
)

func Boss() string {
	bossName := ""
	rarity_tier := 0

	rand.Seed(time.Now().UnixNano())
	randomPercentage := rand.Intn(100)
	Y := rand.Intn(2)
	if randomPercentage == 0 {
		randomPercentage = 1
	}
	if randomPercentage >= 50 {
		rarity_tier = 1
	} else if randomPercentage >= 30 {
		rarity_tier = 2
	} else if randomPercentage >= 15 {
		rarity_tier = 3
	} else if randomPercentage >= 4 {
		rarity_tier = 4
	} else if randomPercentage >= 1 {
		rarity_tier = 5
	}

	switch rarity_tier {
	case 1:
		bossName = "Ganon"
	case 2:
		bossName = "Gohma"
	case 3:
		if Y == 0 {
			bossName = "Odolwa"
		}
		if Y == 1 {
			bossName = "Gleeoks"
		}
	case 4:
		if Y == 0 {
			bossName = "Volvagia"
		}

		if Y == 1 {
			bossName = "Stallord"
		}
	case 5:
		bossName = "Onox"
	}
	return bossName
}

func Enemies() string {

	EnemieName := ""
	rarity_tier := 0

	rand.Seed(time.Now().UnixNano())
	randomPercentage := rand.Intn(100)

	if randomPercentage == 0 {
		randomPercentage = 1
	}
	if randomPercentage >= 50 {
		rarity_tier = 1
	} else if randomPercentage >= 30 {
		rarity_tier = 2
	} else if randomPercentage >= 15 {
		rarity_tier = 3
	} else if randomPercentage >= 4 {
		rarity_tier = 4
	} else if randomPercentage >= 1 {
		rarity_tier = 5
	}

	switch rarity_tier {
	case 1:
		Y := rand.Intn(4)
		if Y == 0 {
			EnemieName = "Skulltula"
		}
		if Y == 1 {
			EnemieName = "Octorock"
		}
		if Y == 2 {
			EnemieName = "Deku Scrub"
		}
		if Y == 3 {
			EnemieName = "Bokoblin"
		}
	case 2:
		Y := rand.Intn(2)
		if Y == 0 {
			EnemieName = "Lizalfos"
		}
		if Y == 1 {
			EnemieName = "Gibdo"
		}
	case 3:
		Y := rand.Intn(3)
		if Y == 0 {
			EnemieName = "Dead Hand"
		}
		if Y == 1 {
			EnemieName = "Gleeoks"
		}
		if Y == 2 {
			EnemieName = "Golden Skulltula"
		}
	case 4:
		Y := rand.Intn(2)
		if Y == 0 {
			EnemieName = "Stalfos"
		}

		if Y == 1 {
			EnemieName = "Darknuts"
		}
	case 5:
		Y := rand.Intn(2)
		if Y == 0 {
			EnemieName = "Guardian"
		}
		if Y == 1 {
			EnemieName = "Lynel"
		}
	}
	return EnemieName
}
func Player() string {
	PlayerName := ""
	rarity_tier := 0

	rand.Seed(time.Now().UnixNano())
	randomPercentage := rand.Intn(100)

	if randomPercentage == 0 {
		randomPercentage = 1
	}
	if randomPercentage >= 50 {
		rarity_tier = 1
	} else if randomPercentage >= 30 {
		rarity_tier = 2
	} else if randomPercentage >= 15 {
		rarity_tier = 3
	} else if randomPercentage >= 4 {
		rarity_tier = 4
	} else if randomPercentage >= 1 {
		rarity_tier = 5
	}

	switch rarity_tier {
	case 1:
		PlayerName = "Link"
	case 2:
		PlayerName = "Young_Link"
	case 3:
		PlayerName = "Sheik"
	case 4:
		PlayerName = "Impa"
	case 5:
		PlayerName = "Hylia"
	}

	return PlayerName
}
