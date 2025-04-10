// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
	cerrdefs "github.com/containerd/errdefs"
	ncTypes "github.com/containerd/nerdctl/v2/pkg/api/types"

	"github.com/runfinch/finch-daemon/pkg/errdefs"
)

func (s *service) Unpause(ctx context.Context, cid string, options ncTypes.ContainerUnpauseOptions) error {
	cont, err := s.getContainer(ctx, cid)
	if err != nil {
		if cerrdefs.IsNotFound(err) {
			return errdefs.NewNotFound(err)
		}
		return err
	}
	status := s.client.GetContainerStatus(ctx, cont)
	if status != containerd.Paused {
		if status == containerd.Running {
			return errdefs.NewConflict(fmt.Errorf("container %s is already running", cid))
		}
		return errdefs.NewConflict(fmt.Errorf("container %s is not paused", cid))
	}

	err = s.nctlContainerSvc.UnpauseContainer(ctx, cid, options)
	if err != nil {
		return err
	}

	return nil
}
