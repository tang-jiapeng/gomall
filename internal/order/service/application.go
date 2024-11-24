package service

import (
	"context"
	"github.com/0X10CC/gorder/common/metrics"
	"github.com/0X10CC/gorder/order/adapters"
	"github.com/0X10CC/gorder/order/app"
	"github.com/0X10CC/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderInmemRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderInmemRepo, logger, metricClient),
		},
	}
}
