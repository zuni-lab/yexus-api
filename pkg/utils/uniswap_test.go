package utils_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func TestCalculatePrice(t *testing.T) {
	sqrtPriceX96 := "1614273077879348794073240928982532"
	sqrtPriceX96Int, ok := big.NewInt(0).SetString(sqrtPriceX96, 10)
	if !ok {
		t.Fatalf("failed to convert sqrtPriceX96 to big.Int")
	}

	price := utils.CalculatePrice(sqrtPriceX96Int, 18, 6, true)

	fmt.Println("price", price)

	var priceNumeric pgtype.Numeric
	priceStr := price.Text('f', 18) // Use 18 decimal places for precision

	log.Info().
		Str("price", priceStr).
		Msg("Calculated price")

	if err := priceNumeric.Scan(priceStr); err != nil {
		t.Fatalf("failed to convert price to numeric: %v", err)
	}

	fmt.Println("priceNumeric", priceNumeric)
}
