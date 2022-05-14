package service

import (
	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
	databaseService "github.com/akromibn37/hospitalityCollaboration/dao/database/service"
	apimodelService "github.com/akromibn37/hospitalityCollaboration/models/service"
	"github.com/akromibn37/hospitalityCollaboration/pkg/logging"
	projutil "github.com/akromibn37/hospitalityCollaboration/util"
)

type Service struct {
}

func (s *Service) ServiceGetAll(t *logging.Timelog) (res *apimodelService.ServiceGetAllResponse, err error) {

	t.TimeInDb = time.Now()
	data, err := databaseService.GetServiceAll()
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()
	resData := make([]apimodelService.ServiceDetail, len(data))
	for i := 0; i < len(data); i++ {
		resData[i].ServiceId = data[i].ServiceId
		resData[i].ServiceName = data[i].ServiceName
		resData[i].ServiceDescription = data[i].ServiceDescription
	}
	res = &apimodelService.ServiceGetAllResponse{Detail: resData}

	return res, nil
}

func (s *Service) ServiceCreate(req *apimodelService.ServiceCreateRequest, t *logging.Timelog) (res *apimodelService.ServiceCreateResponse, err error) {

	serviceId := projutil.GenerateID(constant.SERVICE_ID)
	t.TimeInDb = time.Now()
	err = databaseService.CreateService(req, serviceId)
	if err != nil {
		return nil, err
	}
	t.TimeOutDb = time.Now()

	res = &apimodelService.ServiceCreateResponse{}

	return res, nil
}

func (s *Service) ServiceUpdate(req *apimodelService.ServiceUpdateRequest, t *logging.Timelog) (res *apimodelService.ServiceUpdateResponse, err error) {
	t.TimeInDb = time.Now()
	if err != nil {
		return nil, err
	}
	err = databaseService.UpdateService(req)

	t.TimeOutDb = time.Now()
	res = &apimodelService.ServiceUpdateResponse{}

	return res, nil
}

func (s *Service) ServiceDelete(req *apimodelService.ServiceDeleteRequest, t *logging.Timelog) (res *apimodelService.ServiceDeleteResponse, err error) {
	t.TimeInDb = time.Now()
	if err != nil {
		return nil, err
	}
	err = databaseService.DeleteService(req)

	t.TimeOutDb = time.Now()
	res = &apimodelService.ServiceDeleteResponse{}

	return res, nil
}
