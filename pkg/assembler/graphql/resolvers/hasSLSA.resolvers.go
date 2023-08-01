package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestSlsa is the resolver for the ingestSLSA field.
func (r *mutationResolver) IngestSlsa(ctx context.Context, subject model.ArtifactInputSpec, builtFrom []*model.ArtifactInputSpec, builtBy model.BuilderInputSpec, slsa model.SLSAInputSpec) (*model.HasSlsa, error) {
	return r.Backend.IngestSLSA(ctx, subject, builtFrom, builtBy, slsa)
}

// IngestSLSAs is the resolver for the ingestSLSAs field.
func (r *mutationResolver) IngestSLSAs(ctx context.Context, subjects []*model.ArtifactInputSpec, builtFromList [][]*model.ArtifactInputSpec, builtByList []*model.BuilderInputSpec, slsaList []*model.SLSAInputSpec) ([]*model.HasSlsa, error) {
	return r.Backend.IngestSLSAs(ctx, subjects, builtFromList, builtByList, slsaList)
}

// HasSlsa is the resolver for the HasSLSA field.
func (r *queryResolver) HasSlsa(ctx context.Context, hasSLSASpec *model.HasSLSASpec) ([]*model.HasSlsa, error) {
	return r.Backend.HasSlsa(ctx, hasSLSASpec)
}
