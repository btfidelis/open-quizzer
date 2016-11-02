package validation

import (
	"github.com/btfidelis/quizzer/models"
	"errors"
	"log"
)

var validationErrors []error

type ValidationResult struct {
	Passed  bool
	Errors  []error
}

func ValidateQuiz(q models.Quiz) ValidationResult {
	validationErrors = make([]error, 0)

	if len(q.Answers) < 4 {
		addError("Questions must have at least 4 anwsers")
	}

	if len(q.Question) > 0 {
		addError("Question field is required")
	}

	if q.Type != models.QUIZ_MULTIPLE_ALTERNATIVE && q.Type != models.QUIZ_NORMAL {
		addError("Question type must be Multiple Anternative or Normal")
	}

	if len(q.CorrectAnswers) < 1 && q.Type == models.QUIZ_NORMAL {
		addError("On normal or multiple alternative questions, you need to especify at least one correct question")
	}
	log.Println(q)
	return ValidationResult{
		Passed: len(validationErrors) > 0,
		Errors: validationErrors,
	}

}

func addError(errorMessage string) {
	validationErrors = append(validationErrors, errors.New(errorMessage))
}