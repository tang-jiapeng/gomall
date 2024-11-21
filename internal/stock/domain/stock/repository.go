package stock

import (
	"context"
	"fmt"
	"github.com/0X10CC/gorder/common/genproto/orderpb"
	"strings"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("stock items not found: %s", strings.Join(e.Missing, ","))
}
