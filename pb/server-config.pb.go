// Code generated by protoc-gen-go.
// source: server-config.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	server-config.proto
	mutation.proto
	api.proto

It has these top-level messages:
	ServerConfig
	AccessPolicy
	Account
	Mutation
	LogTreeHash
	LeafNode
	TreeNode
	AccountRef
	LogRef
	MapRef
	LogCreateRequest
	LogCreateResponse
	LogDeleteRequest
	LogDeleteResponse
	LogListRequest
	LogInfo
	LogListResponse
	MapCreateRequest
	MapCreateResponse
	MapDeleteRequest
	MapDeleteResponse
	MapListRequest
	MapInfo
	MapListResponse
	LogTreeHashRequest
	LogTreeHashResponse
	MapTreeHashRequest
	MapTreeHashResponse
	LogInclusionProofRequest
	LogInclusionProofResponse
	LogConsistencyProofRequest
	LogConsistencyProofResponse
	HashableData
	LogAddEntryRequest
	LogAddEntryResponse
	MapSetValueRequest
	MapSetValueResponse
	MapGetValueRequest
	MapGetValueResponse
	LogFetchEntriesRequest
	LogFetchEntriesResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Permission int32

const (
	Permission_PERM_NONE                    Permission = 0
	Permission_PERM_ALL_PERMISSIONS         Permission = 1
	Permission_PERM_ACCOUNT_LIST_LOGS       Permission = 2
	Permission_PERM_ACCOUNT_LIST_MAPS       Permission = 3
	Permission_PERM_LOG_CREATE              Permission = 4
	Permission_PERM_LOG_DELETE              Permission = 5
	Permission_PERM_MAP_CREATE              Permission = 6
	Permission_PERM_MAP_DELETE              Permission = 7
	Permission_PERM_LOG_RAW_ADD             Permission = 8
	Permission_PERM_LOG_READ_ENTRY          Permission = 9
	Permission_PERM_LOG_READ_HASH           Permission = 10
	Permission_PERM_LOG_PROVE_INCLUSION     Permission = 11
	Permission_PERM_MAP_SET_VALUE           Permission = 12
	Permission_PERM_MAP_GET_VALUE           Permission = 13
	Permission_PERM_MAP_MUTATION_READ_ENTRY Permission = 14
	Permission_PERM_MAP_MUTATION_READ_HASH  Permission = 15
)

var Permission_name = map[int32]string{
	0:  "PERM_NONE",
	1:  "PERM_ALL_PERMISSIONS",
	2:  "PERM_ACCOUNT_LIST_LOGS",
	3:  "PERM_ACCOUNT_LIST_MAPS",
	4:  "PERM_LOG_CREATE",
	5:  "PERM_LOG_DELETE",
	6:  "PERM_MAP_CREATE",
	7:  "PERM_MAP_DELETE",
	8:  "PERM_LOG_RAW_ADD",
	9:  "PERM_LOG_READ_ENTRY",
	10: "PERM_LOG_READ_HASH",
	11: "PERM_LOG_PROVE_INCLUSION",
	12: "PERM_MAP_SET_VALUE",
	13: "PERM_MAP_GET_VALUE",
	14: "PERM_MAP_MUTATION_READ_ENTRY",
	15: "PERM_MAP_MUTATION_READ_HASH",
}
var Permission_value = map[string]int32{
	"PERM_NONE":                    0,
	"PERM_ALL_PERMISSIONS":         1,
	"PERM_ACCOUNT_LIST_LOGS":       2,
	"PERM_ACCOUNT_LIST_MAPS":       3,
	"PERM_LOG_CREATE":              4,
	"PERM_LOG_DELETE":              5,
	"PERM_MAP_CREATE":              6,
	"PERM_MAP_DELETE":              7,
	"PERM_LOG_RAW_ADD":             8,
	"PERM_LOG_READ_ENTRY":          9,
	"PERM_LOG_READ_HASH":           10,
	"PERM_LOG_PROVE_INCLUSION":     11,
	"PERM_MAP_SET_VALUE":           12,
	"PERM_MAP_GET_VALUE":           13,
	"PERM_MAP_MUTATION_READ_ENTRY": 14,
	"PERM_MAP_MUTATION_READ_HASH":  15,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}
func (Permission) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerConfig struct {
	ServerCertPath               string     `protobuf:"bytes,1,opt,name=server_cert_path,json=serverCertPath" json:"server_cert_path,omitempty"`
	ServerKeyPath                string     `protobuf:"bytes,2,opt,name=server_key_path,json=serverKeyPath" json:"server_key_path,omitempty"`
	ListenBind                   string     `protobuf:"bytes,3,opt,name=listen_bind,json=listenBind" json:"listen_bind,omitempty"`
	InsecureHttpServerForTesting bool       `protobuf:"varint,4,opt,name=insecure_http_server_for_testing,json=insecureHttpServerForTesting" json:"insecure_http_server_for_testing,omitempty"`
	Accounts                     []*Account `protobuf:"bytes,5,rep,name=accounts" json:"accounts,omitempty"`
	BoltDbPath                   string     `protobuf:"bytes,6,opt,name=bolt_db_path,json=boltDbPath" json:"bolt_db_path,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerConfig) GetServerCertPath() string {
	if m != nil {
		return m.ServerCertPath
	}
	return ""
}

func (m *ServerConfig) GetServerKeyPath() string {
	if m != nil {
		return m.ServerKeyPath
	}
	return ""
}

func (m *ServerConfig) GetListenBind() string {
	if m != nil {
		return m.ListenBind
	}
	return ""
}

func (m *ServerConfig) GetInsecureHttpServerForTesting() bool {
	if m != nil {
		return m.InsecureHttpServerForTesting
	}
	return false
}

func (m *ServerConfig) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetBoltDbPath() string {
	if m != nil {
		return m.BoltDbPath
	}
	return ""
}

type AccessPolicy struct {
	ApiKey        string       `protobuf:"bytes,1,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	NameMatch     string       `protobuf:"bytes,2,opt,name=name_match,json=nameMatch" json:"name_match,omitempty"`
	AllowedFields []string     `protobuf:"bytes,3,rep,name=allowed_fields,json=allowedFields" json:"allowed_fields,omitempty"`
	Permissions   []Permission `protobuf:"varint,4,rep,packed,name=permissions,enum=continusec.verifiabledatastructures.server.Permission" json:"permissions,omitempty"`
}

func (m *AccessPolicy) Reset()                    { *m = AccessPolicy{} }
func (m *AccessPolicy) String() string            { return proto.CompactTextString(m) }
func (*AccessPolicy) ProtoMessage()               {}
func (*AccessPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccessPolicy) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *AccessPolicy) GetNameMatch() string {
	if m != nil {
		return m.NameMatch
	}
	return ""
}

func (m *AccessPolicy) GetAllowedFields() []string {
	if m != nil {
		return m.AllowedFields
	}
	return nil
}

func (m *AccessPolicy) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Account struct {
	Id     string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Policy []*AccessPolicy `protobuf:"bytes,2,rep,name=policy" json:"policy,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetPolicy() []*AccessPolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerConfig)(nil), "continusec.verifiabledatastructures.server.ServerConfig")
	proto.RegisterType((*AccessPolicy)(nil), "continusec.verifiabledatastructures.server.AccessPolicy")
	proto.RegisterType((*Account)(nil), "continusec.verifiabledatastructures.server.Account")
	proto.RegisterEnum("continusec.verifiabledatastructures.server.Permission", Permission_name, Permission_value)
}

func init() { proto.RegisterFile("server-config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x4e, 0xdb, 0x4e,
	0x14, 0xc5, 0xff, 0xb1, 0x43, 0x20, 0x37, 0x21, 0x8c, 0x06, 0x04, 0xd6, 0xbf, 0x54, 0x58, 0x48,
	0xad, 0x22, 0xa4, 0x66, 0x01, 0x52, 0xd5, 0xad, 0x49, 0x0c, 0x44, 0x38, 0xb6, 0x65, 0x3b, 0xf4,
	0x63, 0x33, 0xf2, 0xc7, 0x04, 0x46, 0x18, 0xdb, 0xf2, 0x4c, 0xa8, 0xf2, 0x28, 0x7d, 0x97, 0xee,
	0xfb, 0x5a, 0x95, 0x3f, 0x30, 0x44, 0x6a, 0xa5, 0xb2, 0x9b, 0xfc, 0xee, 0xb9, 0x37, 0xe7, 0x9c,
	0x85, 0x61, 0x97, 0xd3, 0xfc, 0x91, 0xe6, 0x1f, 0xc2, 0x34, 0x59, 0xb0, 0xdb, 0x51, 0x96, 0xa7,
	0x22, 0xc5, 0x27, 0x61, 0x9a, 0x08, 0x96, 0x2c, 0x39, 0x0d, 0x47, 0x8f, 0x34, 0x67, 0x0b, 0xe6,
	0x07, 0x31, 0x8d, 0x7c, 0xe1, 0x73, 0x91, 0x2f, 0x43, 0xb1, 0xcc, 0x29, 0x1f, 0x55, 0x8b, 0xc7,
	0x3f, 0x25, 0xe8, 0xbb, 0xe5, 0x73, 0x5c, 0x9e, 0xc0, 0x43, 0x40, 0xd5, 0x88, 0x84, 0x34, 0x17,
	0x24, 0xf3, 0xc5, 0x9d, 0xd2, 0x52, 0x5b, 0xc3, 0xae, 0x33, 0xa8, 0xf8, 0x98, 0xe6, 0xc2, 0xf6,
	0xc5, 0x1d, 0x7e, 0x0f, 0x3b, 0xb5, 0xf2, 0x9e, 0xae, 0x2a, 0xa1, 0x54, 0x0a, 0xb7, 0x2b, 0x7c,
	0x4d, 0x57, 0xa5, 0xee, 0x08, 0x7a, 0x31, 0xe3, 0x82, 0x26, 0x24, 0x60, 0x49, 0xa4, 0xc8, 0xa5,
	0x06, 0x2a, 0x74, 0xce, 0x92, 0x08, 0x5f, 0x80, 0xca, 0x12, 0x4e, 0xc3, 0x65, 0x4e, 0xc9, 0x9d,
	0x10, 0x19, 0xa9, 0xcf, 0x2e, 0xd2, 0x9c, 0x08, 0xca, 0x05, 0x4b, 0x6e, 0x95, 0xb6, 0xda, 0x1a,
	0x6e, 0x39, 0x87, 0x4f, 0xba, 0x2b, 0x21, 0xb2, 0xca, 0xf6, 0x45, 0x9a, 0x7b, 0x95, 0x06, 0x5b,
	0xb0, 0xe5, 0x87, 0x61, 0xba, 0x4c, 0x04, 0x57, 0x36, 0x54, 0x79, 0xd8, 0x3b, 0x3d, 0x1b, 0xfd,
	0x7b, 0x15, 0x23, 0xad, 0xda, 0x75, 0x9a, 0x23, 0x58, 0x85, 0x7e, 0x90, 0xc6, 0x82, 0x44, 0x41,
	0x15, 0xaf, 0x53, 0x59, 0x2f, 0xd8, 0x24, 0x28, 0xb2, 0x1d, 0xff, 0x6a, 0x41, 0x5f, 0x0b, 0x43,
	0xca, 0xb9, 0x9d, 0xc6, 0x2c, 0x5c, 0xe1, 0x03, 0xd8, 0xf4, 0x33, 0x56, 0x34, 0x52, 0xb7, 0xd6,
	0xf1, 0x33, 0x76, 0x4d, 0x57, 0xf8, 0x2d, 0x40, 0xe2, 0x3f, 0x50, 0xf2, 0xe0, 0x8b, 0xf0, 0xa9,
	0xa8, 0x6e, 0x41, 0x66, 0x05, 0xc0, 0xef, 0x60, 0xe0, 0xc7, 0x71, 0xfa, 0x9d, 0x46, 0x64, 0xc1,
	0x68, 0x1c, 0x71, 0x45, 0x56, 0xe5, 0xa2, 0xcb, 0x9a, 0x5e, 0x94, 0x10, 0x7f, 0x81, 0x5e, 0x46,
	0xf3, 0x07, 0xc6, 0x39, 0x4b, 0x13, 0xae, 0xb4, 0x55, 0x79, 0x38, 0x38, 0xfd, 0xf8, 0x9a, 0x94,
	0x76, 0xb3, 0xee, 0xbc, 0x3c, 0x75, 0x7c, 0x0f, 0x9b, 0x75, 0x01, 0x78, 0x00, 0x12, 0x8b, 0x6a,
	0xfb, 0x12, 0x8b, 0xb0, 0x0d, 0x9d, 0xac, 0x4c, 0xa7, 0x48, 0x65, 0xab, 0x9f, 0x5e, 0xd9, 0x6a,
	0xd3, 0x8e, 0x53, 0xdf, 0x39, 0xf9, 0x21, 0x03, 0x3c, 0x1b, 0xc1, 0xdb, 0xd0, 0xb5, 0x75, 0x67,
	0x46, 0x4c, 0xcb, 0xd4, 0xd1, 0x7f, 0x58, 0x81, 0xbd, 0xf2, 0xa7, 0x66, 0x18, 0xa4, 0x78, 0x4c,
	0x5d, 0x77, 0x6a, 0x99, 0x2e, 0x6a, 0xe1, 0xff, 0x61, 0xbf, 0x9a, 0x8c, 0xc7, 0xd6, 0xdc, 0xf4,
	0x88, 0x31, 0x75, 0x3d, 0x62, 0x58, 0x97, 0x2e, 0x92, 0xfe, 0x3c, 0x9b, 0x69, 0xb6, 0x8b, 0x64,
	0xbc, 0x0b, 0x3b, 0xe5, 0xcc, 0xb0, 0x2e, 0xc9, 0xd8, 0xd1, 0x35, 0x4f, 0x47, 0xed, 0x35, 0x38,
	0xd1, 0x0d, 0xdd, 0xd3, 0xd1, 0x46, 0x03, 0x67, 0x9a, 0xfd, 0xa4, 0xec, 0xac, 0xc1, 0x5a, 0xb9,
	0x89, 0xf7, 0x00, 0x35, 0xeb, 0x8e, 0xf6, 0x99, 0x68, 0x93, 0x09, 0xda, 0xc2, 0x07, 0xb0, 0xfb,
	0x4c, 0x75, 0x6d, 0x42, 0x74, 0xd3, 0x73, 0xbe, 0xa2, 0x2e, 0xde, 0x07, 0xbc, 0x3e, 0xb8, 0xd2,
	0xdc, 0x2b, 0x04, 0xf8, 0x10, 0x94, 0x86, 0xdb, 0x8e, 0x75, 0xa3, 0x93, 0xa9, 0x39, 0x36, 0xe6,
	0x45, 0x62, 0xd4, 0x6b, 0xb6, 0x8a, 0x7f, 0x76, 0x75, 0x8f, 0xdc, 0x68, 0xc6, 0x5c, 0x47, 0xfd,
	0x35, 0x7e, 0xd9, 0xf0, 0x6d, 0xac, 0xc2, 0x61, 0xc3, 0x67, 0x73, 0x4f, 0xf3, 0xa6, 0x96, 0xf9,
	0xd2, 0xc7, 0x00, 0x1f, 0xc1, 0x9b, 0xbf, 0x28, 0x4a, 0x43, 0x3b, 0xe7, 0xed, 0x6f, 0x52, 0x16,
	0x04, 0x9d, 0xf2, 0x53, 0x72, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x19, 0x41, 0x42, 0x61,
	0x04, 0x00, 0x00,
}
