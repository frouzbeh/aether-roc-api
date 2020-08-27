// Code generated by oapi-codegen. DO NOT EDIT.
// Package aether_1_0_0 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package aether_1_0_0

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

// DeleteAccessProfile impl of gNMI access.
func (w *ServerImpl) DeleteAccessProfile(ctx echo.Context) error {

	log.Infof("DeleteAccessProfile")
	return ctx.JSON(200, nil)
}

// GetAccessProfile impl of gNMI access.
func (w *ServerImpl) GetAccessProfile(ctx echo.Context) error {

	log.Infof("GetAccessProfile")
	return ctx.JSON(200, nil)
}

// PostAccessProfile impl of gNMI access.
func (w *ServerImpl) PostAccessProfile(ctx echo.Context) error {

	log.Infof("PostAccessProfile")
	return ctx.JSON(200, nil)
}

// DeleteAccessProfileAccessProfile impl of gNMI access.
func (w *ServerImpl) DeleteAccessProfileAccessProfile(ctx echo.Context, id string) error {

	log.Infof("DeleteAccessProfileAccessProfile")
	return ctx.JSON(200, nil)
}

// GetAccessProfileAccessProfile impl of gNMI access.
func (w *ServerImpl) GetAccessProfileAccessProfile(ctx echo.Context, id string) error {

	log.Infof("GetAccessProfileAccessProfile")
	return ctx.JSON(200, nil)
}

// PostAccessProfileAccessProfile impl of gNMI access.
func (w *ServerImpl) PostAccessProfileAccessProfile(ctx echo.Context, id string) error {

	log.Infof("PostAccessProfileAccessProfile")
	return ctx.JSON(200, nil)
}

// DeleteApnProfile impl of gNMI access.
func (w *ServerImpl) DeleteApnProfile(ctx echo.Context) error {

	log.Infof("DeleteApnProfile")
	return ctx.JSON(200, nil)
}

// GetApnProfile impl of gNMI access.
func (w *ServerImpl) GetApnProfile(ctx echo.Context) error {

	log.Infof("GetApnProfile")
	return ctx.JSON(200, nil)
}

// PostApnProfile impl of gNMI access.
func (w *ServerImpl) PostApnProfile(ctx echo.Context) error {

	log.Infof("PostApnProfile")
	return ctx.JSON(200, nil)
}

// DeleteApnProfileApnProfile impl of gNMI access.
func (w *ServerImpl) DeleteApnProfileApnProfile(ctx echo.Context, id string) error {

	log.Infof("DeleteApnProfileApnProfile")
	return ctx.JSON(200, nil)
}

// GetApnProfileApnProfile impl of gNMI access.
func (w *ServerImpl) GetApnProfileApnProfile(ctx echo.Context, id string) error {

	log.Infof("GetApnProfileApnProfile")
	return ctx.JSON(200, nil)
}

// PostApnProfileApnProfile impl of gNMI access.
func (w *ServerImpl) PostApnProfileApnProfile(ctx echo.Context, id string) error {

	log.Infof("PostApnProfileApnProfile")
	return ctx.JSON(200, nil)
}

// DeleteQosProfile impl of gNMI access.
func (w *ServerImpl) DeleteQosProfile(ctx echo.Context) error {

	log.Infof("DeleteQosProfile")
	return ctx.JSON(200, nil)
}

// GetQosProfile impl of gNMI access.
func (w *ServerImpl) GetQosProfile(ctx echo.Context) error {

	log.Infof("GetQosProfile")
	return ctx.JSON(200, nil)
}

// PostQosProfile impl of gNMI access.
func (w *ServerImpl) PostQosProfile(ctx echo.Context) error {

	log.Infof("PostQosProfile")
	return ctx.JSON(200, nil)
}

// DeleteQosProfileQosProfileApnAmbr impl of gNMI access.
func (w *ServerImpl) DeleteQosProfileQosProfileApnAmbr(ctx echo.Context) error {

	log.Infof("DeleteQosProfileQosProfileApnAmbr")
	return ctx.JSON(200, nil)
}

// GetQosProfileQosProfileApnAmbr impl of gNMI access.
func (w *ServerImpl) GetQosProfileQosProfileApnAmbr(ctx echo.Context) error {

	log.Infof("GetQosProfileQosProfileApnAmbr")
	return ctx.JSON(200, nil)
}

// PostQosProfileQosProfileApnAmbr impl of gNMI access.
func (w *ServerImpl) PostQosProfileQosProfileApnAmbr(ctx echo.Context) error {

	log.Infof("PostQosProfileQosProfileApnAmbr")
	return ctx.JSON(200, nil)
}

// DeleteQosProfileQosProfile impl of gNMI access.
func (w *ServerImpl) DeleteQosProfileQosProfile(ctx echo.Context, id string) error {

	log.Infof("DeleteQosProfileQosProfile")
	return ctx.JSON(200, nil)
}

// GetQosProfileQosProfile impl of gNMI access.
func (w *ServerImpl) GetQosProfileQosProfile(ctx echo.Context, id string) error {

	log.Infof("GetQosProfileQosProfile")
	return ctx.JSON(200, nil)
}

// PostQosProfileQosProfile impl of gNMI access.
func (w *ServerImpl) PostQosProfileQosProfile(ctx echo.Context, id string) error {

	log.Infof("PostQosProfileQosProfile")
	return ctx.JSON(200, nil)
}

// DeleteSubscriber impl of gNMI access.
func (w *ServerImpl) DeleteSubscriber(ctx echo.Context) error {

	log.Infof("DeleteSubscriber")
	return ctx.JSON(200, nil)
}

// GetSubscriber impl of gNMI access.
func (w *ServerImpl) GetSubscriber(ctx echo.Context) error {

	log.Infof("GetSubscriber")
	return ctx.JSON(200, nil)
}

// PostSubscriber impl of gNMI access.
func (w *ServerImpl) PostSubscriber(ctx echo.Context) error {

	log.Infof("PostSubscriber")
	return ctx.JSON(200, nil)
}

// DeleteSubscriberUeProfiles impl of gNMI access.
func (w *ServerImpl) DeleteSubscriberUeProfiles(ctx echo.Context) error {

	log.Infof("DeleteSubscriberUeProfiles")
	return ctx.JSON(200, nil)
}

// GetSubscriberUeProfiles impl of gNMI access.
func (w *ServerImpl) GetSubscriberUeProfiles(ctx echo.Context) error {

	log.Infof("GetSubscriberUeProfiles")
	return ctx.JSON(200, nil)
}

// PostSubscriberUeProfiles impl of gNMI access.
func (w *ServerImpl) PostSubscriberUeProfiles(ctx echo.Context) error {

	log.Infof("PostSubscriberUeProfiles")
	return ctx.JSON(200, nil)
}

// DeleteSubscriberUeProfilesAccessProfile impl of gNMI access.
func (w *ServerImpl) DeleteSubscriberUeProfilesAccessProfile(ctx echo.Context, accessProfile string) error {

	log.Infof("DeleteSubscriberUeProfilesAccessProfile")
	return ctx.JSON(200, nil)
}

// GetSubscriberUeProfilesAccessProfile impl of gNMI access.
func (w *ServerImpl) GetSubscriberUeProfilesAccessProfile(ctx echo.Context, accessProfile string) error {

	log.Infof("GetSubscriberUeProfilesAccessProfile")
	return ctx.JSON(200, nil)
}

// PostSubscriberUeProfilesAccessProfile impl of gNMI access.
func (w *ServerImpl) PostSubscriberUeProfilesAccessProfile(ctx echo.Context, accessProfile string) error {

	log.Infof("PostSubscriberUeProfilesAccessProfile")
	return ctx.JSON(200, nil)
}

// DeleteSubscriberUeServingPlmn impl of gNMI access.
func (w *ServerImpl) DeleteSubscriberUeServingPlmn(ctx echo.Context) error {

	log.Infof("DeleteSubscriberUeServingPlmn")
	return ctx.JSON(200, nil)
}

// GetSubscriberUeServingPlmn impl of gNMI access.
func (w *ServerImpl) GetSubscriberUeServingPlmn(ctx echo.Context) error {

	log.Infof("GetSubscriberUeServingPlmn")
	return ctx.JSON(200, nil)
}

// PostSubscriberUeServingPlmn impl of gNMI access.
func (w *ServerImpl) PostSubscriberUeServingPlmn(ctx echo.Context) error {

	log.Infof("PostSubscriberUeServingPlmn")
	return ctx.JSON(200, nil)
}

// DeleteSubscriberUe impl of gNMI access.
func (w *ServerImpl) DeleteSubscriberUe(ctx echo.Context, ueid string) error {

	log.Infof("DeleteSubscriberUe")
	return ctx.JSON(200, nil)
}

// GetSubscriberUe impl of gNMI access.
func (w *ServerImpl) GetSubscriberUe(ctx echo.Context, ueid string) error {

	log.Infof("GetSubscriberUe")
	return ctx.JSON(200, nil)
}

// PostSubscriberUe impl of gNMI access.
func (w *ServerImpl) PostSubscriberUe(ctx echo.Context, ueid string) error {

	log.Infof("PostSubscriberUe")
	return ctx.JSON(200, nil)
}

// DeleteUpProfile impl of gNMI access.
func (w *ServerImpl) DeleteUpProfile(ctx echo.Context) error {

	log.Infof("DeleteUpProfile")
	return ctx.JSON(200, nil)
}

// GetUpProfile impl of gNMI access.
func (w *ServerImpl) GetUpProfile(ctx echo.Context) error {

	log.Infof("GetUpProfile")
	return ctx.JSON(200, nil)
}

// PostUpProfile impl of gNMI access.
func (w *ServerImpl) PostUpProfile(ctx echo.Context) error {

	log.Infof("PostUpProfile")
	return ctx.JSON(200, nil)
}

// DeleteUpProfileUpProfile impl of gNMI access.
func (w *ServerImpl) DeleteUpProfileUpProfile(ctx echo.Context, id string) error {

	log.Infof("DeleteUpProfileUpProfile")
	return ctx.JSON(200, nil)
}

// GetUpProfileUpProfile impl of gNMI access.
func (w *ServerImpl) GetUpProfileUpProfile(ctx echo.Context, id string) error {

	log.Infof("GetUpProfileUpProfile")
	return ctx.JSON(200, nil)
}

// PostUpProfileUpProfile impl of gNMI access.
func (w *ServerImpl) PostUpProfileUpProfile(ctx echo.Context, id string) error {

	log.Infof("PostUpProfileUpProfile")
	return ctx.JSON(200, nil)
}

// register template override
