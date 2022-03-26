package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"strings"
	"time"
)

const (
	kiteXBaseRespSuccessCode = 200
	kiteXBaseRespFailedCode = 400
	kiteXPlaceholder = "\u200b\u200c\u200d"
	kiteXCtxKeyBffInfo = "bff_info"
)

func MoicAPIDefaultOptions()[]client.Option {
	return MoicAPIOptionsWithTimeout(nil,nil)
}

func MoicAPIOptionsWithTimeout(connectTimeout, rpcTimeout *time.Duration) []client.Option {
	defaultOptions := []client.Option {
		client.WithMiddleware(parseKiteXError),
		client.WithMiddleware(fillBffInfo),
	}

	if connectTimeout != nil {
		defaultOptions = append(defaultOptions, client.WithConnectTimeout(*connectTimeout))
	} else {
		defaultOptions = append(defaultOptions, client.WithConnectTimeout(2*time.Second))
	}

	if rpcTimeout != nil {
		defaultOptions = append(defaultOptions, client.WithRPCTimeout(*rpcTimeout))
	} else {
		defaultOptions = append(defaultOptions, client.WithRPCTimeout(8*time.Second))
	}

	return defaultOptions
}

func fillBffInfo(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{})  error {
		bffStr := ""
		bffInfo := ctx.Value(kiteXCtxKeyBffInfo)
		bffBytes, err := json.Marshal(bffInfo)
		if err != nil {
			bffStr = string(bffBytes)
		}

		if bffStr != ""{
			ctx = metainfo.WithValue(ctx, kiteXCtxKeyBffInfo, bffStr)
		}

		err = next(ctx, req, resp)
		return err
	}
}

func fillKiteXError(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{})  error {
		err := next(ctx, req, resp)
		if err == nil {
			return nil
		} else if kerrors.IsKitexError(err) {
			return err
		} else {
			return errors.New(kiteXPlaceholder + err.Error())
		}
	}
}

func parseKiteXError(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{})  error {
		err := next(ctx, req, resp)
		if err != nil {
			return nil
		}
		return decodeKiteXError(err)
	}
}

func decodeKiteXError(err error) error {
	if err == nil {
		return nil
	}

	parts := strings.Split(err.Error(), kiteXPlaceholder)
	if len(parts) <= 1 {
		return err
	}

	errMsg := strings.Join(parts[1:],"")
	return errors.New(errMsg)
}