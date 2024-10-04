// Copyright 2024 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package chartrepo manages a list of configured helm repositories.
package chartrepo

import (
	"context"
	"fmt"
	"os/exec"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/pipe-cd/pipecd/pkg/config"
)

var updateGroup = &singleflight.Group{}

// Add installs all specified Helm Chart repositories.
// https://helm.sh/docs/topics/chart_repository/
// helm repo add fantastic-charts https://fantastic-charts.storage.googleapis.com
// helm repo add fantastic-charts https://fantastic-charts.storage.googleapis.com --username my-username --password my-password
func Add(ctx context.Context, helmPath string, repos []config.HelmChartRepository, logger *zap.Logger) error {
	for _, repo := range repos {
		args := []string{"repo", "add", repo.Name, repo.Address}
		if repo.Insecure {
			args = append(args, "--insecure-skip-tls-verify")
		}
		if repo.Username != "" || repo.Password != "" {
			args = append(args, "--username", repo.Username, "--password", repo.Password)
		}
		cmd := exec.CommandContext(ctx, helmPath, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add chart repository %s: %s (%w)", repo.Name, string(out), err)
		}
		logger.Info(fmt.Sprintf("successfully added chart repository: %s", repo.Name))
	}
	return nil
}

func Update(ctx context.Context, helmPath string, logger *zap.Logger) error {
	_, err, _ := updateGroup.Do("update", func() (interface{}, error) {
		return nil, update(ctx, helmPath, logger)
	})
	return err
}

func update(ctx context.Context, helmPath string, logger *zap.Logger) error {
	logger.Info("start updating Helm chart repositories")

	args := []string{"repo", "update"}
	cmd := exec.CommandContext(ctx, helmPath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update Helm chart repositories: %s (%w)", string(out), err)
	}

	logger.Info("successfully updated Helm chart repositories")
	return nil
}
