package main

import "go.uber.org/zap"

func InitZap() *zap.Logger {
	z, _ := zap.NewProduction()
	return z
}
