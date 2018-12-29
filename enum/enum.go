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

func (lb LobbyType) String() (s string) {
	switch lb {
	case LobbyTypeInvalid:
		s = "Invalid"
	case LobbyTypePublicMatchmaking:
		s = "Public matchmaking"
	case LobbyTypePractise:
		s = "Practise"
	case LobbyTypeTournament:
		s = "Tournament"
	case LobbyTypeTutorial:
		s = "Tutorial"
	case LobbyTypeCoopWithBots:
		s = "Co-op with bots."
	case LobbyTypeTeamMatch:
		s = "Team match"
	case LobbyTypeSoloQueue:
		s = "Solo Queue"
	case LobbyTypeRanked:
		s = "Ranked"
	case LobbyTypeSolo:
		s = "1v1 Mid"
	default:
		s = "not support lobbytype"
	}

	return s
}

func (mod D2GameMode) String() (s string) {
	switch mod {
	case None:
		s = "None"
	case AllPick:
		s = "All pick"
	case CaptainMode:
		s = "Captain's Mode"
	case RandomDraft:
		s = "Random Draft"
	case SingleDraft:
		s = "Single Draft"
	case AllRandom:
		s = "All Random"
	case Intro:
		s = "Intro"
	case Diretide:
		s = "Diretide"
	case ReverseCaptainMode:
		s = "Reverse Captain Mode"
	case TheGreeviling:
		s = "The Greeviling"
	case Tutorial:
		s = "Tutorial"
	case MidOnly:
		s = "Mid Only"
	case LeastPlayed:
		s = "Least Played"
	case NewPlayerPool:
		s = "New Player Pool"
	case CompendiumMatchmaking:
		s = "Compendium Matchmaking"
	case CoopVSBots:
		s = "Co-op vs Bots"
	case CaptainsDraft:
		s = "Captains Draft"
	case AllRandomDeathmatch:
		s = "AllRandom Deathmatch"
	case MidOnly1V1:
		s = "1v1 Mid Only"
	case RankedMatchmaking:
		s = "Ranked Matchmaking"
	default:
		s = "not support D2GameMode"
	}

	return s
}