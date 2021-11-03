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
	totalcompanyNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_company_api_company_not_found_total",
		Help: "Total number of companys that were not found",
	})
)

type companyAPI struct {
	pb.UnimplementedbsscompanyApiServiceServer
	repo repo.Repo
}

// NewcompanyAPI returns api of bss-company-api service
func NewcompanyAPI(r repo.Repo) pb.bsscompanyApiServiceServer {
	return &companyAPI{repo: r}
}

func (o *companyAPI) DescribecompanyV1(
	ctx context.Context,
	req *pb.DescribecompanyV1Request,
) (*pb.DescribecompanyV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribecompanyV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	company, err := o.repo.Describecompany(ctx, req.companyId)
	if err != nil {
		log.Error().Err(err).Msg("DescribecompanyV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if company == nil {
		log.Debug().Uint64("companyId", req.companyId).Msg("company not found")
		totalcompanyNotFound.Inc()

		return nil, status.Error(codes.NotFound, "company not found")
	}

	log.Debug().Msg("DescribecompanyV1 - success")

	return &pb.DescribecompanyV1Response{
		Value: &pb.company{
			Id:  company.ID,
			Foo: company.Foo,
		},
	}, nil
}
