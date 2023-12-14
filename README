# Multiple Choice Test Engine

This Go application implements a simple multiple-choice test engine with features such as loading tests from files, displaying questions, prompting the user for answers, validating answers, scoring, and more.

## Functionality

### 1. Ability to Load a Test from a File

The application allows loading a test from a JSON file using the `LoadTestFromFile` function in `testengine/loader.go`.

### 2. Question Structure

Each question is represented by a `Question` struct in `model/question.go` with attributes such as Text, Options, and CorrectAns.

### 3. Test Engine

The test engine is implemented in `testengine/engine.go` and includes functions to run the test, display questions, prompt users, and validate answers.

### 4. Scoring

Scoring is handled by the `TestEngine` struct in `testengine/engine.go`, which keeps track of the user's score.

### 5. Additional Features

- Shuffling questions for each test attempt.
- Timer for the test duration.
- Storing user details (name, email, etc.) and scores in a PostgreSQL database.

## How to Run

1. Install dependencies:

   ```bash
   go mod download


2. Run the application:
     go run .\cmd\

3. Access the API endpoint:
    curl -X POST http://localhost:8080/start-test -d '{"name": "John", "email": "john@example.com"}'
    Then in the terminal, where Go application running there you can start answering questions
