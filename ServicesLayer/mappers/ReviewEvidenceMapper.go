package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)


func CreateSliceOfReviewEvidenceDTOAsResponse(reviewEvidences []businessEntities.ReviewEvidence) []dataTransferObjects.ReviewEvidenceDTO {
	var response []dataTransferObjects.ReviewEvidenceDTO

	for _, evidenceItem := range reviewEvidences {
		evidence := dataTransferObjects.ReviewEvidenceDTO {
			Name: evidenceItem.Name,
		}

		response = append(response, evidence)
	}

	return response
}

func CreateSliceOfReviewEvidenceEntities(reviewDTOs []dataTransferObjects.ReviewEvidenceDTO) []businessEntities.ReviewEvidence {
	var response []businessEntities.ReviewEvidence

	for _, evidenceItem := range reviewDTOs {
		evidence := businessEntities.ReviewEvidence {
			ID: uuid.NewV4(),
			Name: evidenceItem.Name,
		}
		response = append(response, evidence)
	}

	return response
}