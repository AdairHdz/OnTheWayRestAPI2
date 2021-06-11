package mappers

import (
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)

func CreateServiceRequestEntity(serviceRequestDTO dataTransferObjects.ReceivedServiceRequestDTO) businessEntities.ServiceRequest {
	response := businessEntities.ServiceRequest{
		ID:                 uuid.NewV4(),
		Cost:               serviceRequestDTO.Cost,
		Date:               time.Now(),
		AddressID:          serviceRequestDTO.DeliveryAddressID,
		Description:        serviceRequestDTO.Description,
		KindOfService:      serviceRequestDTO.KindOfService,
		ServiceStatus:      businessEntities.PendingOfAcceptance,
		ServiceRequesterID: serviceRequestDTO.ServiceRequesterID,
		ServiceProviderID:  serviceRequestDTO.ServiceProviderID,
	}

	return response
}

func CreateServiceRequestDTOAsResponse(serviceRequest businessEntities.ServiceRequest) dataTransferObjects.ResponseServiceRequestDTO {
	formattedDate := serviceRequest.Date.Format("2006-01-02")
	response := dataTransferObjects.ResponseServiceRequestDTO{
		ID:                 serviceRequest.ID,
		Date:               formattedDate,
		Status:             serviceRequest.ServiceStatus,
		Cost:               serviceRequest.Cost,
		DeliveryAddressID:  serviceRequest.AddressID,
		Description:        serviceRequest.Description,
		KindOfService:      serviceRequest.KindOfService,
		ServiceProviderID:  serviceRequest.ServiceProviderID,
		ServiceRequesterID: serviceRequest.ServiceRequesterID,
	}

	return response
}

func CreateServiceRequestDTOWithDetailsAsResponse(serviceRequest businessEntities.ServiceRequest) dataTransferObjects.ResponseServiceRequestDTOWithDetails {
	formattedDate := serviceRequest.Date.Format("2006-01-02")
	response := dataTransferObjects.ResponseServiceRequestDTOWithDetails{
		ID:               serviceRequest.ID,
		Date:             formattedDate,
		Status:           serviceRequest.ServiceStatus,
		Cost:             serviceRequest.Cost,
		DeliveryAddress:  CreateAddressDTOWithCityAsResponse(serviceRequest.DeliveryAddress),
		Description:      serviceRequest.Description,
		KindOfService:    serviceRequest.KindOfService,
		ServiceProvider:  CreateUserDTOWithNameOnlyAsResponse(serviceRequest.ServiceProvider.User, serviceRequest.ServiceProvider.ID),
		ServiceRequester: CreateUserDTOWithNameOnlyAsResponse(serviceRequest.ServiceRequester.User, serviceRequest.ServiceRequester.ID),
	}
	return response
}
