package trace

import (
	"context"

	trpc "trpc.group/trpc-go/trpc-go"
)

const TraceIDKey = "trace_id"

func ID(ctx context.Context) string {
	return string(trpc.GetMetaData(ctx, TraceIDKey))
}
