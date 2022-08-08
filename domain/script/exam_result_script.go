package script

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"time"
	"zmed_exam_manager/infrastructure/dynamo"
	"zmed_exam_manager/infrastructure/s3"
	"zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
	"zmed_exam_manager/utils"
)

type errorExamKeyData struct {
	examKey string
	appErr  app_errors.AppError
}

var examKeyChannel chan errorExamKeyData

func HandleExamsResultProcessing(examResultProvider *s3.Repository,
	examsRepository *dynamo.Repository) {
	ctx := context.Background()
	pullExamError := 0
	examKeyChannel = make(chan errorExamKeyData)
	handleStuckExam(ctx, examResultProvider)

	for {
		objects, appErr := examResultProvider.PullS3CompletedExams(ctx)
		if appErr != nil {
			handlePullExamsError(&pullExamError, appErr)
		}
		for _, exam := range objects {
			patientId, examId, appErr := getExamResultInfo(exam)
			if appErr != nil {
				examKeyChannel <- errorExamKeyData{examKey: *exam.Key, appErr: appErr}
			}
			examResult := zmed_model.Exam{
				Id:         examId,
				PatientId:  patientId,
				Status:     "Finished",
				FinishedAt: time.Now().String(),
			}
			_, appErr = examsRepository.Persist(ctx, &examResult)
			if appErr != nil {
				examKeyChannel <- errorExamKeyData{examKey: *exam.Key, appErr: appErr}
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func getExamResultInfo(exams types.Object) (string, string, app_errors.AppError) {
	patientId, appErr := utils.GetPatientId(*exams.Key)
	if appErr != nil {
		return "", "", appErr
	}
	examId, appErr := utils.GetExamId(*exams.Key)
	if appErr != nil {
		return "", "", appErr
	}
	return patientId, examId, nil
}

func handlePullExamsError(pullExamError *int, appErr app_errors.AppError) {
	if *pullExamError >= 3 {
		fmt.Println("PullExamError", time.Now().String(), appErr)
		*pullExamError = 0
	} else {
		*pullExamError += 1
	}
}

func handleStuckExam(ctx context.Context, examResultProvider *s3.Repository) {
	for {
		select {
		case errorExamKeyData := <-examKeyChannel:
			appErr := examResultProvider.MoveExamToStuckFolder(ctx, &errorExamKeyData.examKey)
			if appErr != nil {
				fmt.Println("Error, ", time.Now().String(), ", err1: ",
					errorExamKeyData.appErr, ", err2: ", appErr)
			}
		}
	}
}
