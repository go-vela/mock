// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

// Version represents the supported
// Docker API version for the mock.
const Version = "v1.40"

// New returns a client that is capable of handling
// Docker client calls and returning stub responses.
// nolint:golint // returning unexported type intentionally
func New() (*mock, error) {
	return &mock{
		ConfigService:       ConfigService{},
		ContainerService:    ContainerService{},
		DistributionService: DistributionService{},
		ImageService:        ImageService{},
		NetworkService:      NetworkService{},
		NodeService:         NodeService{},
		PluginService:       PluginService{},
		SecretService:       SecretService{},
		ServiceService:      ServiceService{},
		SystemService:       SystemService{},
		SwarmService:        SwarmService{},
		VolumeService:       VolumeService{},
		Version:             Version,
	}, nil
}
