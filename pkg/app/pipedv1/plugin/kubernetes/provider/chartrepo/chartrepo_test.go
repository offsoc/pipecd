package chartrepo

import (
	"context"
	"os/exec"
	"testing"

	"github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/kubernetes/toolregistry"
	"github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/toolregistry/toolregistrytest"
	"github.com/pipe-cd/pipecd/pkg/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func helm(t *testing.T) string {
	t.Helper()

	r, err := toolregistrytest.NewToolRegistry(t)
	require.NoError(t, err)
	t.Cleanup(func() { r.Close() })

	registry := toolregistry.NewRegistry(r)

	helmPath, err := registry.Helm(context.Background(), "3.16.1")
	require.NoError(t, err)

	return helmPath
}

func TestAdd(t *testing.T) {
	t.Parallel()

	var (
		ctx      = context.Background()
		helmPath = helm(t)
	)

	repos := []config.HelmChartRepository{
		{
			// example from https://helm.sh/docs/intro/quickstart/#initialize-a-helm-chart-repository
			Name:    "bitnami",
			Address: "https://charts.bitnami.com/bitnami",
		},
	}

	require.NoError(t, Add(ctx, helmPath, repos, zap.NewNop()))
	t.Cleanup(func() {
		for _, repo := range repos {
			cmd := exec.CommandContext(ctx, helmPath, "repo", "remove", repo.Name)
			out, err := cmd.CombinedOutput()
			require.NoError(t, err, string(out))
		}
	})

	out, err := exec.CommandContext(ctx, helmPath, "repo", "list").CombinedOutput()
	require.NoError(t, err)
	require.Contains(t, string(out), "bitnami")
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	var (
		ctx      = context.Background()
		helmPath = helm(t)
	)

	require.NoError(t, Update(ctx, helmPath, zap.NewNop()))
}
