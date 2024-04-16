package config_test

import (
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/pkg/config"
	"github.com/greenplum-db/gpdb/gp/internal/pkg/enums"
	"github.com/greenplum-db/gpdb/gp/testutils"
	"github.com/greenplum-db/gpdb/gp/testutils/exectest"
)

func init() {
	exectest.RegisterMains()
}

func TestConfig(t *testing.T) {

	t.Run("config not present at location", func(t *testing.T) {

		testConfig := config.New()
		err := testConfig.Load(".")
		if err == nil {
			t.Fatalf("Expected %s, got %d", "error", err)
		}

	})

	t.Run("config loaded succesfully with defaults", func(t *testing.T) {

		testConfig := config.New()
		testConfig.SetName("gpdeploy.testconf")
		err := testConfig.Load("../../../test/data")

		testutils.Assert(t, nil, err, "Failed to load configs")
		testutils.Assert(t, enums.DeploymentTypeMirrorless, testConfig.GetDatabaseConfig().DeploymentType, "")

		testutils.Assert(t, 4506, testConfig.GetInfraConfig().RequestPort, "")
		testutils.Assert(t, 4505, testConfig.GetInfraConfig().PublishPort, "")
	})

	t.Run("config loaded succesfully overriding defaults", func(t *testing.T) {

		testConfig := config.New()
		testConfig.SetName("gpdeploy_override.testconf")
		err := testConfig.Load("../../../test/data")

		testutils.Assert(t, nil, err, "Failed to load configs")
		testutils.Assert(t, enums.DeploymentTypeMirrored, testConfig.GetDatabaseConfig().DeploymentType, "")
		testutils.Assert(t, 5001, testConfig.GetInfraConfig().RequestPort, "")
		testutils.Assert(t, 5002, testConfig.GetInfraConfig().PublishPort, "")

	})
}
