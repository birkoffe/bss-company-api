package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (o *CompanyAPI) DescribeCompanyV1(
	ctx context.Context,
	req *pb.DescribeCompanyV1Request,
) (*pb.DescribeCompanyV1Response, error) {
	log.Debug().Msg("DescribeCompanyV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeCompanyV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	company, err := o.repo.DescribeCompany(ctx, req.CompanyId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeCompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if company == nil {
		log.Debug().Uint64("companyId", req.CompanyId).Msg("company not found")
		totalCompanyNotFound.Inc()

		return nil, status.Error(codes.NotFound, "company not found")
	}

	log.Debug().Msg("DescribeCompanyV1 - success")

	return &pb.DescribeCompanyV1Response{
		Value: &pb.Company{
			Id:      company.ID,
			Name:    company.Name,
			Address: company.Address,
		},
	}, nil
}

func (o *CompanyAPI) CreateCompanyV1(
	ctx context.Context,
	req *pb.CreateCompanyV1Request,
) (*pb.CreateCompanyV1Response, error) {
	log.Debug().Msg("CreateCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}

func (o *CompanyAPI) ListCompanyV1(
	ctx context.Context,
	req *pb.ListCompanyV1Request,
) (*pb.ListCompanyV1Response, error) {
	log.Debug().Msg("ListCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}

func (o *CompanyAPI) RemoveCompanyV1(
	ctx context.Context,
	req *pb.RemoveCompanyV1Request,
) (*pb.RemoveCompanyV1Response, error) {
	log.Debug().Msg("RemoveCompanyV1")

	return nil, status.Error(codes.Internal, "not implemented")
}
