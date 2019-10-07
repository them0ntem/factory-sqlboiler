package factory

import (
	"github.com/icrowley/fake"
	"github.com/volatiletech/sqlboiler/boil"
	"themontem/factory/models"
	"time"
)

type ModelEpisodeFunc func(dbInsert bool, opts ...EpisodeOption) *models.Episode

type EpisodeOption func(*EpisodeGenerator)

type EpisodeGenerator struct {
	ID       UUIDFunc
	Season   NullInt64Func
	Num      NullInt64Func
	Title    NullStringFunc
	Stardate NullDecimalFunc
}

func (g *EpisodeGenerator) generate(dbInsert bool) *models.Episode {
	episode := &models.Episode{
		ID:       g.ID(),
		Season:   g.Season(),
		Num:      g.Num(),
		Title:    g.Title(),
		Stardate: g.Stardate(),
	}

	if dbInsert {
		if err := episode.Insert(ctx, DB, boil.Infer()); err != nil {
			panic(err)
		}
	} else {
		episode.ID = g.ID()
		episode.CreatedAt = time.Now().In(boil.GetLocation())
		episode.UpdatedAt = time.Now().In(boil.GetLocation())
	}

	return episode
}

func NewEpisodeGenerator(opts ...EpisodeOption) *EpisodeGenerator {
	generator := &EpisodeGenerator{
		ID:       UUID,
		Season:   nullInt64(false, random.Int63n, 1000),
		Num:      nullInt64(false, random.Int63n, 1000),
		Title:    nullString(false, fake.Brand),
		Stardate: nullDecimal(false, random.Int),
	}

	for _, opt := range opts {
		opt(generator)
	}

	return generator
}

func NewEpisode(dbInsert bool, opts ...EpisodeOption) *models.Episode {
	generator := NewEpisodeGenerator(opts...)
	return generator.generate(dbInsert)
}
