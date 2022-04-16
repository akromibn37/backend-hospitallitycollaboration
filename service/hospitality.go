package service

import (
	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	databaseHospitality "github.com/akromibn37/hospitalityCollaboration/dao/database/hospitality"
	apimodelHospitality "github.com/akromibn37/hospitalityCollaboration/models/hospitality"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	projutil "github.com/akromibn37/hospitalityCollaboration/util"
)

type Hospitality struct {
}

func (h *Hospitality) HospitalityGetAll(t *logging.Timelog) (res *apimodelHospitality.HospitalityGetAllResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseHospitality.GetHospitalityAll()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelHospitality.HospitalityDetail, len(data))
	for i := 0; i < len(data); i++ {
		resData[i].HospitalityId = data[i].HospitalityId
		resData[i].HospitalityName = data[i].HospitalityName
		resData[i].HospitalityDescription = data[i].HospitalityDescription
		resData[i].HospitalityType = data[i].HospitalityType
		resData[i].HospitalityContactName = data[i].HospitalityContactName
		resData[i].HospitalityPhoneNumber = data[i].HospitalityPhoneNumber
	}
	res = &apimodelHospitality.HospitalityGetAllResponse{Detail: resData}

	return res, nil
}

func (h *Hospitality) HospitalityCreate(req *apimodelHospitality.HospitalityCreateRequest, t *logging.Timelog) (res *apimodelHospitality.HospitalityCreateResponse, err error) {

	hospitalityId := projutil.GenerateID(constant.HOSPITALITY_ID)
	t.TimeInDb = time.Now()
	err = databaseHospitality.CreateHospitality(req, hospitalityId)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()

	res = &apimodelHospitality.HospitalityCreateResponse{}

	return res, nil
}

func (h *Hospitality) HospitalityUpdate(req *apimodelHospitality.HospitalityUpdateRequest, t *logging.Timelog) (res *apimodelHospitality.HospitalityUpdateResponse, err error) {
	t.TimeInDb = time.Now()
	if err != nil {
		return nil, err
	}
	err = databaseHospitality.UpdateHospitality(req)

	t.TimeOutDb = time.Now()
	res = &apimodelHospitality.HospitalityUpdateResponse{}

	return res, nil
}
