package utils

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

// BigFloatToNumeric converts a big.Float to pgtype.Numeric
// preserving full precision including decimals
func BigFloatToNumeric(f *big.Float) (pgtype.Numeric, error) {
	str := f.Text('f', -1)
	var numeric pgtype.Numeric
	err := numeric.Scan(str)
	if err != nil {
		return pgtype.Numeric{}, err
	}
	return numeric, nil
}
