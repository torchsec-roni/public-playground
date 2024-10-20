package my_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestSnapshotOverride(t *testing.T) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	snapshotName := "mysnapshot"
	err = postgresContainer.Snapshot(ctx, postgres.WithSnapshotName(snapshotName))
	if err != nil {
		log.Fatalf("Failed to snapshot test DB: %v", err)
	}

	// Despite the following comment on the Snapshot method, it does not work:
	// "If a snapshot already exists under the given/default name, it will be overwritten with the new snapshot."
	// Instead, we get the following error:
	// `Failed to snapshot test DB: non-zero exit code for restore command: (ERROR:  cannot drop a template database`
	err = postgresContainer.Snapshot(ctx, postgres.WithSnapshotName(snapshotName))
	if err != nil {
		log.Fatalf("Failed to snapshot test DB: %v", err)
	}
}

func TestDefaultSnapshot(t *testing.T) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	// Despite the following comment on the Snapshot method, it does not work:
	// "By default, the snapshot will be created under a database called migrated_template."
	// Instead, we get the following error:
	// `Failed to snapshot test DB: non-zero exit code for restore command: ~ERROR:  zero-length delimited identifier at or near """"
	//  LINE 1: DROP DATABASE IF EXISTS ""`
	err = postgresContainer.Snapshot(ctx)
	if err != nil {
		log.Fatalf("Failed to snapshot test DB: %v", err)
	}
}

func TestDefaultSnapshot2(t *testing.T) {
	ctx := context.Background()

	ctr, err := postgres.Run(ctx, "docker.io/postgres:16-alpine",
		postgres.WithDatabase("users"),
		postgres.BasicWaitStrategies(),
	)
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	err = ctr.Snapshot(ctx)
	require.NoError(t, err)
}
