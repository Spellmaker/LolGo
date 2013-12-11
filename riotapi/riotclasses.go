package riotapi

import ()

type date string

type ChampionDTO struct {
	Active            bool
	AttackRank        int
	BotEnabled        bool
	BotMmEnabled      bool
	DefenseRank       int
	DifficultyRank    int
	FreeToPlay        bool
	ID                int64
	MagicRank         int
	Name              string
	RankedPlayEnabled bool
}

type GameDTO struct {
	ChampionId    int
	CreateDate    int64
	CreateDateStr string
	FellowPlayers []PlayerDTO
	GameID        int64
	GameMode      string
	GameType      string
	Invalid       bool
	Level         int
	MapID         int
	Spell1        int
	Spell2        int
	Statistics    []StatisticDTO
	SubType       string
	TeamID        int
}

type PlayerDTO struct {
	ChampionID int
	SummonerID int64
	TeamID     int
}

type StatisticDTO struct {
	ID    int
	Name  string
	Value int
}

type LeagueDTO struct {
	Entries   []LeagueItemDTO
	Name      string
	Queue     string
	Tier      string
	timestamp int64
}

type LeagueItemDTO struct {
	IsFreshBlood     bool
	IsHotStreak      bool
	IsInactive       bool
	IsVeteran        bool
	LastPlayed       int64
	LeagueName       string
	LeaguePoints     int
	Losses           int
	MiniSeries       MiniSeriesDTO
	PlayerOrTeamID   string
	PlayerOrTeamName string
	QueueType        string
	Rank             string
	Tier             string
	TimeUntilDecay   int64
	Wins             int
}

type MiniSeriesDTO struct {
	Losses               int
	Progress             []byte
	Target               int
	TimeLeftToPlayMillis int64
	Wins                 int
}

type PlayerStatsSummary struct {
	AggregatedStats       []AggregatedStatDTO
	Losses                int
	ModifyDate            int64
	ModifyDateStr         string
	PlayerStatSummaryType string
	Wins                  int
}

type AggregatedStatDTO struct {
	Count int
	ID    int
	Name  string
}

type RankedStatsDTO struct {
	Champions     []ChampionStatsDTO
	ModifyDate    int64
	ModifyDateStr int64
	SummonerID    int64
}

type ChampionStatsDTO struct {
	ID    int
	Name  string
	Stats []ChampionStatDTO
}

type ChampionStatDTO struct {
	Count int
	ID    int
	Name  string
	Value int
}

type MasteryPageDTO struct {
	Current bool
	Name    string
	Talents []TalentDTO
}

type TalentDTO struct {
	ID   int
	Name string
	Rank int
}

type RunePageDTO struct {
	Current bool
	ID      int64
	Name    string
	Slots   []RuneSlotDTO
}

type RuneSlotDTO struct {
	Rune       RuneDTO
	RuneSlotID int
}

type RuneDTO struct {
	Description string
	ID          int
	Name        string
	Tier        int
}

type SummonerDTO struct {
	ID              int64
	Name            string
	ProfileIconID   int
	RevisionDate    int64
	RevisionDateStr string
	SummonerLevel   int64
}

type SummonerNameDTO struct {
	ID   int64
	Name string
}

type TeamDTO struct {
	CreateDate                    date
	LastGameDate                  date
	LastJoinDate                  date
	LastJoinedRankedTeamQueueDate date
	MatchHistory                  []MatchHistorySummaryDTO
	MessageOfDate                 MessageOfDayDTO
	ModifyDate                    date
	Name                          string
	Roster                        RosterDTO
	SecondLastJoinDate            date
	Status                        string
	Tag                           string
	TeamID                        TeamIDDTO
	TeamStatSummary               TeamStatSummaryDTO
	ThirdLastJoinDate             date
	Timestamp                     int64
}

type MatchHistorySummaryDTO struct {
	Assists           int
	Date              int64
	Deaths            int
	GameID            int64
	GameMode          string
	Invalid           bool
	Kills             int
	MapID             int
	OpposingTeamKills int
	OpposingTeamName  string
	Win               bool
}

type MessageOfDayDTO struct {
	CreateDate int64
	Message    string
	Version    int
}

type RosterDTO struct {
	MemberList []TeamMemberInfoDTO
	OwnerID    int64
}

type TeamIDDTO struct {
	FullID string
}

type TeamStatSummaryDTO struct {
	TeamID          TeamIDDTO
	TeamStatDetails []TeamStatDetailDTO
}

type TeamMemberInfoDTO struct {
	InviteDate date
	JoinDate   date
	PlayerID   int64
	Status     string
}

type TeamStatDetailDTO struct {
	AverageGamesPlayed int
	Losses             int
	MaxRating          int
	Rating             int
	SeedRating         int
	TeamID             TeamIDDTO
	TeamStatType       string
	Wins               int
}
