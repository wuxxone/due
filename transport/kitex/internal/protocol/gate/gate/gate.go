// Code generated by Kitex v0.9.1. DO NOT EDIT.

package gate

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	gate "github.com/dobyte/due/transport/kitex/v2/internal/protocol/gate"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Bind": kitex.NewMethodInfo(
		bindHandler,
		newGateBindArgs,
		newGateBindResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Unbind": kitex.NewMethodInfo(
		unbindHandler,
		newGateUnbindArgs,
		newGateUnbindResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetIP": kitex.NewMethodInfo(
		getIPHandler,
		newGateGetIPArgs,
		newGateGetIPResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Stat": kitex.NewMethodInfo(
		statHandler,
		newGateStatArgs,
		newGateStatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Disconnect": kitex.NewMethodInfo(
		disconnectHandler,
		newGateDisconnectArgs,
		newGateDisconnectResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Push": kitex.NewMethodInfo(
		pushHandler,
		newGatePushArgs,
		newGatePushResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Multicast": kitex.NewMethodInfo(
		multicastHandler,
		newGateMulticastArgs,
		newGateMulticastResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Broadcast": kitex.NewMethodInfo(
		broadcastHandler,
		newGateBroadcastArgs,
		newGateBroadcastResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	gateServiceInfo                = NewServiceInfo()
	gateServiceInfoForClient       = NewServiceInfoForClient()
	gateServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return gateServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return gateServiceInfoForStreamClient
}

// for transporter client
func serviceInfoForClient() *kitex.ServiceInfo {
	return gateServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Gate"
	handlerType := (*gate.Gate)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "gate",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func bindHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateBindArgs)
	realResult := result.(*gate.GateBindResult)
	success, err := handler.(gate.Gate).Bind(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateBindArgs() interface{} {
	return gate.NewGateBindArgs()
}

func newGateBindResult() interface{} {
	return gate.NewGateBindResult()
}

func unbindHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateUnbindArgs)
	realResult := result.(*gate.GateUnbindResult)
	success, err := handler.(gate.Gate).Unbind(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateUnbindArgs() interface{} {
	return gate.NewGateUnbindArgs()
}

func newGateUnbindResult() interface{} {
	return gate.NewGateUnbindResult()
}

func getIPHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateGetIPArgs)
	realResult := result.(*gate.GateGetIPResult)
	success, err := handler.(gate.Gate).GetIP(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateGetIPArgs() interface{} {
	return gate.NewGateGetIPArgs()
}

func newGateGetIPResult() interface{} {
	return gate.NewGateGetIPResult()
}

func statHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateStatArgs)
	realResult := result.(*gate.GateStatResult)
	success, err := handler.(gate.Gate).Stat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateStatArgs() interface{} {
	return gate.NewGateStatArgs()
}

func newGateStatResult() interface{} {
	return gate.NewGateStatResult()
}

func disconnectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateDisconnectArgs)
	realResult := result.(*gate.GateDisconnectResult)
	success, err := handler.(gate.Gate).Disconnect(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateDisconnectArgs() interface{} {
	return gate.NewGateDisconnectArgs()
}

func newGateDisconnectResult() interface{} {
	return gate.NewGateDisconnectResult()
}

func pushHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GatePushArgs)
	realResult := result.(*gate.GatePushResult)
	success, err := handler.(gate.Gate).Push(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGatePushArgs() interface{} {
	return gate.NewGatePushArgs()
}

func newGatePushResult() interface{} {
	return gate.NewGatePushResult()
}

func multicastHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateMulticastArgs)
	realResult := result.(*gate.GateMulticastResult)
	success, err := handler.(gate.Gate).Multicast(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateMulticastArgs() interface{} {
	return gate.NewGateMulticastArgs()
}

func newGateMulticastResult() interface{} {
	return gate.NewGateMulticastResult()
}

func broadcastHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gate.GateBroadcastArgs)
	realResult := result.(*gate.GateBroadcastResult)
	success, err := handler.(gate.Gate).Broadcast(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGateBroadcastArgs() interface{} {
	return gate.NewGateBroadcastArgs()
}

func newGateBroadcastResult() interface{} {
	return gate.NewGateBroadcastResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Bind(ctx context.Context, req *gate.BindRequest) (r *gate.BindResponse, err error) {
	var _args gate.GateBindArgs
	_args.Req = req
	var _result gate.GateBindResult
	if err = p.c.Call(ctx, "Bind", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Unbind(ctx context.Context, req *gate.UnbindRequest) (r *gate.UnbindResponse, err error) {
	var _args gate.GateUnbindArgs
	_args.Req = req
	var _result gate.GateUnbindResult
	if err = p.c.Call(ctx, "Unbind", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIP(ctx context.Context, req *gate.GetIPRequest) (r *gate.GetIPResponse, err error) {
	var _args gate.GateGetIPArgs
	_args.Req = req
	var _result gate.GateGetIPResult
	if err = p.c.Call(ctx, "GetIP", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Stat(ctx context.Context, req *gate.StatRequest) (r *gate.StatResponse, err error) {
	var _args gate.GateStatArgs
	_args.Req = req
	var _result gate.GateStatResult
	if err = p.c.Call(ctx, "Stat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Disconnect(ctx context.Context, req *gate.DisconnectRequest) (r *gate.DisconnectResponse, err error) {
	var _args gate.GateDisconnectArgs
	_args.Req = req
	var _result gate.GateDisconnectResult
	if err = p.c.Call(ctx, "Disconnect", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Push(ctx context.Context, req *gate.PushRequest) (r *gate.PushResponse, err error) {
	var _args gate.GatePushArgs
	_args.Req = req
	var _result gate.GatePushResult
	if err = p.c.Call(ctx, "Push", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Multicast(ctx context.Context, req *gate.MulticastRequest) (r *gate.MulticastResponse, err error) {
	var _args gate.GateMulticastArgs
	_args.Req = req
	var _result gate.GateMulticastResult
	if err = p.c.Call(ctx, "Multicast", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Broadcast(ctx context.Context, req *gate.BroadcastRequest) (r *gate.BroadcastResponse, err error) {
	var _args gate.GateBroadcastArgs
	_args.Req = req
	var _result gate.GateBroadcastResult
	if err = p.c.Call(ctx, "Broadcast", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
