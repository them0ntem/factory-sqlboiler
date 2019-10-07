//go:generate sqlboiler --wipe crdb

package main

import (
	"database/sql"
	"github.com/davecgh/go-spew/spew"
	"github.com/volatiletech/null"
	"themontem/factory/factory"
	"themontem/factory/models"

	_ "github.com/lib/pq"
)

func main() {
	dbInstance, err := sql.Open("postgres", "host=192.168.56.101 port=36257 user=root dbname=startrek sslmode=disable")
	if err != nil {
		panic(err)
	}

	factory.DB = dbInstance

	episode := factory.NewEpisode(true, func(g *factory.EpisodeGenerator) {
		g.Season = func() null.Int64 {
			return null.NewInt64(5, false)
		}
	})

	quote := factory.NewQuote(true, func(qg *factory.QuoteGenerator) {
		qg.Episode = func(dbInsert bool, opts ...factory.EpisodeOption) *models.Episode {
			return factory.NewEpisode(true, func(eg *factory.EpisodeGenerator) {
				eg.Num = func() null.Int64 {
					return null.Int64From(100)
				}
			})
		}
	})

	quoteGenerator := factory.NewQuoteGenerator(func(generator *factory.QuoteGenerator) {
		generator.Characters = func() null.String {
			return null.StringFrom("themontem")
		}
	})

	q1 := factory.NewQuote(true, factory.LoadGenerator(quoteGenerator))

	spew.Dump(episode)
	spew.Dump(quote)
	spew.Dump(q1)
}
