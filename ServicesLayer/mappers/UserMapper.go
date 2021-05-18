package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	uuid "github.com/satori/go.uuid"
)


func CreateUserDTOAsResponse(user businessEntities.User, CustomID uuid.UUID) dataTransferObjects.ResponseUserDTO {
	response := dataTransferObjects.ResponseUserDTO{
		ID: CustomID,
		Names: user.Names,
		LastName: user.LastName,
		EmailAddress: user.EmailAddress,
		UserType: user.UserType,
		Verified: user.Verified,
		StateID: user.StateID,
	}
	
	return response
}

func CreateUserEntity(receivedUserDTO dataTransferObjects.ReceivedUserDTO) (businessEntities.User, error) {

	hashedPassword, hashingError := hashing.GenerateHash(receivedUserDTO.Password)

	if hashingError != nil {
		return businessEntities.User{}, hashingError
	}

	userEntity := businessEntities.User{
		ID: uuid.NewV4(),
		Names: receivedUserDTO.Names,
		LastName: receivedUserDTO.LastName,
		EmailAddress: receivedUserDTO.EmailAddress,
		Password: hashedPassword,
		UserType: receivedUserDTO.UserType,
		Verified: false,
		StateID: receivedUserDTO.StateID,
	}

	return userEntity, nil
}

func CreateUserDTOWithNameOnlyAsResponse(user businessEntities.User, CustomID uuid.UUID) dataTransferObjects.ResponseUserDTOWithNamesOnly {
	response := dataTransferObjects.ResponseUserDTOWithNamesOnly {
		ID: CustomID,
		Names: user.Names,
		LastName: user.LastName,
	}
	
	return response
}