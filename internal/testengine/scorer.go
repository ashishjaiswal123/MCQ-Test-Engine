package testengine

import (
	"fmt"
	"test-engine/internal/model"
)

// ValidateAnswer validates the user's answer against the correct answer
func (engine *TestEngine) ValidateAnswer(question *model.Question) {
	if question.UserAttempt == question.CorrectAns {
		engine.UserScore++
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Incorrect! Correct answer: %s\n\n", question.CorrectAns)
	}
}

// CalculateFinalScore calculates the final score
func (engine *TestEngine) CalculateFinalScore() int {
	return engine.UserScore
}
