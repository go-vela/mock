// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// NetworkService implements all the network
// related functions for the Docker mock.
type NetworkService struct{}

// WARNING: DO NOT REMOVE THIS UNDER ANY CIRCUMSTANCES
//
// This line serves as a quick and efficient way to ensure that our
// NetworkService satisfies the NetworkAPIClient interface that
// the Docker client expects.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#NetworkAPIClient
var _ client.NetworkAPIClient = (*NetworkService)(nil)

// NetworkConnect is a helper function to simulate
// a mocked call to connect to a Docker network.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkConnect
func (n *NetworkService) NetworkConnect(ctx context.Context, network, container string, config *network.EndpointSettings) error {
	return nil
}

// NetworkCreate is a helper function to simulate
// a mocked call to create a Docker network.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkCreate
func (n *NetworkService) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	return types.NetworkCreateResponse{}, nil
}

// NetworkDisconnect is a helper function to simulate
// a mocked call to disconnect from a Docker network.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkDisconnect
func (n *NetworkService) NetworkDisconnect(ctx context.Context, network, container string, force bool) error {
	return nil
}

// NetworkInspect is a helper function to simulate
// a mocked call to inspect a Docker network.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkInspect
func (n *NetworkService) NetworkInspect(ctx context.Context, network string, options types.NetworkInspectOptions) (types.NetworkResource, error) {
	return types.NetworkResource{}, nil
}

// NetworkInspectWithRaw is a helper function to simulate
// a mocked call to inspect a Docker network and return
// the raw body received from the API.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkInspectWithRaw
func (n *NetworkService) NetworkInspectWithRaw(ctx context.Context, network string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error) {
	return types.NetworkResource{}, nil, nil
}

// NetworkList is a helper function to simulate
// a mocked call to list Docker networks.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkList
func (n *NetworkService) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	return nil, nil
}

// NetworkRemove is a helper function to simulate
// a mocked call to remove Docker a network.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworkRemove
func (n *NetworkService) NetworkRemove(ctx context.Context, network string) error {
	return nil
}

// NetworksPrune is a helper function to simulate
// a mocked call to prune Docker networks.
//
// https://pkg.go.dev/github.com/docker/docker/client?tab=doc#Client.NetworksPrune
func (n *NetworkService) NetworksPrune(ctx context.Context, pruneFilter filters.Args) (types.NetworksPruneReport, error) {
	return types.NetworksPruneReport{}, nil
}
