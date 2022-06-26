package opentelemetry

import (
	"context"

	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
)

const (
	lDebug = "DEBUG"
	lInfo  = "INFO"
	lWarn  = "WARN"
	lErr   = "ERROR"
	lFatal = "FATAL"
)

type Adapter struct{}

func (Adapter) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	ctx, ocs := octrace.StartSpan(ctx, name)
	l := log.G(ctx).WithField("method", name)

	s := &span{s: ocs, l: l}
	ctx = log.WithLogger(ctx, s.Logger())

	return ctx, s
}
