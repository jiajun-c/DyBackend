// Code generated by Kitex v0.4.4. DO NOT EDIT.

package thumbupservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	thumbuppart "tiktok/cmd/thumbup/kitex_gen/thumbuppart"
)

func serviceInfo() *kitex.ServiceInfo {
	return thumbupServiceServiceInfo
}

var thumbupServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ThumbupService"
	handlerType := (*thumbuppart.ThumbupService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetFavoriteAction": kitex.NewMethodInfo(getFavoriteActionHandler, newGetFavoriteActionArgs, newGetFavoriteActionResult, false),
		"GetFavoriteList":   kitex.NewMethodInfo(getFavoriteListHandler, newGetFavoriteListArgs, newGetFavoriteListResult, false),
		"GetCommentAction":  kitex.NewMethodInfo(getCommentActionHandler, newGetCommentActionArgs, newGetCommentActionResult, false),
		"GetCommentList":    kitex.NewMethodInfo(getCommentListHandler, newGetCommentListArgs, newGetCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "thumbup",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getFavoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(thumbuppart.FavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(thumbuppart.ThumbupService).GetFavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteActionArgs:
		success, err := handler.(thumbuppart.ThumbupService).GetFavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteActionArgs() interface{} {
	return &GetFavoriteActionArgs{}
}

func newGetFavoriteActionResult() interface{} {
	return &GetFavoriteActionResult{}
}

type GetFavoriteActionArgs struct {
	Req *thumbuppart.FavoriteActionRequest
}

func (p *GetFavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(thumbuppart.FavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(thumbuppart.FavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteActionArgs_Req_DEFAULT *thumbuppart.FavoriteActionRequest

func (p *GetFavoriteActionArgs) GetReq() *thumbuppart.FavoriteActionRequest {
	if !p.IsSetReq() {
		return GetFavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFavoriteActionResult struct {
	Success *thumbuppart.FavoriteActionResponse
}

var GetFavoriteActionResult_Success_DEFAULT *thumbuppart.FavoriteActionResponse

func (p *GetFavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(thumbuppart.FavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(thumbuppart.FavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteActionResult) GetSuccess() *thumbuppart.FavoriteActionResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*thumbuppart.FavoriteActionResponse)
}

func (p *GetFavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(thumbuppart.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(thumbuppart.ThumbupService).GetFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteListArgs:
		success, err := handler.(thumbuppart.ThumbupService).GetFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteListArgs() interface{} {
	return &GetFavoriteListArgs{}
}

func newGetFavoriteListResult() interface{} {
	return &GetFavoriteListResult{}
}

type GetFavoriteListArgs struct {
	Req *thumbuppart.FavoriteListRequest
}

func (p *GetFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(thumbuppart.FavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(thumbuppart.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteListArgs_Req_DEFAULT *thumbuppart.FavoriteListRequest

func (p *GetFavoriteListArgs) GetReq() *thumbuppart.FavoriteListRequest {
	if !p.IsSetReq() {
		return GetFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFavoriteListResult struct {
	Success *thumbuppart.FavoriteListResponse
}

var GetFavoriteListResult_Success_DEFAULT *thumbuppart.FavoriteListResponse

func (p *GetFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(thumbuppart.FavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(thumbuppart.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteListResult) GetSuccess() *thumbuppart.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*thumbuppart.FavoriteListResponse)
}

func (p *GetFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(thumbuppart.CommentActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(thumbuppart.ThumbupService).GetCommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentActionArgs:
		success, err := handler.(thumbuppart.ThumbupService).GetCommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentActionArgs() interface{} {
	return &GetCommentActionArgs{}
}

func newGetCommentActionResult() interface{} {
	return &GetCommentActionResult{}
}

type GetCommentActionArgs struct {
	Req *thumbuppart.CommentActionRequest
}

func (p *GetCommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(thumbuppart.CommentActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentActionArgs) Unmarshal(in []byte) error {
	msg := new(thumbuppart.CommentActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentActionArgs_Req_DEFAULT *thumbuppart.CommentActionRequest

func (p *GetCommentActionArgs) GetReq() *thumbuppart.CommentActionRequest {
	if !p.IsSetReq() {
		return GetCommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentActionResult struct {
	Success *thumbuppart.CommentActionResponse
}

var GetCommentActionResult_Success_DEFAULT *thumbuppart.CommentActionResponse

func (p *GetCommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(thumbuppart.CommentActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentActionResult) Unmarshal(in []byte) error {
	msg := new(thumbuppart.CommentActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentActionResult) GetSuccess() *thumbuppart.CommentActionResponse {
	if !p.IsSetSuccess() {
		return GetCommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*thumbuppart.CommentActionResponse)
}

func (p *GetCommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(thumbuppart.CommentListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(thumbuppart.ThumbupService).GetCommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentListArgs:
		success, err := handler.(thumbuppart.ThumbupService).GetCommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentListResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentListArgs() interface{} {
	return &GetCommentListArgs{}
}

func newGetCommentListResult() interface{} {
	return &GetCommentListResult{}
}

type GetCommentListArgs struct {
	Req *thumbuppart.CommentListRequest
}

func (p *GetCommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(thumbuppart.CommentListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentListArgs) Unmarshal(in []byte) error {
	msg := new(thumbuppart.CommentListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentListArgs_Req_DEFAULT *thumbuppart.CommentListRequest

func (p *GetCommentListArgs) GetReq() *thumbuppart.CommentListRequest {
	if !p.IsSetReq() {
		return GetCommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentListResult struct {
	Success *thumbuppart.CommentListResponse
}

var GetCommentListResult_Success_DEFAULT *thumbuppart.CommentListResponse

func (p *GetCommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(thumbuppart.CommentListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentListResult) Unmarshal(in []byte) error {
	msg := new(thumbuppart.CommentListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentListResult) GetSuccess() *thumbuppart.CommentListResponse {
	if !p.IsSetSuccess() {
		return GetCommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*thumbuppart.CommentListResponse)
}

func (p *GetCommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetFavoriteAction(ctx context.Context, Req *thumbuppart.FavoriteActionRequest) (r *thumbuppart.FavoriteActionResponse, err error) {
	var _args GetFavoriteActionArgs
	_args.Req = Req
	var _result GetFavoriteActionResult
	if err = p.c.Call(ctx, "GetFavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, Req *thumbuppart.FavoriteListRequest) (r *thumbuppart.FavoriteListResponse, err error) {
	var _args GetFavoriteListArgs
	_args.Req = Req
	var _result GetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentAction(ctx context.Context, Req *thumbuppart.CommentActionRequest) (r *thumbuppart.CommentActionResponse, err error) {
	var _args GetCommentActionArgs
	_args.Req = Req
	var _result GetCommentActionResult
	if err = p.c.Call(ctx, "GetCommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, Req *thumbuppart.CommentListRequest) (r *thumbuppart.CommentListResponse, err error) {
	var _args GetCommentListArgs
	_args.Req = Req
	var _result GetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
