package model

// Question struct to represent each question
type Question struct {
	Text        string   `json:"question"`
	Options     []string `json:"choices"`
	CorrectAns  string   `json:"correctAnswer"`
	UserAttempt string   `json:"userAttempt,omitempty"`
}

// Test struct to represent a test
type Test struct {
	Questions []Question `json:"questions"`
}
