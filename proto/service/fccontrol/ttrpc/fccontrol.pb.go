// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fccontrol.proto

package fccontrol

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto1 "github.com/firecracker-microvm/firecracker-containerd/proto"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("fccontrol.proto", fileDescriptor_b99f53e2bf82c5ef) }

var fileDescriptor_b99f53e2bf82c5ef = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4b, 0x4f, 0xfa, 0x50,
	0x10, 0xc5, 0xe9, 0xe2, 0xcf, 0x63, 0xfe, 0xf2, 0xba, 0x51, 0x62, 0x30, 0xe9, 0xc6, 0xad, 0x19,
	0x0c, 0xba, 0x34, 0xc4, 0xf8, 0x22, 0x26, 0x36, 0x1a, 0x88, 0x2c, 0xdc, 0x5d, 0xda, 0x29, 0x1a,
	0xa1, 0x53, 0xdb, 0xdb, 0x85, 0x3b, 0xe3, 0xa7, 0x63, 0xe9, 0xd2, 0xa5, 0xf4, 0x93, 0x18, 0xe8,
	0x03, 0x68, 0x4c, 0xba, 0x3b, 0xf3, 0x6b, 0xcf, 0x9c, 0xd3, 0x7b, 0x0b, 0x75, 0xdb, 0x34, 0xd9,
	0x51, 0x1e, 0x4f, 0xd1, 0xf5, 0x58, 0x71, 0xfb, 0x60, 0xc2, 0x3c, 0x99, 0x52, 0x67, 0x35, 0x8d,
	0x03, 0xbb, 0x43, 0x33, 0x57, 0xbd, 0xc7, 0x0f, 0x9b, 0xf6, 0x8b, 0x47, 0xa6, 0x27, 0xcd, 0x57,
	0xf2, 0x22, 0xd4, 0xfd, 0xfc, 0x07, 0xff, 0x6f, 0xd6, 0x54, 0x74, 0xa0, 0x7c, 0xe9, 0x91, 0x54,
	0x34, 0x32, 0x44, 0x03, 0x13, 0x39, 0xa0, 0xb7, 0x80, 0x7c, 0xd5, 0x6e, 0x6e, 0x10, 0xdf, 0x65,
	0xc7, 0x27, 0x71, 0x0c, 0xc5, 0xa1, 0x62, 0x77, 0x64, 0x88, 0x1a, 0x46, 0x22, 0x79, 0xb9, 0x85,
	0x51, 0x17, 0x4c, 0xba, 0xe0, 0xf5, 0xb2, 0x8b, 0xe8, 0x42, 0xa5, 0x4f, 0x6a, 0x64, 0xdc, 0x3a,
	0x36, 0x8b, 0x26, 0xa6, 0x3a, 0xf1, 0x89, 0x4d, 0x14, 0xa7, 0xf4, 0xa0, 0x3a, 0x5c, 0x42, 0x83,
	0x94, 0xb4, 0xa4, 0x92, 0x62, 0x0f, 0xb7, 0xe6, 0xbc, 0xcc, 0x2b, 0x68, 0x3c, 0xba, 0xd6, 0xaa,
	0x79, 0xba, 0x62, 0x1f, 0xb3, 0x28, 0x6f, 0x4b, 0x0f, 0xaa, 0xfd, 0x4c, 0x8b, 0xfe, 0xdf, 0x2d,
	0x32, 0x38, 0xfe, 0x8a, 0x2e, 0x94, 0x1e, 0x64, 0xe0, 0x2f, 0xcf, 0xb6, 0x8e, 0xb1, 0xca, 0xcb,
	0x3c, 0x85, 0xf2, 0x80, 0xfc, 0x60, 0x16, 0x5d, 0x48, 0x22, 0xf3, 0x5c, 0x67, 0xb0, 0x73, 0xc7,
	0xd2, 0x1a, 0x3a, 0xd2, 0xf5, 0x9f, 0x59, 0x89, 0x5d, 0xdc, 0x1c, 0xf3, 0xdc, 0xe7, 0x50, 0x8b,
	0xee, 0x39, 0xf5, 0xb7, 0x70, 0x1b, 0xe4, 0x6d, 0x38, 0x82, 0xd2, 0xbd, 0x6d, 0x4f, 0x59, 0x5a,
	0xa2, 0x8e, 0xb1, 0x4a, 0x3c, 0x8d, 0x35, 0x88, 0xce, 0xe5, 0xe2, 0x70, 0xbe, 0xd0, 0x0b, 0xdf,
	0x0b, 0xbd, 0xf0, 0x11, 0xea, 0xda, 0x3c, 0xd4, 0xb5, 0xaf, 0x50, 0xd7, 0x7e, 0x42, 0x5d, 0x7b,
	0xaa, 0xa4, 0xff, 0xf7, 0xb8, 0xb8, 0x8a, 0x38, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x39,
	0x1c, 0x7a, 0xf3, 0x02, 0x00, 0x00,
}

type FirecrackerService interface {
	CreateVM(ctx context.Context, req *proto1.CreateVMRequest) (*proto1.CreateVMResponse, error)
	StopVM(ctx context.Context, req *proto1.StopVMRequest) (*empty.Empty, error)
	GetVMInfo(ctx context.Context, req *proto1.GetVMInfoRequest) (*proto1.GetVMInfoResponse, error)
	SetVMMetadata(ctx context.Context, req *proto1.SetVMMetadataRequest) (*empty.Empty, error)
	UpdateVMMetadata(ctx context.Context, req *proto1.UpdateVMMetadataRequest) (*empty.Empty, error)
	GetVMMetadata(ctx context.Context, req *proto1.GetVMMetadataRequest) (*proto1.GetVMMetadataResponse, error)
	PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error)
	ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error)
	LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*empty.Empty, error)
	CreateSnapshot(ctx context.Context, req *proto1.CreateSnapshotRequest) (*empty.Empty, error)
	Offload(ctx context.Context, req *proto1.OffloadRequest) (*proto1.OffloadResponse, error)
}

func RegisterFirecrackerService(srv *github_com_containerd_ttrpc.Server, svc FirecrackerService) {
	srv.Register("Firecracker", map[string]github_com_containerd_ttrpc.Method{
		"CreateVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.CreateVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.CreateVM(ctx, &req)
		},
		"StopVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.StopVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.StopVM(ctx, &req)
		},
		"GetVMInfo": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetVMInfoRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetVMInfo(ctx, &req)
		},
		"SetVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.SetVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.SetVMMetadata(ctx, &req)
		},
		"UpdateVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.UpdateVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.UpdateVMMetadata(ctx, &req)
		},
		"GetVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetVMMetadata(ctx, &req)
		},
		"PauseVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.PauseVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.PauseVM(ctx, &req)
		},
		"ResumeVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.ResumeVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.ResumeVM(ctx, &req)
		},
		"LoadSnapshot": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.LoadSnapshotRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.LoadSnapshot(ctx, &req)
		},
		"CreateSnapshot": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.CreateSnapshotRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.CreateSnapshot(ctx, &req)
		},
		"Offload": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.OffloadRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Offload(ctx, &req)
		},
	})
}

type firecrackerClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewFirecrackerClient(client *github_com_containerd_ttrpc.Client) FirecrackerService {
	return &firecrackerClient{
		client: client,
	}
}

func (c *firecrackerClient) CreateVM(ctx context.Context, req *proto1.CreateVMRequest) (*proto1.CreateVMResponse, error) {
	var resp proto1.CreateVMResponse
	if err := c.client.Call(ctx, "Firecracker", "CreateVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) StopVM(ctx context.Context, req *proto1.StopVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "StopVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) GetVMInfo(ctx context.Context, req *proto1.GetVMInfoRequest) (*proto1.GetVMInfoResponse, error) {
	var resp proto1.GetVMInfoResponse
	if err := c.client.Call(ctx, "Firecracker", "GetVMInfo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) SetVMMetadata(ctx context.Context, req *proto1.SetVMMetadataRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "SetVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) UpdateVMMetadata(ctx context.Context, req *proto1.UpdateVMMetadataRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "UpdateVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) GetVMMetadata(ctx context.Context, req *proto1.GetVMMetadataRequest) (*proto1.GetVMMetadataResponse, error) {
	var resp proto1.GetVMMetadataResponse
	if err := c.client.Call(ctx, "Firecracker", "GetVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "PauseVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "ResumeVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "LoadSnapshot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) CreateSnapshot(ctx context.Context, req *proto1.CreateSnapshotRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "CreateSnapshot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) Offload(ctx context.Context, req *proto1.OffloadRequest) (*proto1.OffloadResponse, error) {
	var resp proto1.OffloadResponse
	if err := c.client.Call(ctx, "Firecracker", "Offload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
