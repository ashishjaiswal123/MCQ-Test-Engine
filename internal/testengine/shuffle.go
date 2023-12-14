package testengine

import (
	"math/rand"
)

// ShuffleQuestions shuffles the order of questions for each test attempt
func (engine *TestEngine) ShuffleQuestions() {
	rand.Shuffle(len(engine.Test.Questions), func(i, j int) {
		engine.Test.Questions[i], engine.Test.Questions[j] = engine.Test.Questions[j], engine.Test.Questions[i]
	})
}
