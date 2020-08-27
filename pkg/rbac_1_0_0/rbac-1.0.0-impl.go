// Code generated by oapi-codegen. DO NOT EDIT.
// Package rbac_1_0_0 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package rbac_1_0_0

import (
	"github.com/labstack/echo/v4"
)

// server-interface template override

import "github.com/onosproject/onos-lib-go/pkg/logging"

// Implement the Server Interface for access to gNMI
var log = logging.GetLogger("rbac")

// ServerImpl -
type ServerImpl struct {
}

// DeleteRbac impl of gNMI access.
func (w *ServerImpl) DeleteRbac(ctx echo.Context) error {

	log.Infof("DeleteRbac")
	return ctx.JSON(200, nil)
}

// GetRbac impl of gNMI access.
func (w *ServerImpl) GetRbac(ctx echo.Context) error {

	log.Infof("GetRbac")
	return ctx.JSON(200, nil)
}

// PostRbac impl of gNMI access.
func (w *ServerImpl) PostRbac(ctx echo.Context) error {

	log.Infof("PostRbac")
	return ctx.JSON(200, nil)
}

// DeleteRbacGroupRole impl of gNMI access.
func (w *ServerImpl) DeleteRbacGroupRole(ctx echo.Context, roleid string) error {

	log.Infof("DeleteRbacGroupRole")
	return ctx.JSON(200, nil)
}

// GetRbacGroupRole impl of gNMI access.
func (w *ServerImpl) GetRbacGroupRole(ctx echo.Context, roleid string) error {

	log.Infof("GetRbacGroupRole")
	return ctx.JSON(200, nil)
}

// PostRbacGroupRole impl of gNMI access.
func (w *ServerImpl) PostRbacGroupRole(ctx echo.Context, roleid string) error {

	log.Infof("PostRbacGroupRole")
	return ctx.JSON(200, nil)
}

// DeleteRbacGroup impl of gNMI access.
func (w *ServerImpl) DeleteRbacGroup(ctx echo.Context, groupid string) error {

	log.Infof("DeleteRbacGroup")
	return ctx.JSON(200, nil)
}

// GetRbacGroup impl of gNMI access.
func (w *ServerImpl) GetRbacGroup(ctx echo.Context, groupid string) error {

	log.Infof("GetRbacGroup")
	return ctx.JSON(200, nil)
}

// PostRbacGroup impl of gNMI access.
func (w *ServerImpl) PostRbacGroup(ctx echo.Context, groupid string) error {

	log.Infof("PostRbacGroup")
	return ctx.JSON(200, nil)
}

// DeleteRbacRolePermission impl of gNMI access.
func (w *ServerImpl) DeleteRbacRolePermission(ctx echo.Context) error {

	log.Infof("DeleteRbacRolePermission")
	return ctx.JSON(200, nil)
}

// GetRbacRolePermission impl of gNMI access.
func (w *ServerImpl) GetRbacRolePermission(ctx echo.Context) error {

	log.Infof("GetRbacRolePermission")
	return ctx.JSON(200, nil)
}

// PostRbacRolePermission impl of gNMI access.
func (w *ServerImpl) PostRbacRolePermission(ctx echo.Context) error {

	log.Infof("PostRbacRolePermission")
	return ctx.JSON(200, nil)
}

// DeleteRbacRole impl of gNMI access.
func (w *ServerImpl) DeleteRbacRole(ctx echo.Context, roleid string) error {

	log.Infof("DeleteRbacRole")
	return ctx.JSON(200, nil)
}

// GetRbacRole impl of gNMI access.
func (w *ServerImpl) GetRbacRole(ctx echo.Context, roleid string) error {

	log.Infof("GetRbacRole")
	return ctx.JSON(200, nil)
}

// PostRbacRole impl of gNMI access.
func (w *ServerImpl) PostRbacRole(ctx echo.Context, roleid string) error {

	log.Infof("PostRbacRole")
	return ctx.JSON(200, nil)
}

// register template override