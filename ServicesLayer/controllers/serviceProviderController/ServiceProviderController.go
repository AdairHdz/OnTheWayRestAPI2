package serviceProviderController

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceProviderManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/fileAnalyzer"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var (
	serviceProviderMgtService = serviceProviderManagementService.ServiceProviderManagementService{}
)


func RegisterServiceProvider() gin.HandlerFunc{
	return func(context *gin.Context){
		receivedData := dataTransferObjects.ReceivedUserDTO{}
		context.BindJSON(&receivedData)

		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userEntity, mappingError := mappers.CreateUserEntity(receivedData, 1)

		if mappingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProviderEntity := businessEntities.ServiceProvider{
			ID: uuid.NewV4(),
			User: userEntity,
			AverageScore: 0,
			PriceRates: nil,
		}

		serviceProviderMgtService := serviceProviderManagementService.ServiceProviderManagementService{}
		registryError := serviceProviderMgtService.Register(serviceProviderEntity)

		if registryError != nil {
			context.AbortWithStatus(http.StatusConflict)
		}

		response := mappers.CreateUserDTOAsResponse(serviceProviderEntity.User, serviceProviderEntity.ID)
		
		context.JSON(http.StatusCreated, response)
	}
}

func FindMatches() gin.HandlerFunc{
	return func(context *gin.Context){		
		// maxPriceRate, parsingError := strconv.ParseFloat(context.Param("maxPriceRate"), 32)

		// if parsingError != nil {
		// 	return
		// }

		// city := context.Param("city")
		// kindOfService, parsingError := strconv.ParseUint(context.Param("kindOfService"), 10, 8)

		// if parsingError != nil {
		// 	return
		// }

		// context.Params

	}
}

func Find() gin.HandlerFunc{
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProvider, searchError := serviceProviderMgtService.Find(serviceProviderID)

		if searchError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateUserDTOAsResponse(serviceProvider.User, serviceProvider.ID)
		context.JSON(http.StatusOK, response)
	}
}

func Update() gin.HandlerFunc{
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		receivedData := struct{
			Names string `json:"names"`
			LastName string `json:"lastName"`
			Password string `json:"password"`
		}{}

		bindingError := context.BindJSON(&receivedData)

		if bindingError != nil {
			return
		}

		serviceProvider := businessEntities.ServiceProvider{ }
		serviceProvider.Find(serviceProviderID)

		if receivedData.Names != ""{
			serviceProvider.User.Names = receivedData.Names
		}

		if receivedData.LastName != "" {
			serviceProvider.User.LastName = receivedData.LastName
		}
				
		if receivedData.Password != "" {
			hashedPassword, hashingError := hashing.GenerateHash(serviceProvider.User.Password)	
			if hashingError != nil {
				context.AbortWithStatus(http.StatusConflict)
				return
			}
	
			serviceProvider.User.Password = hashedPassword
		}				
		
		validator :=  validators.GetValidator()
		validationErrors := validator.Var(serviceProvider.User.Names, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceProvider.User.LastName, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceProvider.User.Password, "required,max=80")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		repository := repositories.Repository{}
		updateError := repository.Update(&serviceProvider.User)

		if updateError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusOK)
	}
}

func UpdateServiceProviderImage() gin.HandlerFunc {
	return func(context *gin.Context){
		providerID := context.Param("providerId")
		path := "./images/" + providerID
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			os.Mkdir(path, 777)
		}
		
		serviceProvider := businessEntities.ServiceProvider{}

		
		serviceProvider.Find(uuid.FromStringOrNil(providerID))		


		if serviceProvider.ID == uuid.Nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		dirIsEmpty, err := fileAnalyzer.DirIsEmpty(path)
	
		
		file, _ := context.FormFile("image")		
		fileExtension := filepath.Ext(file.Filename)

		if !fileAnalyzer.ImageHasValidFormat(fileExtension) {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		if !dirIsEmpty {
			pathOfImageToBeDeleted := path + serviceProvider.BusinessPicture
			os.Remove(pathOfImageToBeDeleted)
		}
		
		err = context.SaveUploadedFile(file, path + "/profile_picture" + fileExtension)
		
		if err != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProvider.BusinessPicture = file.Filename

		repo := repositories.Repository{}
		err = repo.Update(&serviceProvider)
		
		context.Status(http.StatusOK)
	}
}