package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

// CommandHandler 定义泛型接收一个查询 Q，返回结果 R
type CommandHandler[C, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

func ApplyCommandDecorators[H, R any](handler CommandHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		logger: logger,
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: metricsClient,
		},
	}
}
