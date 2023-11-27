package infrastructure

import (
	"fmt"

	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) {

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("zap.NewProduction failed: %w", err)
	}
	return logger.Sugar(), nil

}
