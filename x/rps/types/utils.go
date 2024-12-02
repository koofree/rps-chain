package types

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

// the provided string is a valid SHA256 hash
func isValidHash(s string) bool {
	pattern := "^[a-fA-F0-9]{64}$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

// isMoveRevealed is a healer function to compare the commitment
// and the hash of the revealed move and the salt
func isMoveRevealed(commitment, move, salt string) bool {
	hash := CalculateHash(move, salt)
	return hash == commitment
}

// CalculateHash is a helper function to calculate the hash of a move and a salt
func CalculateHash(move, salt string) string {
	// concatenate move and salt
	combined := []byte(move + salt)

	// calculate the hash
	hash := sha256.Sum256(combined)

	// return the hash as a hex string
	return hex.EncodeToString(hash[:])
}
