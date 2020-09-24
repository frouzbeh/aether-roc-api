// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/openconfig/gnmi/proto/gnmi"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_gnmiGetRbacV100targetRbac(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rbacFromGnmi, err := ioutil.ReadFile("../testdata/RbacFromGnmi.json")
	assert.NilError(t, err, "error loadign testdata file")
	mockClient := southbound.NewMockGnmiClient(ctrl)
	mockClient.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
			gr := gnmi.GetResponse{
				Notification: []*gnmi.Notification{
					{
						Update: []*gnmi.Update{
							{
								Val: &gnmi.TypedValue{
									Value: &gnmi.TypedValue_JsonVal{JsonVal: rbacFromGnmi},
								},
							},
						},
					},
				},
			}
			return &gr, nil
		},
	)

	serverImpl := ServerImpl{GnmiClient: mockClient}
	rbacResource, err := serverImpl.gnmiGetRbacV100targetRbac(
		context.Background(), "/rbac/v1.0.0/internal/rbac", "internal")

	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, rbacResource != nil)
	assert.Equal(t, 2, len(*rbacResource.ListRbacV100targetRbacRole), "expected 2 roles")
	assert.Equal(t, 1, len(*rbacResource.ListRbacV100targetRbacGroup), "expected 1 group")
	group1 := (*rbacResource.ListRbacV100targetRbacGroup)[0]
	assert.Equal(t, "menlo-admins", *group1.Groupid)
	assert.Equal(t, 1, len(*group1.ListRbacV100targetRbacGroupRole), "This is wrong - should be 2")
}