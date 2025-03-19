package worker

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/yexus-api/pkg/db"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

type Scheduler struct {
	raw   *gocron.Scheduler
	atUtc time.Time
	jobs  []func()
	log   *zerolog.Event
}

func NewScheduler(at time.Time, name ...string) *Scheduler {
	nameStr := "Scheduler"
	if len(name) > 0 {
		nameStr = name[0]
	}

	atUtc := at.UTC()

	log.Info().Msgf("Scheduler %s at %s", nameStr, atUtc.String())

	return &Scheduler{
		raw:   gocron.NewScheduler(time.UTC),
		atUtc: atUtc,
		log: log.Info().Func(func(e *zerolog.Event) {
			e.Any(utils.YellowMsg("scheduler"), nameStr).Any(utils.BlueMsg("at"), atUtc.String())
		}),
	}
}

func (s *Scheduler) Start() {
	for _, job := range s.jobs {
		s.raw.Every(1).Day().At(s.atUtc).Do(job)
	}
	s.raw.StartAsync()
}

func (s *Scheduler) Shutdown() {
	s.raw.Stop()
}

func (s *Scheduler) AddJob(job func()) {
	// Wrap the original job with panic recovery
	wrappedJob := func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().
					Interface("panic", r).
					Msg("Recovered from panic in scheduled job")
			}
		}()
		job()
	}
	s.jobs = append(s.jobs, wrappedJob)
}

func (s *Scheduler) Len() int {
	return len(s.jobs)
}

// Run all the jobs at once
func Init() {
	var wg sync.Once
	wg.Do(func() {
		ctx := context.Background()

		var shouldFetch = false

		latestYieldMetric, err := db.DB.GetLatestYieldMetric(ctx)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				shouldFetch = true
			} else {
				log.Error().Msgf("failed to get latest yield metric: %s", err)
				return
			}
		}

		if latestYieldMetric.CreatedAt.Time.Add(time.Hour * 24).Before(time.Now()) {
			log.Info().Msg("Data is more than 24 hours old, fetching and updating yield metrics")
			shouldFetch = true
		}

		if !shouldFetch {
			log.Info().Msg("No need to fetch and update yield metrics")
			return
		}

		FetchAndUpdateYieldMetrics(ctx)

	})
}
