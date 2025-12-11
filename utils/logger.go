package utils

import (
	"log"
)

// LogUserAction logs user activities asynchronously.
// In a real high-load system, this would push to a queue or external system.
func LogUserAction(action string, userID int) {
	// Simple standard output logging
	log.Printf("AUDIT LOG: ActionType=%s, UserID=%d\n", action, userID)
}
