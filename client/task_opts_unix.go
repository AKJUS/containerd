//go:build !windows

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package client

import (
	"context"
)

// WithNoNewKeyring causes tasks not to be created with a new keyring for secret storage.
// There is an upper limit on the number of keyrings in a linux system
func WithNoNewKeyring(ctx context.Context, c *Client, ti *TaskInfo) error {
	opts, err := ti.getRuncOptions()
	if err != nil {
		return err
	}
	opts.NoNewKeyring = true
	return nil
}

// WithNoPivotRoot instructs the runtime not to you pivot_root
func WithNoPivotRoot(_ context.Context, _ *Client, ti *TaskInfo) error {
	opts, err := ti.getRuncOptions()
	if err != nil {
		return err
	}
	opts.NoPivotRoot = true
	return nil
}

// WithShimCgroup sets the existing cgroup for the shim
func WithShimCgroup(path string) NewTaskOpts {
	return func(ctx context.Context, c *Client, ti *TaskInfo) error {
		opts, err := ti.getRuncOptions()
		if err != nil {
			return err
		}
		opts.ShimCgroup = path
		return nil
	}
}

// WithUIDOwner allows console I/O to work with the remapped UID in user namespace
func WithUIDOwner(uid uint32) NewTaskOpts {
	return func(ctx context.Context, c *Client, ti *TaskInfo) error {
		opts, err := ti.getRuncOptions()
		if err != nil {
			return err
		}
		opts.IoUid = uid
		return nil
	}
}

// WithGIDOwner allows console I/O to work with the remapped GID in user namespace
func WithGIDOwner(gid uint32) NewTaskOpts {
	return func(ctx context.Context, c *Client, ti *TaskInfo) error {
		opts, err := ti.getRuncOptions()
		if err != nil {
			return err
		}
		opts.IoGid = gid
		return nil
	}
}
