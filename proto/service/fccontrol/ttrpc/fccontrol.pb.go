// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fccontrol.proto

package fccontrol

import (
	context "context"
	fmt "fmt"
	math "math"

	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto1 "github.com/firecracker-microvm/firecracker-containerd/proto"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4b, 0x4e, 0xce,
	0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xa1,
	0x92, 0x82, 0x69, 0x99, 0x45, 0xa9, 0xc9, 0x45, 0x89, 0xc9, 0xd9, 0xa9, 0x45, 0x10, 0x21, 0xa3,
	0x2b, 0xcc, 0x5c, 0xdc, 0x6e, 0x08, 0x51, 0x21, 0x7d, 0x2e, 0x0e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92,
	0xd4, 0x30, 0x5f, 0x21, 0x01, 0x3d, 0x18, 0x33, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x4a,
	0x10, 0x49, 0xa4, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x88, 0x8b, 0x3d, 0x20, 0xb1, 0xb4,
	0x18, 0xa4, 0x9e, 0x5f, 0x0f, 0xca, 0x82, 0x29, 0x17, 0xd3, 0x83, 0xb8, 0x46, 0x0f, 0xe6, 0x1a,
	0x3d, 0x57, 0x90, 0x6b, 0x84, 0x4c, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x21, 0x96, 0xc0,
	0x98, 0x84, 0x74, 0x19, 0x70, 0xb1, 0x05, 0x97, 0xe4, 0x17, 0x84, 0xf9, 0x0a, 0xf1, 0xe9, 0x41,
	0x18, 0x84, 0x74, 0x18, 0x71, 0x71, 0xba, 0xa7, 0x96, 0x84, 0xf9, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b,
	0x09, 0xea, 0xc1, 0xd9, 0x30, 0x7d, 0x42, 0xc8, 0x42, 0x50, 0xff, 0xd8, 0x71, 0xf1, 0x06, 0x83,
	0x04, 0x7d, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x85, 0x44, 0xf5, 0x50, 0xf8, 0x84, 0xec,
	0x74, 0xe1, 0x12, 0x08, 0x2d, 0x48, 0x01, 0x87, 0x11, 0xdc, 0x08, 0x09, 0x3d, 0x74, 0x21, 0x42,
	0xa6, 0xd8, 0x71, 0xf1, 0xba, 0xa3, 0xb9, 0xc2, 0x1d, 0xbb, 0x2b, 0xd0, 0x84, 0x21, 0xbe, 0x70,
	0x52, 0x3e, 0xf1, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x86, 0x47, 0x72, 0x8c, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x14, 0x27, 0x3c, 0xc5, 0x24,
	0xb1, 0x81, 0x2d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x26, 0xe6, 0x13, 0x45, 0x02,
	0x00, 0x00,
}

type FirecrackerService interface {
	CreateVM(ctx context.Context, req *proto1.CreateVMRequest) (*proto1.CreateVMResponse, error)
	PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error)
	ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error)
	StopVM(ctx context.Context, req *proto1.StopVMRequest) (*empty.Empty, error)
	GetVMInfo(ctx context.Context, req *proto1.GetVMInfoRequest) (*proto1.GetVMInfoResponse, error)
	SetVMMetadata(ctx context.Context, req *proto1.SetVMMetadataRequest) (*empty.Empty, error)
	UpdateVMMetadata(ctx context.Context, req *proto1.UpdateVMMetadataRequest) (*empty.Empty, error)
	GetVMMetadata(ctx context.Context, req *proto1.GetVMMetadataRequest) (*proto1.GetVMMetadataResponse, error)
	PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error)
	ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error)
	LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*proto1.LoadResponse, error)
	CreateSnapshot(ctx context.Context, req *proto1.CreateSnapshotRequest) (*empty.Empty, error)
	Offload(ctx context.Context, req *proto1.OffloadRequest) (*empty.Empty, error)
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

func (c *firecrackerClient) LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*proto1.LoadResponse, error) {
	var resp proto1.LoadResponse
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

func (c *firecrackerClient) Offload(ctx context.Context, req *proto1.OffloadRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "Offload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
