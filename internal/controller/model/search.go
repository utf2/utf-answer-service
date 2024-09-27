package transfer

import "github.com/google/uuid"

type SearchStudentAnswersRequest struct {
	StudentID uuid.UUID
	FormID    uuid.UUID
}

type SearchStudentAnswersResponse struct {
	Answers []AnswerDTO
}

type AnswerDTO struct {
	ID         uuid.UUID
	StudentID  uuid.UUID
	QuestionID uuid.UUID
	AnswerID   uuid.UUID
}
