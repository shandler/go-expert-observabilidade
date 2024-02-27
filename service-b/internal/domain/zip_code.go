package domain

import (
	"github.com/shandler/go-expert-observabilidade/service-b/internal/dto"
)

type ZipCode interface {
	Search(request dto.SearchRequest) dto.SearchResponse
}
