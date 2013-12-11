package riotapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type api struct {
	platform string
	apikey   string
}

func NewAPI(platform, apikey string) *api {
	a := new(api)
	a.platform = platform
	a.apikey = apikey

	return a
}

func (a *api) getBody(baseurl string, hasparam bool) ([]byte, error) {
	var url string
	if hasparam {
		url = "http://prod.api.pvp.net/api/lol/" + a.platform + baseurl + "&api_key=" + a.apikey
	} else {
		url = "http://prod.api.pvp.net/api/lol/" + a.platform + baseurl + "?api_key=" + a.apikey
	}

	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

//Champion
func (a *api) GetChampions(f2ponly bool) ([]ChampionDTO, error) {
	var f2p string

	if f2ponly {
		f2p = "true"
	} else {
		f2p = "false"
	}

	body, err := a.getBody("/v1.1/champion?freeToPlay="+f2p, true)

	if err != nil {
		return nil, err
	}

	type championlist struct {
		Champions []ChampionDTO
	}

	var res championlist
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res.Champions, nil
}

//Game
func (a *api) GetRecentGames(summonerid int64) ([]GameDTO, error) {
	body, err := a.getBody(fmt.Sprintf("/v1.1/game/by-summoner/%d/recent", summonerid), false)

	if err != nil {
		return nil, err
	}

	type gamelist struct {
		Games      []GameDTO
		Summonerid int64
	}

	var res gamelist
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res.Games, nil
}

//League
func (a *api) GetLeagues(summonerid int64) (map[string]LeagueDTO, error) {
	body, err := a.getBody(fmt.Sprintf("/v1.1/league/by-summoner/%d", summonerid), false)

	if err != nil {
		return nil, err
	}

	var res map[string]LeagueDTO
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

//Stats
func (a *api) GetAgStats(summonerid int64, season string) ([]AggregatedStatDTO, error) {
	body, err := a.getBody(fmt.Sprintf("/v1.1/stats/by-summoner/%d/summary?season="+season, summonerid), true)

	if err != nil {
		return nil, err
	}

	type statlist struct {
		PlayerStatSummaries []AggregatedStatDTO
		SummonerId          int64
	}

	var res statlist
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res.PlayerStatSummaries, nil
}

func (a *api) GetRankedStats(summonerid int64, season string) (RankedStatsDTO, error) {
	var res RankedStatsDTO
	body, err := a.getBody(fmt.Sprintf("/v1.1/stats/by-summoner/%d/ranked?season="+season, summonerid), true)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (a *api) GetMasteries(summonerid int64) ([]MasteryPageDTO, error) {
	body, err := a.getBody(fmt.Sprintf("/v1.1/summoner/%d/masteries", summonerid), false)

	if err != nil {
		return nil, err
	}

	type masterylist struct {
		Pages      []MasteryPageDTO
		SummonerId int64
	}

	var res masterylist
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res.Pages, nil
}

func (a *api) GetRunes(summonerid int64) ([]RunePageDTO, error) {
	body, err := a.getBody(fmt.Sprintf("/v1.1/summoner/%d/runes", summonerid), false)

	if err != nil {
		return nil, err
	}

	type runepages struct {
		Pages      []RunePageDTO
		SummonerId int64
	}

	var res runepages
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res.Pages, nil
}

func (a *api) GetSummonerByName(name string) (SummonerDTO, error) {
	var res SummonerDTO
	body, err := a.getBody("/v1.1/summoner/by-name/"+name, false)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (a *api) GetSummonerByID(summonerid int64) (SummonerDTO, error) {
	var res SummonerDTO
	body, err := a.getBody(fmt.Sprintf("/v1.1/summoner/%d", summonerid), false)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

//Get Summoners
//Get Teams
