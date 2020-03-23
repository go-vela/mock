// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

import (
	"github.com/docker/docker/client"
)

// Version represents the supported
// Docker API version for the mock.
const Version = "v1.38"

type mock struct {
	// Services
	ContainerService
	ImageService
	NetworkService
	VolumeService

	// Docker API version for the mock
	Version string
}

// New returns a client that is capable of handling
// Docker client calls and returning stub responses.
func New() (*mock, error) {
	return &mock{
		ContainerService: ContainerService{},
		ImageService:     ImageService{},
		NetworkService:   NetworkService{},
		VolumeService:    VolumeService{},
		Version:          Version,
	}, nil
}

// WARNING: DO NOT REMOVE THIS UNDER ANY CIRCUMSTANCES
//
// This line serves as a quick and efficient way to ensure
// that our mock satisfies the Go interface that the
// Docker client expects.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#CommonAPIClient
var _ client.CommonAPIClient = (*mock)(nil)
