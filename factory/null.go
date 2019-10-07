package factory

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/types"
)

func nullString(shouldBeNull bool, nextString func() string) NullStringFunc {
	if shouldBeNull {
		return func() null.String {
			return null.NewString("", false)
		}
	}

	return func() null.String {
		return null.StringFrom(nextString())
	}
}

func nullInt64(shouldBeNull bool, nextInt64 func(n int64) int64, maxValue int64) NullInt64Func {
	if shouldBeNull {
		return func() null.Int64 {
			return null.NewInt64(0, false)
		}
	}

	return func() null.Int64 {
		return null.Int64From(nextInt64(maxValue))
	}
}

func nullDecimal(shouldBeNull bool, nextInt func() int) NullDecimalFunc {
	if shouldBeNull {
		return func() types.NullDecimal {
			return types.NewNullDecimal(nil)
		}
	}

	return func() types.NullDecimal {
		randVal := fmt.Sprintf("%d.%d", nextInt()%10, nextInt()%10)
		random, success := new(decimal.Big).SetString(randVal)
		if !success {
			panic("randVal could not be turned into a decimal")
		}

		return types.NewNullDecimal(random)
	}
}
