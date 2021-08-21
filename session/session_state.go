package session

type SessionState int

const (
	NewSession SessionState = iota
	NewCharacter
	NewPassword
	RepeatPassword
	SelectRace
	ExistingCharacter
	ExistingCharacterPassword
	Playing
)
