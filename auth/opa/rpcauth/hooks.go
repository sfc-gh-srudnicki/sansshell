package rpcauth

import (
	"context"
	"net"
)

// AuthzHookFunc implements RpcAuthzHook for a simple function
type RpcAuthzHookFunc func(context.Context, *RpcAuthInput) error

func (r RpcAuthzHookFunc) Hook(ctx context.Context, input *RpcAuthInput) error {
	return r(ctx, input)
}

// An HookPredicate returns true if a conditional hook should run
type HookPredicate func(*RpcAuthInput) bool

// HookIf wraps an existing hook, and only executes it when
// the provided condition returns true
func HookIf(hook RpcAuthzHook, condition HookPredicate) RpcAuthzHook {
	return &conditionalHook{
		hook:      hook,
		predicate: condition,
	}
}

type conditionalHook struct {
	hook      RpcAuthzHook
	predicate HookPredicate
}

func (c *conditionalHook) Hook(ctx context.Context, input *RpcAuthInput) error {
	if c.predicate(input) {
		return c.hook.Hook(ctx, input)
	}
	return nil
}

// HostNetHook returns an RpcAuthzHook that sets host networking information.
func HostNetHook(addr net.Addr) RpcAuthzHook {
	return RpcAuthzHookFunc(func(ctx context.Context, input *RpcAuthInput) error {
		if input.Host == nil {
			input.Host = &HostAuthInput{}
		}
		input.Host.Net = NetInputFromAddr(addr)
		return nil
	})
}