// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: lorawan-stack/api/qrcodegenerator.proto

/*
Package ttnpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ttnpb

import (
	"context"
	"io"
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_EndDeviceQRCodeGenerator_GetFormat_0(ctx context.Context, marshaler runtime.Marshaler, client EndDeviceQRCodeGeneratorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetQRCodeFormatRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["format_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "format_id")
	}

	protoReq.FormatId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "format_id", err)
	}

	msg, err := client.GetFormat(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_EndDeviceQRCodeGenerator_GetFormat_0(ctx context.Context, marshaler runtime.Marshaler, server EndDeviceQRCodeGeneratorServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetQRCodeFormatRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["format_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "format_id")
	}

	protoReq.FormatId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "format_id", err)
	}

	msg, err := server.GetFormat(ctx, &protoReq)
	return msg, metadata, err

}

func request_EndDeviceQRCodeGenerator_ListFormats_0(ctx context.Context, marshaler runtime.Marshaler, client EndDeviceQRCodeGeneratorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq types.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.ListFormats(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_EndDeviceQRCodeGenerator_ListFormats_0(ctx context.Context, marshaler runtime.Marshaler, server EndDeviceQRCodeGeneratorServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq types.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.ListFormats(ctx, &protoReq)
	return msg, metadata, err

}

func request_EndDeviceQRCodeGenerator_Generate_0(ctx context.Context, marshaler runtime.Marshaler, client EndDeviceQRCodeGeneratorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateEndDeviceQRCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Generate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_EndDeviceQRCodeGenerator_Generate_0(ctx context.Context, marshaler runtime.Marshaler, server EndDeviceQRCodeGeneratorServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateEndDeviceQRCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Generate(ctx, &protoReq)
	return msg, metadata, err

}

func request_QRCodeParser_Parse_0(ctx context.Context, marshaler runtime.Marshaler, client QRCodeParserClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ParseQRCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Parse(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_QRCodeParser_Parse_0(ctx context.Context, marshaler runtime.Marshaler, server QRCodeParserServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ParseQRCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Parse(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterEndDeviceQRCodeGeneratorHandlerServer registers the http handlers for service EndDeviceQRCodeGenerator to "mux".
// UnaryRPC     :call EndDeviceQRCodeGeneratorServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterEndDeviceQRCodeGeneratorHandlerFromEndpoint instead.
func RegisterEndDeviceQRCodeGeneratorHandlerServer(ctx context.Context, mux *runtime.ServeMux, server EndDeviceQRCodeGeneratorServer) error {

	mux.Handle("GET", pattern_EndDeviceQRCodeGenerator_GetFormat_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_EndDeviceQRCodeGenerator_GetFormat_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_GetFormat_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_EndDeviceQRCodeGenerator_ListFormats_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_EndDeviceQRCodeGenerator_ListFormats_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_ListFormats_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_EndDeviceQRCodeGenerator_Generate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_EndDeviceQRCodeGenerator_Generate_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_Generate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterQRCodeParserHandlerServer registers the http handlers for service QRCodeParser to "mux".
// UnaryRPC     :call QRCodeParserServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterQRCodeParserHandlerFromEndpoint instead.
func RegisterQRCodeParserHandlerServer(ctx context.Context, mux *runtime.ServeMux, server QRCodeParserServer) error {

	mux.Handle("POST", pattern_QRCodeParser_Parse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_QRCodeParser_Parse_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_QRCodeParser_Parse_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterEndDeviceQRCodeGeneratorHandlerFromEndpoint is same as RegisterEndDeviceQRCodeGeneratorHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEndDeviceQRCodeGeneratorHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEndDeviceQRCodeGeneratorHandler(ctx, mux, conn)
}

// RegisterEndDeviceQRCodeGeneratorHandler registers the http handlers for service EndDeviceQRCodeGenerator to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEndDeviceQRCodeGeneratorHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEndDeviceQRCodeGeneratorHandlerClient(ctx, mux, NewEndDeviceQRCodeGeneratorClient(conn))
}

// RegisterEndDeviceQRCodeGeneratorHandlerClient registers the http handlers for service EndDeviceQRCodeGenerator
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "EndDeviceQRCodeGeneratorClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "EndDeviceQRCodeGeneratorClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "EndDeviceQRCodeGeneratorClient" to call the correct interceptors.
func RegisterEndDeviceQRCodeGeneratorHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EndDeviceQRCodeGeneratorClient) error {

	mux.Handle("GET", pattern_EndDeviceQRCodeGenerator_GetFormat_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_EndDeviceQRCodeGenerator_GetFormat_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_GetFormat_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_EndDeviceQRCodeGenerator_ListFormats_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_EndDeviceQRCodeGenerator_ListFormats_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_ListFormats_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_EndDeviceQRCodeGenerator_Generate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_EndDeviceQRCodeGenerator_Generate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EndDeviceQRCodeGenerator_Generate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_EndDeviceQRCodeGenerator_GetFormat_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"qr-codes", "end-devices", "formats", "format_id"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_EndDeviceQRCodeGenerator_ListFormats_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"qr-codes", "end-devices", "formats"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_EndDeviceQRCodeGenerator_Generate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"qr-codes", "end-devices"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_EndDeviceQRCodeGenerator_GetFormat_0 = runtime.ForwardResponseMessage

	forward_EndDeviceQRCodeGenerator_ListFormats_0 = runtime.ForwardResponseMessage

	forward_EndDeviceQRCodeGenerator_Generate_0 = runtime.ForwardResponseMessage
)

// RegisterQRCodeParserHandlerFromEndpoint is same as RegisterQRCodeParserHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterQRCodeParserHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterQRCodeParserHandler(ctx, mux, conn)
}

// RegisterQRCodeParserHandler registers the http handlers for service QRCodeParser to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterQRCodeParserHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterQRCodeParserHandlerClient(ctx, mux, NewQRCodeParserClient(conn))
}

// RegisterQRCodeParserHandlerClient registers the http handlers for service QRCodeParser
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "QRCodeParserClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "QRCodeParserClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "QRCodeParserClient" to call the correct interceptors.
func RegisterQRCodeParserHandlerClient(ctx context.Context, mux *runtime.ServeMux, client QRCodeParserClient) error {

	mux.Handle("POST", pattern_QRCodeParser_Parse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_QRCodeParser_Parse_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_QRCodeParser_Parse_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_QRCodeParser_Parse_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"qr-codes"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_QRCodeParser_Parse_0 = runtime.ForwardResponseMessage
)
