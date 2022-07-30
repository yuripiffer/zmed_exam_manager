package model

type Exam struct {
	Id         string `json:"exam_id"`
	PatientId  string `json:"patient_id"`
	Status     string `json:"status"`
	ExamType   int    `json:"exam_type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	FinishedAt string `json:"finished_at"`
	IsRevoked  string `json:"is_revoked"`
}
