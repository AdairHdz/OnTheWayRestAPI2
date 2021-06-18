package reviewManagementService

import "github.com/gin-gonic/gin"

type IReviewManagementService interface {
	Register() gin.HandlerFunc
	Find() gin.HandlerFunc
	UploadEvidence() gin.HandlerFunc
}
