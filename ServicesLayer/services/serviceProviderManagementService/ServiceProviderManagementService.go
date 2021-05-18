package serviceProviderManagementService

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/fileAnalyzer"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceProviderManagementService struct { }

func (ServiceProviderManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProvider := businessEntities.ServiceProvider{}
		searchError := serviceProvider.Find(serviceProviderID)		

		if searchError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}
		
		response := mappers.CreateServiceProviderDTOAsResponse(serviceProvider)

		context.JSON(http.StatusOK, response)
	}
}

func (ServiceProviderManagementService) FindMatches() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (ServiceProviderManagementService) Update() gin.HandlerFunc {
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

		updateError := serviceProvider.Update()

		if updateError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusOK)
	}
}

func (ServiceProviderManagementService) UpdateServiceProviderImage() gin.HandlerFunc{
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
			pathOfImageToBeDeleted := path + "/" + serviceProvider.BusinessPicture
			
			fmt.Println(pathOfImageToBeDeleted)
			os.Remove(pathOfImageToBeDeleted)
		}
		
		err = context.SaveUploadedFile(file, path + "/" + file.Filename)
		
		if err != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProvider.BusinessPicture = file.Filename

		databaseError := serviceProvider.Update()

		if databaseError != nil{
			context.AbortWithStatus(http.StatusConflict)
			return
		}
		
		context.Status(http.StatusOK)
	}
}
