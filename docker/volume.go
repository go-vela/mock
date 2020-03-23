// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// VolumeService implements all the volume
// related functions for the Docker mock.
type VolumeService struct{ client.VolumeAPIClient }

// WARNING: DO NOT REMOVE THIS UNDER ANY CIRCUMSTANCES
//
// This line serves as a quick and efficient way to ensure that our
// VolumeService satisfies the VolumeAPIClient interface that
// the Docker client expects.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#VolumeAPIClient
var _ client.VolumeAPIClient = (*VolumeService)(nil)

// VolumeCreate is a helper function to simulate
// a mocked call to create a Docker volume.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumeCreate
func (v *VolumeService) VolumeCreate(ctx context.Context, options volume.VolumeCreateBody) (types.Volume, error) {
	return types.Volume{}, nil
}

// VolumeInspect is a helper function to simulate
// a mocked call to inspect a Docker volume.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumeInspect
func (v *VolumeService) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error) {
	return types.Volume{}, nil
}

// VolumeInspectWithRaw is a helper function to simulate
// a mocked call to inspect a Docker volume and return
// the raw body received from the API.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumeInspectWithRaw
func (v *VolumeService) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error) {
	return types.Volume{}, nil, nil
}

// VolumeList is a helper function to simulate
// a mocked call to list Docker volumes.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumeList
func (v *VolumeService) VolumeList(ctx context.Context, filter filters.Args) (volume.VolumeListOKBody, error) {
	return volume.VolumeListOKBody{}, nil
}

// VolumeRemove is a helper function to simulate
// a mocked call to remove Docker a volume.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumeRemove
func (v *VolumeService) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	return nil
}

// VolumesPrune is a helper function to simulate
// a mocked call to prune Docker volumes.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.VolumesPrune
func (v *VolumeService) VolumesPrune(ctx context.Context, pruneFilter filters.Args) (types.VolumesPruneReport, error) {
	return types.VolumesPruneReport{}, nil
}
