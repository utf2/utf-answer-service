package controller

import transfer "github.com/utf2/utf-answer-service/internal/controller/model"

type AnswerControllerAPI interface {
	Create(transfer.StudentAnswerCreateRequest) transfer.StudentAnswerCreateResponse
	SearchStudentFormAnswers(transfer.SearchStudentAnswersRequest) transfer.SearchStudentAnswersResponse
}
