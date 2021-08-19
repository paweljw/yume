package network

type ConnectionState int

const (
	NewConnection ConnectionState = iota
	NewCharacter
	NewPassword
	RepeatPassword
	SelectRace
	ExistingCharacter
	ExistingCharacterPassword
	Playing
)

var NewCharacterStates = []ConnectionState{
	NewCharacter,
	NewPassword,
	RepeatPassword,
	SelectRace,
}

func isNewCharacterState(c ConnectionState) bool {
	for _, state := range NewCharacterStates {
		if state == c {
			return true
		}
	}

	return false
}
