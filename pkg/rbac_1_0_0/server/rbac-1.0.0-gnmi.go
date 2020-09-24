// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"encoding/json"
	"fmt"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	modelplugin "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
)

// gnmiDeleteRbacV100targetRbac deletes an instance of RbacV100targetRbac.
func (i *ServerImpl) gnmiDeleteRbacV100targetRbac(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacV100targetRbac returns an instance of RbacV100targetRbac.
func (i *ServerImpl) gnmiGetRbacV100targetRbac(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacV100targetRbac, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListRbacV100targetRbac()
}

// gnmiPostRbacV100targetRbac adds an instance of RbacV100targetRbac.
func (i *ServerImpl) gnmiPostRbacV100targetRbac(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacV100targetRbac)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.RbacV100targetRbac %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacV100targetRbac(jsonObj, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacV100targetRbac to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacV100targetRbacGroup deletes an instance of RbacV100targetRbacGroup.
func (i *ServerImpl) gnmiDeleteRbacV100targetRbacGroup(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacV100targetRbacGroup returns an instance of RbacV100targetRbacGroup.
func (i *ServerImpl) gnmiGetRbacV100targetRbacGroup(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacV100targetRbacGroup, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListRbacV100targetRbacGroup()
}

// gnmiPostRbacV100targetRbacGroup adds an instance of RbacV100targetRbacGroup.
func (i *ServerImpl) gnmiPostRbacV100targetRbacGroup(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacV100targetRbacGroup)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.RbacV100targetRbacGroup %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacGroup(jsonObj, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacV100targetRbacGroup to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacV100targetRbacGroupRole deletes an instance of RbacV100targetRbacGroupRole.
func (i *ServerImpl) gnmiDeleteRbacV100targetRbacGroupRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacV100targetRbacGroupRole returns an instance of RbacV100targetRbacGroupRole.
func (i *ServerImpl) gnmiGetRbacV100targetRbacGroupRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacV100targetRbacGroupRole, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListRbacV100targetRbacGroupRole()
}

// gnmiPostRbacV100targetRbacGroupRole adds an instance of RbacV100targetRbacGroupRole.
func (i *ServerImpl) gnmiPostRbacV100targetRbacGroupRole(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacV100targetRbacGroupRole)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.RbacV100targetRbacGroupRole %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacGroupRole(jsonObj, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacV100targetRbacGroupRole to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacV100targetRbacRole deletes an instance of RbacV100targetRbacRole.
func (i *ServerImpl) gnmiDeleteRbacV100targetRbacRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacV100targetRbacRole returns an instance of RbacV100targetRbacRole.
func (i *ServerImpl) gnmiGetRbacV100targetRbacRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacV100targetRbacRole, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListRbacV100targetRbacRole()
}

// gnmiPostRbacV100targetRbacRole adds an instance of RbacV100targetRbacRole.
func (i *ServerImpl) gnmiPostRbacV100targetRbacRole(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacV100targetRbacRole)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.RbacV100targetRbacRole %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacRole(jsonObj, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacV100targetRbacRole to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacV100targetRbacRolePermission deletes an instance of RbacV100targetRbacRolePermission.
func (i *ServerImpl) gnmiDeleteRbacV100targetRbacRolePermission(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacV100targetRbacRolePermission returns an instance of RbacV100targetRbacRolePermission.
func (i *ServerImpl) gnmiGetRbacV100targetRbacRolePermission(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacV100targetRbacRolePermission, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListRbacV100targetRbacRolePermission()
}

// gnmiPostRbacV100targetRbacRolePermission adds an instance of RbacV100targetRbacRolePermission.
func (i *ServerImpl) gnmiPostRbacV100targetRbacRolePermission(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacV100targetRbacRolePermission)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.RbacV100targetRbacRolePermission %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacRolePermission(jsonObj, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacV100targetRbacRolePermission to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteTarget deletes an instance of target.
func (i *ServerImpl) gnmiDeleteTarget(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetTarget returns an instance of target.
func (i *ServerImpl) gnmiGetTarget(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.Target, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.handlePropListTarget()
}

// gnmiPostTarget adds an instance of target.
func (i *ServerImpl) gnmiPostTarget(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	return nil, fmt.Errorf("Not implemented")

}

type Translator interface {
	handlePropListRbacV100targetRbac() (*types.RbacV100targetRbac, error)
	handlePropListRbacV100targetRbacGroup() (*types.RbacV100targetRbacGroup, error)
	handlePropListRbacV100targetRbacGroupRole() (*types.RbacV100targetRbacGroupRole, error)
	handlePropListRbacV100targetRbacRole() (*types.RbacV100targetRbacRole, error)
	handlePropListRbacV100targetRbacRolePermission() (*types.RbacV100targetRbacRolePermission, error)
	handlePropListTarget() (*types.Target, error)
}