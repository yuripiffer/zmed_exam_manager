package patient_provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
)

type provider struct {
	host           string
	getPatientPath string
}

func NewProvider() *provider {
	return &provider{
		host:           config.ENV.PatientManagerHost,
		getPatientPath: config.ENV.PatientManagerGetPath,
	}
}

func (p *provider) GetPatient(document string) (*zmed_model.Patient, app_errors.AppError) {
	url := fmt.Sprintf("%s/%s", p.host, p.getPatientPath)

	var payload = map[string]interface{}{
		"document": document,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, app_errors.NewInternalServerError("Get Patient Request Error", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJson))
	if err != nil {
		return nil, app_errors.NewInternalServerError("Get Patient Request Error", err)
	}

	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 8 * time.Second}
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, app_errors.NewInternalServerError("Get Patient Request Error", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, app_errors.NewInternalServerError("Get Patient Request Error", err)
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	var patient *zmed_model.Patient

	err = json.Unmarshal(responseBody, &patient)
	if err != nil {
		return nil, app_errors.NewInternalServerError("Get Patient Request Error", err)
	}
	return patient, nil
}
