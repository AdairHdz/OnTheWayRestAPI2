package mappers

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)

func CreateSliceOfReviewEvidenceDTOAsResponse(reviewID string, reviewEvidences []businessEntities.ReviewEvidence) []dataTransferObjects.ReviewEvidenceRespondeDTO {
	var response []dataTransferObjects.ReviewEvidenceRespondeDTO

	for _, evidenceItem := range reviewEvidences {
		evidence := dataTransferObjects.ReviewEvidenceRespondeDTO{
			Link: fmt.Sprintf("reviews/%s/%s", reviewID, evidenceItem.Name),
			Name: evidenceItem.Name,
		}

		response = append(response, evidence)
	}

	return response
}

func CreateSliceOfReviewEvidenceEntities(reviewDTOs []dataTransferObjects.ReviewEvidenceDTO) []businessEntities.ReviewEvidence {
	var response []businessEntities.ReviewEvidence

	for _, evidenceItem := range reviewDTOs {
		escapedPath := strings.Replace(evidenceItem.Name, "\\", "/", -1)
		evidence := businessEntities.ReviewEvidence{
			ID:   uuid.NewV4(),
			Name: filepath.Base(escapedPath),
		}
		response = append(response, evidence)
	}

	return response
}
