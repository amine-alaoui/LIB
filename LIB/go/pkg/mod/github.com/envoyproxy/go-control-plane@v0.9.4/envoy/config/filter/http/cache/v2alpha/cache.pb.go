// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/cache/v2alpha/cache.proto

package envoy_config_filter_http_cache_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CacheConfig struct {
	TypedConfig          *any.Any                      `protobuf:"bytes,1,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	AllowedVaryHeaders   []*matcher.StringMatcher      `protobuf:"bytes,2,rep,name=allowed_vary_headers,json=allowedVaryHeaders,proto3" json:"allowed_vary_headers,omitempty"`
	KeyCreatorParams     *CacheConfig_KeyCreatorParams `protobuf:"bytes,3,opt,name=key_creator_params,json=keyCreatorParams,proto3" json:"key_creator_params,omitempty"`
	MaxBodyBytes         uint32                        `protobuf:"varint,4,opt,name=max_body_bytes,json=maxBodyBytes,proto3" json:"max_body_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CacheConfig) Reset()         { *m = CacheConfig{} }
func (m *CacheConfig) String() string { return proto.CompactTextString(m) }
func (*CacheConfig) ProtoMessage()    {}
func (*CacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d6a0b399e44d47, []int{0}
}

func (m *CacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig.Unmarshal(m, b)
}
func (m *CacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig.Marshal(b, m, deterministic)
}
func (m *CacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig.Merge(m, src)
}
func (m *CacheConfig) XXX_Size() int {
	return xxx_messageInfo_CacheConfig.Size(m)
}
func (m *CacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig proto.InternalMessageInfo

func (m *CacheConfig) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func (m *CacheConfig) GetAllowedVaryHeaders() []*matcher.StringMatcher {
	if m != nil {
		return m.AllowedVaryHeaders
	}
	return nil
}

func (m *CacheConfig) GetKeyCreatorParams() *CacheConfig_KeyCreatorParams {
	if m != nil {
		return m.KeyCreatorParams
	}
	return nil
}

func (m *CacheConfig) GetMaxBodyBytes() uint32 {
	if m != nil {
		return m.MaxBodyBytes
	}
	return 0
}

type CacheConfig_KeyCreatorParams struct {
	ExcludeScheme           bool                           `protobuf:"varint,1,opt,name=exclude_scheme,json=excludeScheme,proto3" json:"exclude_scheme,omitempty"`
	ExcludeHost             bool                           `protobuf:"varint,2,opt,name=exclude_host,json=excludeHost,proto3" json:"exclude_host,omitempty"`
	QueryParametersIncluded []*route.QueryParameterMatcher `protobuf:"bytes,3,rep,name=query_parameters_included,json=queryParametersIncluded,proto3" json:"query_parameters_included,omitempty"`
	QueryParametersExcluded []*route.QueryParameterMatcher `protobuf:"bytes,4,rep,name=query_parameters_excluded,json=queryParametersExcluded,proto3" json:"query_parameters_excluded,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                       `json:"-"`
	XXX_unrecognized        []byte                         `json:"-"`
	XXX_sizecache           int32                          `json:"-"`
}

func (m *CacheConfig_KeyCreatorParams) Reset()         { *m = CacheConfig_KeyCreatorParams{} }
func (m *CacheConfig_KeyCreatorParams) String() string { return proto.CompactTextString(m) }
func (*CacheConfig_KeyCreatorParams) ProtoMessage()    {}
func (*CacheConfig_KeyCreatorParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d6a0b399e44d47, []int{0, 0}
}

func (m *CacheConfig_KeyCreatorParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Unmarshal(m, b)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Marshal(b, m, deterministic)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.Merge(m, src)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Size() int {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Size(m)
}
func (m *CacheConfig_KeyCreatorParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig_KeyCreatorParams proto.InternalMessageInfo

func (m *CacheConfig_KeyCreatorParams) GetExcludeScheme() bool {
	if m != nil {
		return m.ExcludeScheme
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetExcludeHost() bool {
	if m != nil {
		return m.ExcludeHost
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersIncluded() []*route.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersIncluded
	}
	return nil
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersExcluded() []*route.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersExcluded
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheConfig)(nil), "envoy.config.filter.http.cache.v2alpha.CacheConfig")
	proto.RegisterType((*CacheConfig_KeyCreatorParams)(nil), "envoy.config.filter.http.cache.v2alpha.CacheConfig.KeyCreatorParams")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/cache/v2alpha/cache.proto", fileDescriptor_f8d6a0b399e44d47)
}

var fileDescriptor_f8d6a0b399e44d47 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6b, 0xdb, 0x4c,
	0x10, 0x45, 0x4e, 0xf2, 0x61, 0xd6, 0x4e, 0x08, 0x22, 0x10, 0xc5, 0x87, 0xaf, 0x4e, 0x69, 0x8b,
	0x43, 0xe9, 0x2e, 0x28, 0xfd, 0x03, 0x91, 0x5b, 0x48, 0x29, 0x05, 0xd7, 0x86, 0x5e, 0xc5, 0x58,
	0x1a, 0x5b, 0x22, 0xd2, 0xae, 0xb2, 0xbb, 0x56, 0xbc, 0xbf, 0xa2, 0xd7, 0x9e, 0x7b, 0xec, 0x6f,
	0xeb, 0xa9, 0xc7, 0x1e, 0x4a, 0xd1, 0xee, 0x1a, 0x9a, 0xd0, 0x43, 0xa0, 0x17, 0xa1, 0x9d, 0x99,
	0xf7, 0x66, 0xde, 0xbc, 0x21, 0x31, 0xf2, 0x56, 0x18, 0x96, 0x09, 0xbe, 0x2a, 0xd7, 0x6c, 0x55,
	0x56, 0x1a, 0x25, 0x2b, 0xb4, 0x6e, 0x58, 0x06, 0x59, 0x81, 0xac, 0x8d, 0xa1, 0x6a, 0x0a, 0x70,
	0x2f, 0xda, 0x48, 0xa1, 0x45, 0xf8, 0xc2, 0x62, 0xa8, 0xc3, 0x50, 0x87, 0xa1, 0x1d, 0x86, 0xba,
	0x2a, 0x8f, 0x19, 0x5d, 0x38, 0x6e, 0x68, 0x4a, 0xd6, 0xc6, 0x4c, 0x8a, 0x8d, 0x46, 0xf7, 0x4d,
	0x33, 0x51, 0x37, 0x82, 0x23, 0xd7, 0xca, 0x51, 0x8e, 0x9e, 0xb8, 0x52, 0x6d, 0x1a, 0x64, 0x35,
	0xe8, 0xac, 0x40, 0xc9, 0x94, 0x96, 0x25, 0x5f, 0xfb, 0x82, 0xb3, 0xb5, 0x10, 0xeb, 0x0a, 0x99,
	0x7d, 0x2d, 0x37, 0x2b, 0x06, 0xdc, 0xf8, 0xd4, 0xff, 0x9b, 0xbc, 0x01, 0x06, 0x9c, 0x0b, 0x0d,
	0xba, 0x14, 0x5c, 0xb1, 0xba, 0x5c, 0x4b, 0xd0, 0x7e, 0xdc, 0xd1, 0x69, 0x0b, 0x55, 0x99, 0x83,
	0x46, 0xb6, 0xfb, 0x71, 0x89, 0xa7, 0xdf, 0xf7, 0xc9, 0x60, 0xda, 0x4d, 0x3c, 0xb5, 0x4a, 0xc2,
	0x2b, 0x32, 0xec, 0x06, 0xc8, 0x53, 0xa7, 0x2c, 0x0a, 0xc6, 0xc1, 0x64, 0x10, 0x9f, 0x50, 0xd7,
	0x9a, 0xee, 0x5a, 0xd3, 0x2b, 0x6e, 0x92, 0xfe, 0xcf, 0xe4, 0xe0, 0x6b, 0xd0, 0xeb, 0x07, 0xf3,
	0x81, 0xc5, 0x78, 0x8a, 0x05, 0x39, 0x81, 0xaa, 0x12, 0x77, 0x98, 0xa7, 0x2d, 0x48, 0x93, 0x16,
	0x08, 0x39, 0x4a, 0x15, 0xf5, 0xc6, 0x7b, 0x93, 0x41, 0x7c, 0x4e, 0xdd, 0xe6, 0x3a, 0x04, 0xf5,
	0x32, 0xe9, 0xc2, 0xca, 0xfc, 0xe0, 0x5e, 0xf3, 0xd0, 0xc3, 0x3f, 0x81, 0x34, 0xd7, 0x0e, 0x1c,
	0x4a, 0x12, 0xde, 0xa0, 0x49, 0x33, 0x89, 0xa0, 0x85, 0x4c, 0x1b, 0x90, 0x50, 0xab, 0x68, 0xcf,
	0x4e, 0xf7, 0x86, 0x3e, 0xce, 0x0c, 0xfa, 0x87, 0x50, 0xfa, 0x1e, 0xcd, 0xd4, 0x91, 0xcd, 0x2c,
	0xd7, 0xfc, 0xf8, 0xe6, 0x41, 0x24, 0x7c, 0x46, 0x8e, 0x6a, 0xd8, 0xa6, 0x4b, 0x91, 0x9b, 0x74,
	0x69, 0x34, 0xaa, 0x68, 0x7f, 0x1c, 0x4c, 0x0e, 0xe7, 0xc3, 0x1a, 0xb6, 0x89, 0xc8, 0x4d, 0xd2,
	0xc5, 0x46, 0xdf, 0x7a, 0xe4, 0xf8, 0x21, 0x59, 0xf8, 0x9c, 0x1c, 0xe1, 0x36, 0xab, 0x36, 0x39,
	0xa6, 0x2a, 0x2b, 0xb0, 0x46, 0xbb, 0xc8, 0xfe, 0xfc, 0xd0, 0x47, 0x17, 0x36, 0x18, 0x9e, 0x93,
	0xe1, 0xae, 0xac, 0x10, 0x4a, 0x47, 0x3d, 0x5b, 0x34, 0xf0, 0xb1, 0x6b, 0xa1, 0x74, 0x88, 0xe4,
	0xec, 0x76, 0x83, 0xd2, 0x38, 0xc9, 0xa8, 0x51, 0xaa, 0xb4, 0xe4, 0x36, 0x9f, 0x47, 0x7b, 0x76,
	0xa5, 0x17, 0x5e, 0x3f, 0x34, 0x25, 0x6d, 0x63, 0x6a, 0xcf, 0x8b, 0x7e, 0xec, 0x40, 0xb3, 0x1d,
	0x66, 0xb7, 0xda, 0xd3, 0xdb, 0x7b, 0x61, 0xf5, 0xce, 0x33, 0xfd, 0xb5, 0x8d, 0x1f, 0x23, 0x8f,
	0xf6, 0xff, 0xb5, 0xcd, 0x5b, 0xcf, 0x94, 0xdc, 0xfd, 0xf8, 0xf2, 0xeb, 0xf3, 0xc1, 0xab, 0xf0,
	0xa5, 0xa3, 0xc2, 0xad, 0x46, 0xae, 0xba, 0x7b, 0xf5, 0xae, 0xa9, 0x7b, 0xb6, 0x5d, 0x5a, 0xdb,
	0xc8, 0xeb, 0x52, 0xb8, 0xd6, 0x8d, 0x14, 0x5b, 0xf3, 0x48, 0xb3, 0x13, 0x62, 0xdd, 0x9e, 0x75,
	0x07, 0x3b, 0x0b, 0x96, 0xff, 0xd9, 0xcb, 0xbd, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xa7,
	0xa7, 0x38, 0xec, 0x03, 0x00, 0x00,
}
