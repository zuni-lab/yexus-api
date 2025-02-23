package handlers

import (
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/pkg/evm"
)

func HandleSwap(event *evm.UniswapV3Swap) error {
	log.Info().Any("event", event).Msgf("[Manager] [HandleSwap] handled %s event", event.Raw.Address.Hex())
	return nil
}
