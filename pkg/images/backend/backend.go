///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package backend

import (
	"context"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// Backend is the interface for image manager backend, like Knative
type Backend interface {
	AddImage(ctx context.Context, image *v1.Image) (*v1.Image, error)
	GetImage(ctx context.Context, meta *v1.Meta) (*v1.Image, error)
	DeleteImage(ctx context.Context, meta *v1.Meta) error
	ListImage(ctx context.Context, meta *v1.Meta) ([]*v1.Image, error)
	UpdateImage(ctx context.Context, image *v1.Image) (*v1.Image, error)
}
