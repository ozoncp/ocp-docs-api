package api

import (
	"context"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpDocsApiServer
}

func NewDocsApi() desc.OcpDocsApiServer {
	return &api{}
}

func (a *api) ListDocsV1(
	ctx context.Context,
	req *desc.ListDocsV1Request,
) (*desc.ListDocsV1Response, error) {
	log.Print("ListDocsV1", req)
	return &desc.ListDocsV1Response{}, nil
}

func (a *api) DescribeDocV1(
	ctx context.Context,
	req *desc.DescribeDocV1Request,
) (*desc.DescribeDocV1Response, error) {
	log.Print("DescribeDocV1", req)
	return &desc.DescribeDocV1Response{}, nil
}

func (a *api) CreateDocV1(
	ctx context.Context,
	req *desc.CreateDocV1Request,
) (*desc.CreateDocV1Response, error) {
	log.Print("CreateDocV1", req)
	return &desc.CreateDocV1Response{}, nil
}

func (a *api) RemoveDocV1(
	ctx context.Context,
	req *desc.RemoveDocV1Request,
) (*desc.RemoveDocV1Response, error) {
	log.Print("RemoveDocV1", req)
	return &desc.RemoveDocV1Response{}, nil
}