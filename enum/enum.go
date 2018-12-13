package enum

type LobbyType int
const (
	LobbyTypeInvalid LobbyType = -1
	LobbyTypePublicMatchmaking LobbyType = iota
	LobbyTypePractise
	LobbyTypeTournament
	LobbyTypeTutorial
	LobbyTypeCoopWithBots		// Co-op with bots.
	LobbyTypeTeamMatch
	LobbyTypeSoloQueue
	LobbyTypeRanked
	LobbyTypeSolo				// 1v1 Mid
)

type D2GameMode int
const (
	None D2GameMode = iota
	AllPick
	CaptainMode
	RandomDraft
	SingleDraft
	AllRandom
	Intro
	Diretide
	ReverseCaptainMode
	TheGreeviling
	Tutorial
	MidOnly
	LeastPlayed
	NewPlayerPool
	CompendiumMatchmaking
	CoopVSBots				// Co-op vs Bots
	CaptainsDraft
	AbilityDraft
	AllRandomDeathmatch
	MidOnly1V1				// 1v1 Mid Only
	RankedMatchmaking
)

