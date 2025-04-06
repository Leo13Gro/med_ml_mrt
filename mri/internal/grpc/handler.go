package grpc

import (
	"mri/internal/generated/grpc/service"
	"mri/internal/grpc/device"
	"mri/internal/grpc/image"
	"mri/internal/grpc/mri"
	"mri/internal/grpc/node"
	"mri/internal/grpc/segment"
)

type Handler struct {
	device.DeviceHandler
	mri.MriHandler
	image.ImageHandler
	node.NodeHandler
	segment.SegmentHandler

	service.UnsafeMriSrvServer
}

func New(
	deviceHandler device.DeviceHandler,
	mriHandler mri.MriHandler,
	imageHandler image.ImageHandler,
	nodeHandler node.NodeHandler,
	segmentHandler segment.SegmentHandler,
) *Handler {
	return &Handler{
		DeviceHandler:  deviceHandler,
		MriHandler:     mriHandler,
		ImageHandler:   imageHandler,
		NodeHandler:    nodeHandler,
		SegmentHandler: segmentHandler,
	}
}
