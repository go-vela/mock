// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

// ImageService implements all the image
// related functions for the Docker mock.
type ImageService struct{}

// BuildCachePrune is a helper function to simulate
// a mocked call to prune the build cache for the
// Docker daemon.
func (i *ImageService) BuildCachePrune(ctx context.Context, opts types.BuildCachePruneOptions) (*types.BuildCachePruneReport, error) {
	return nil, nil
}

// BuildCancel is a helper function to simulate
// a mocked call to cancel building a Docker image.
func (i *ImageService) BuildCancel(ctx context.Context, id string) error {
	return nil
}

// ImageBuild is a helper function to simulate
// a mocked call to build a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageBuild
func (i *ImageService) ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	return types.ImageBuildResponse{}, nil
}

// ImageCreate is a helper function to simulate
// a mocked call to create a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageCreate
func (i *ImageService) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error) {
	return nil, nil
}

// ImageHistory is a helper function to simulate
// a mocked call to inspect the history for a
// Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageHistory
func (i *ImageService) ImageHistory(ctx context.Context, image string) ([]image.HistoryResponseItem, error) {
	return nil, nil
}

// ImageImport is a helper function to simulate
// a mocked call to import a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageImport
func (i *ImageService) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error) {
	return nil, nil
}

// ImageInspectWithRaw is a helper function to simulate
// a mocked call to inspect a Docker image and return
// the raw body received from the API.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageInspectWithRaw
func (i *ImageService) ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error) {
	return types.ImageInspect{}, nil, nil
}

// ImageList is a helper function to simulate
// a mocked call to list Docker images.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageList
func (i *ImageService) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	return nil, nil
}

// ImageLoad is a helper function to simulate
// a mocked call to load a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageLoad
func (i *ImageService) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error) {
	return types.ImageLoadResponse{}, nil
}

// ImagePull is a helper function to simulate
// a mocked call to pull a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImagePull
func (i *ImageService) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	return nil, nil
}

// ImagePush is a helper function to simulate
// a mocked call to push a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImagePush
func (i *ImageService) ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error) {
	return nil, nil
}

// ImageRemove is a helper function to simulate
// a mocked call to remove a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageRemove
func (i *ImageService) ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	return nil, nil
}

// ImageSave is a helper function to simulate
// a mocked call to save a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageSave
func (i *ImageService) ImageSave(ctx context.Context, images []string) (io.ReadCloser, error) {
	return nil, nil
}

// ImageSearch is a helper function to simulate
// a mocked call to search for a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageSearch
func (i *ImageService) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error) {
	return nil, nil
}

// ImageTag is a helper function to simulate
// a mocked call to tag a Docker image.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImageTag
func (i *ImageService) ImageTag(ctx context.Context, image, ref string) error {
	return nil
}

// ImagesPrune is a helper function to simulate
// a mocked call to prune Docker images.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.ImagesPrune
func (i *ImageService) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error) {
	return types.ImagesPruneReport{}, nil
}

// WARNING: DO NOT REMOVE THIS UNDER ANY CIRCUMSTANCES
//
// This line serves as a quick and efficient way to ensure that our
// ImageService satisfies the ImageAPIClient interface that
// the Docker client expects.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#ImageAPIClient
var _ client.ImageAPIClient = (*ImageService)(nil)