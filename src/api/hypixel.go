package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/utility"
)

var HYPIXEL_API_KEY = os.Getenv("HYPIXEL_API_KEY")

func GetPlayer(uuid string) (*models.Player, error) {
	var rawReponse models.HypixelPlayerResponse
	var response models.Player

	if !utility.IsUUID(uuid) {
		respUUID, err := GetUUID(uuid)
		if err != nil {
			return &response, err
		}

		uuid = respUUID
	}

	cache, err := redis.Get(fmt.Sprintf(`player:%s`, uuid))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), &rawReponse)
		if err == nil {
			return &rawReponse.Player, nil
		}
	}

	resp, err := http.Get(fmt.Sprintf("https://api.hypixel.net/v2/player?key=%s&uuid=%s", HYPIXEL_API_KEY, uuid))

	if err != nil {
		return &response, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &response, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &rawReponse)
	if err != nil {
		return &rawReponse.Player, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`player:%s`, uuid), string(body), 24*60*60)
	return &rawReponse.Player, nil
}

func GetProfiles(uuid string) (*models.HypixelProfilesResponse, error) {
	var response models.HypixelProfilesResponse
	if !utility.IsUUID(uuid) {
		respUUID, err := GetUUID(uuid)
		if err != nil {
			return &response, err
		}

		uuid = respUUID
	}

	cache, err := redis.Get(fmt.Sprintf(`profiles:%s`, uuid))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), &response)
		if err == nil {
			return &response, nil
		}
	}

	resp, err := http.Get(fmt.Sprintf("https://api.hypixel.net/v2/skyblock/profiles?key=%s&uuid=%s", HYPIXEL_API_KEY, uuid))
	if err != nil {
		return &response, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &response, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return &response, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`profiles:%s`, uuid), string(body), 5*60) // Cache for 5 minutes
	return &response, nil
}

func GetProfile(uuid string, profileId ...string) (*models.Profile, error) {
	profiles, err := GetProfiles(uuid)
	if err != nil {
		return &models.Profile{}, err
	}

	// If no profileId provided, return the first profile or selected profile
	if len(profileId) == 0 || (len(profileId) == 1 && profileId[0] == "") {
		if len(profiles.Profiles) == 0 {
			return &models.Profile{}, fmt.Errorf("no profiles found for UUID %s", uuid)
		}

		for _, profile := range profiles.Profiles {
			if profile.Selected {
				return &profile, nil
			}
		}

		return &profiles.Profiles[0], nil
	}

	// If profileId is provided, search for it
	targetProfileId := profileId[0]
	for _, profile := range profiles.Profiles {
		if profile.ProfileID == targetProfileId || profile.CuteName == targetProfileId {
			return &profile, nil
		}
	}

	return &models.Profile{}, fmt.Errorf("profile with ID %s not found for UUID %s", targetProfileId, uuid)
}

func GetMuseum(profileId string) (*map[string]models.Museum, error) {
	var rawReponse models.HypixelMuseumResponse

	cache, err := redis.Get(fmt.Sprintf(`museum:%s`, profileId))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), &rawReponse)
		if err == nil {
			return &rawReponse.Members, nil
		}
	}

	resp, err := http.Get(fmt.Sprintf("https://api.hypixel.net/v2/skyblock/museum?key=%s&profile=%s", HYPIXEL_API_KEY, profileId))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &rawReponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`museum:%s`, profileId), string(body), 60*30) // Cache for 30 minutes
	return &rawReponse.Members, nil
}
