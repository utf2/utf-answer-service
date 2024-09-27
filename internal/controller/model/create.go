package transfer

import "github.com/google/uuid"

type StudentAnswerCreateRequest struct {
	StudentID      uuid.UUID
	QuestionID     uuid.UUID
	ChosenAnswerID uuid.UUID
}

type StudentAnswerCreateResponse struct {
	Success         bool
	StudentAnswerID uuid.UUID
}
