package main

import (
	pm "catch/metrics_main/metrics"
	"context"
)

func (gs *GameServer) GetNumberOfActiveUsers(context.Context, *pm.GetRequest) (*pm.MetricsReply, error) {
	return &pm.MetricsReply{MetricName: "number_of_active_users", MetricValue: int64(len(gs.userCache))}, nil
}

func (gs *GameServer) GetGameBooferSize(context.Context, *pm.GetRequest) (*pm.MetricsReply, error) {
	return &pm.MetricsReply{MetricName: "game_boofer_size", MetricValue: int64(len(gs.gameBoofer.games))}, nil
}

func (gs *GameServer) GetGameBooferGamesAvailable(context.Context, *pm.GetRequest) (*pm.MetricsReply, error) {
	return &pm.MetricsReply{MetricName: "game_boofer_games_available", MetricValue: int64(gs.gameBoofer.ptr)}, nil
}
