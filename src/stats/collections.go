package stats

import (
	"skycrypt/src/api"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
)

type CollectionsOutput struct {
	Categories       map[string]CollectionCategory `json:"categories"`
	TotalCollections int                           `json:"totalCollections"`
	MaxedCollections int                           `json:"maxedCollections"`
}

type CollectionCategory struct {
	Name       string                   `json:"name"`
	Texture    string                   `json:"texture"`
	Items      []CollectionCategoryItem `json:"items"`
	TotalTiers int                      `json:"totalTiers"`
	MaxedTiers int                      `json:"maxedTiers"`
}

type CollectionCategoryItem struct {
	Name        string                         `json:"name"`
	Id          string                         `json:"id"`
	Texture     string                         `json:"texture"`
	Amount      int                            `json:"amount"`
	TotalAmount int                            `json:"totalAmount"`
	Tier        int                            `json:"tier"`
	MaxTier     int                            `json:"maxTier"`
	Amounts     []CollectionCategoryItemAmount `json:"amounts"`
}

type CollectionCategoryItemAmount struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

type BossCollectionsFloorData struct {
	floorId int
	item    CollectionCategoryItem
}

func getBossCollections(userProfile *models.Member) CollectionCategory {
	bossCollections := []CollectionCategoryItem{}

	dungeons := GetFloorCompletions(userProfile)
	var floorItems []BossCollectionsFloorData

	for floor, amount := range dungeons.Total {
		if floor == "total" {
			continue
		}

		index, _ := utility.ParseInt(floor)
		boss := constants.BOSS_COLLECTIONS[index-1]
		tier := 0
		for _, t := range boss.Collections {
			if t <= amount {
				tier++
			}
		}

		item := CollectionCategoryItem{
			Name:        boss.Name,
			Id:          floor,
			Texture:     boss.Texture,
			Amount:      amount,
			TotalAmount: amount,
			Tier:        tier,
			MaxTier:     len(boss.Collections),
			Amounts: []CollectionCategoryItemAmount{
				{
					Username: "Normal",
					Amount:   dungeons.Normal[floor],
				},
				{
					Username: "Master Mode",
					Amount:   dungeons.Master[floor],
				},
			},
		}
		floorId, _ := utility.ParseInt(floor)
		floorItems = append(floorItems, BossCollectionsFloorData{
			floorId: floorId,
			item:    item,
		})

	}

	kuudraTier := 0
	KUUDRA_CONSTANTS := constants.BOSS_COLLECTIONS[7]
	kuudraCompletions := GetKuudraCompletions(userProfile)
	for _, t := range KUUDRA_CONSTANTS.Collections {
		if t <= kuudraCompletions {
			kuudraTier++
		}
	}

	kuudraAmounts := []CollectionCategoryItemAmount{}
	for tier := range constants.KUUDRA_COMPLETIONS_MULTIPLIER {
		amounts := userProfile.CrimsonIsle.Kuudra
		if amounts == nil {
			amounts = &map[string]int{}
		}

		if tier == "none" {
			tier = "basic"
		}

		kuudraAmounts = append(kuudraAmounts, CollectionCategoryItemAmount{
			Username: utility.TitleCase(tier),
			Amount:   (*amounts)[tier],
		})
	}

	floorItems = append(floorItems, BossCollectionsFloorData{
		floorId: 8,
		item: CollectionCategoryItem{
			Name:        KUUDRA_CONSTANTS.Name,
			Id:          "kuudra",
			Texture:     KUUDRA_CONSTANTS.Texture,
			Amount:      kuudraCompletions,
			TotalAmount: kuudraCompletions,
			Tier:        kuudraTier,
			MaxTier:     len(KUUDRA_CONSTANTS.Collections),
			Amounts:     kuudraAmounts,
		},
	})

	utility.SortSlice(floorItems, func(i, j int) bool {
		return floorItems[i].floorId < floorItems[j].floorId
	})

	maxedTiers := 0
	for _, fd := range floorItems {
		bossCollections = append(bossCollections, fd.item)
		if fd.item.MaxTier == fd.item.Tier {
			maxedTiers++
		}
	}

	return CollectionCategory{
		Name:       "Boss",
		Texture:    "/api/item/SKULL_ITEM:1",
		Items:      bossCollections,
		TotalTiers: len(bossCollections),
		MaxedTiers: maxedTiers,
	}
}

func GetCollections(userProfile *models.Member, profile *models.Profile) CollectionsOutput {
	usernames := map[string]string{}
	for memberId := range profile.Members {
		username, err := api.GetUsername(memberId)
		if err != nil {
			username = "Unknown"
		}

		usernames[memberId] = username
	}

	output := CollectionsOutput{
		Categories: map[string]CollectionCategory{},
	}

	output.Categories["BOSSES"] = getBossCollections(userProfile)
	for categoryId, categoryData := range constants.COLLECTIONS {
		category := CollectionCategory{
			Name:    categoryData.Name,
			Texture: categoryData.Texture,
			Items:   []CollectionCategoryItem{},
		}

		for _, itemData := range categoryData.Collections {
			amount := (*userProfile.Collections)[itemData.Id]

			totalAmount, amounts := 0, []CollectionCategoryItemAmount{}
			for memberId, memberData := range profile.Members {
				memberAmount := (*memberData.Collections)[itemData.Id]
				totalAmount += memberAmount
				if memberAmount > 0 {
					amounts = append(amounts, CollectionCategoryItemAmount{
						Username: usernames[memberId],
						Amount:   memberAmount,
					})
				}
			}

			tier := 0
			for _, tierData := range itemData.Tiers {
				if amount >= tierData.AmountRequired {
					tier = tierData.Tier
				}
			}

			item := CollectionCategoryItem{
				Name:        itemData.Name,
				Id:          itemData.Id,
				Texture:     itemData.Texture,
				Amount:      amount,
				TotalAmount: totalAmount,
				Tier:        tier,
				MaxTier:     itemData.MaxTier,
				Amounts:     amounts,
			}

			category.Items = append(category.Items, item)
			category.TotalTiers++
			if item.MaxTier == item.Tier {
				category.MaxedTiers++
			}
		}

		output.TotalCollections += len(category.Items)
		output.MaxedCollections += category.MaxedTiers
		output.Categories[categoryId] = category
	}

	return output
}
