package factory

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/types"
	"math/rand"
	"time"
)

var (
	DB     *sql.DB
	ctx    = context.Background()
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type UUIDFunc func() string
type NullStringFunc func() null.String
type NullInt64Func func() null.Int64
type NullDecimalFunc func() types.NullDecimal
