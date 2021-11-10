package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonmp/bss-company-api/internal/repo"

	pb "github.com/ozonmp/bss-company-api/pkg/bss-company-api"
)

var (
	totalCompanyNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_company_api_company_not_found_total",
		Help: "Total number of companys that were not found",
	})
)

type CompanyAPI struct {
	pb.UnimplementedBssCompanyApiServiceServer
	repo repo.Repo
}

// NewCompanyAPI returns api of bss-company-api service
func NewCompanyAPI(r repo.Repo) pb.BssCompanyApiServiceServer {
	return &CompanyAPI{repo: r}
}
