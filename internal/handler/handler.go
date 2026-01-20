package handler

import (
	pb "github.com/jne100/golang-service-layout/api"
	"github.com/jne100/golang-service-layout/internal/controller"
	"github.com/jne100/golang-service-layout/internal/handler/argsvalidator"
)

type handler struct {
	pb.UnimplementedInventoryServiceServer
	argsValidator argsvalidator.ArgsValidator
	ctrl          controller.Controller
}
