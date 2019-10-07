package factory

import (
	"github.com/icrowley/fake"
	"github.com/volatiletech/sqlboiler/boil"
	"themontem/factory/models"
	"time"
)

type QuoteOption func(*QuoteGenerator)

func LoadGenerator(lg *QuoteGenerator) QuoteOption {
	return func(g *QuoteGenerator) {
		*g = *lg
	}
}

type QuoteGenerator struct {
	ID         UUIDFunc
	Quote      NullStringFunc
	Characters NullStringFunc
	Stardate   NullDecimalFunc
	Episode    ModelEpisodeFunc
}

func (g QuoteGenerator) generate(dbInsert bool) *models.Quote {
	episode := g.Episode(dbInsert)

	quote := &models.Quote{
		Quote:      g.Quote(),
		Characters: g.Characters(),
		Stardate:   g.Stardate(),
		EpisodeID:  episode.ID,
	}

	if dbInsert {
		if err := quote.Insert(ctx, DB, boil.Infer()); err != nil {
			panic(err)
		}
	} else {
		quote.ID = g.ID()
		quote.CreatedAt = time.Now().In(boil.GetLocation())
		quote.UpdatedAt = time.Now().In(boil.GetLocation())
	}

	quote.R = quote.R.NewStruct()
	quote.R.Episode = episode

	return quote
}

func NewQuoteGenerator(opts ...QuoteOption) *QuoteGenerator {
	generator := &QuoteGenerator{
		ID:         UUID,
		Quote:      nullString(false, fake.Brand),
		Characters: nullString(false, fake.Brand),
		Stardate:   nullDecimal(false, random.Int),
		Episode:    NewEpisode,
	}

	for _, opt := range opts {
		opt(generator)
	}

	return generator
}

func NewQuote(dbInsert bool, opts ...QuoteOption) *models.Quote {
	generator := NewQuoteGenerator(opts...)
	return generator.generate(dbInsert)
}
