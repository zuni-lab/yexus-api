package db

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func Init(ctx context.Context, dbSource string, migrationURL string) {
	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(migrationURL, dbSource)

	DB = NewStore(connPool)

	log.Info().Msg("db initialized successfully")

	runSeed()
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runSeed() {
	log.Info().Msg("running seed")
	result, err := DB.CreateBatchPoolsTx(context.Background(), CreateBatchPoolsTxParams{
		Pools: poolsSeed,
	})
	if err != nil {
		log.Warn().Err(err).Msg("seed operation encountered an error (might be due to existing data)")
		return
	}
	log.Info().Msgf("successfully seeded %d pools", len(result.Pools))
}

func Close() {
	DB.connPool.Close()
}
