package labservicesapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/labservices/mgmt/2018-10-15/labservices"
	"github.com/Azure/go-autorest/autorest"
)

// ProviderOperationsClientAPI contains the set of methods on the ProviderOperationsClient type.
type ProviderOperationsClientAPI interface {
	List(ctx context.Context) (result labservices.ProviderOperationResultPage, err error)
}

var _ ProviderOperationsClientAPI = (*labservices.ProviderOperationsClient)(nil)

// GlobalUsersClientAPI contains the set of methods on the GlobalUsersClient type.
type GlobalUsersClientAPI interface {
	GetEnvironment(ctx context.Context, userName string, environmentOperationsPayload labservices.EnvironmentOperationsPayload, expand string) (result labservices.GetEnvironmentResponse, err error)
	GetOperationBatchStatus(ctx context.Context, userName string, operationBatchStatusPayload labservices.OperationBatchStatusPayload) (result labservices.OperationBatchStatusResponse, err error)
	GetOperationStatus(ctx context.Context, userName string, operationStatusPayload labservices.OperationStatusPayload) (result labservices.OperationStatusResponse, err error)
	GetPersonalPreferences(ctx context.Context, userName string, personalPreferencesOperationsPayload labservices.PersonalPreferencesOperationsPayload) (result labservices.GetPersonalPreferencesResponse, err error)
	ListEnvironments(ctx context.Context, userName string, listEnvironmentsPayload labservices.ListEnvironmentsPayload) (result labservices.ListEnvironmentsResponse, err error)
	ListLabs(ctx context.Context, userName string) (result labservices.ListLabsResponse, err error)
	Register(ctx context.Context, userName string, registerPayload labservices.RegisterPayload) (result autorest.Response, err error)
	ResetPassword(ctx context.Context, userName string, resetPasswordPayload labservices.ResetPasswordPayload) (result labservices.GlobalUsersResetPasswordFuture, err error)
	StartEnvironment(ctx context.Context, userName string, environmentOperationsPayload labservices.EnvironmentOperationsPayload) (result labservices.GlobalUsersStartEnvironmentFuture, err error)
	StopEnvironment(ctx context.Context, userName string, environmentOperationsPayload labservices.EnvironmentOperationsPayload) (result labservices.GlobalUsersStopEnvironmentFuture, err error)
}

var _ GlobalUsersClientAPI = (*labservices.GlobalUsersClient)(nil)

// LabAccountsClientAPI contains the set of methods on the LabAccountsClient type.
type LabAccountsClientAPI interface {
	CreateLab(ctx context.Context, resourceGroupName string, labAccountName string, createLabProperties labservices.CreateLabProperties) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labAccount labservices.LabAccount) (result labservices.LabAccount, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string) (result labservices.LabAccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, expand string) (result labservices.LabAccount, err error)
	GetRegionalAvailability(ctx context.Context, resourceGroupName string, labAccountName string) (result labservices.GetRegionalAvailabilityResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationLabAccountPage, err error)
	ListBySubscription(ctx context.Context, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationLabAccountPage, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, labAccount labservices.LabAccountFragment) (result labservices.LabAccount, err error)
}

var _ LabAccountsClientAPI = (*labservices.LabAccountsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	Get(ctx context.Context, locationName string, operationName string) (result labservices.OperationResult, err error)
}

var _ OperationsClientAPI = (*labservices.OperationsClient)(nil)

// GalleryImagesClientAPI contains the set of methods on the GalleryImagesClient type.
type GalleryImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage labservices.GalleryImage) (result labservices.GalleryImage, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, expand string) (result labservices.GalleryImage, err error)
	List(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationGalleryImagePage, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage labservices.GalleryImageFragment) (result labservices.GalleryImage, err error)
}

var _ GalleryImagesClientAPI = (*labservices.GalleryImagesClient)(nil)

// LabsClientAPI contains the set of methods on the LabsClient type.
type LabsClientAPI interface {
	AddUsers(ctx context.Context, resourceGroupName string, labAccountName string, labName string, addUsersPayload labservices.AddUsersPayload) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab labservices.Lab) (result labservices.Lab, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (result labservices.LabsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string) (result labservices.Lab, err error)
	List(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationLabPage, err error)
	Register(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab labservices.LabFragment) (result labservices.Lab, err error)
}

var _ LabsClientAPI = (*labservices.LabsClient)(nil)

// EnvironmentSettingsClientAPI contains the set of methods on the EnvironmentSettingsClient type.
type EnvironmentSettingsClientAPI interface {
	ClaimAny(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting labservices.EnvironmentSetting) (result labservices.EnvironmentSettingsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result labservices.EnvironmentSettingsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string) (result labservices.EnvironmentSetting, err error)
	List(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationEnvironmentSettingPage, err error)
	Publish(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, publishPayload labservices.PublishPayload) (result autorest.Response, err error)
	Start(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result labservices.EnvironmentSettingsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string) (result labservices.EnvironmentSettingsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentSetting labservices.EnvironmentSettingFragment) (result labservices.EnvironmentSetting, err error)
}

var _ EnvironmentSettingsClientAPI = (*labservices.EnvironmentSettingsClient)(nil)

// EnvironmentsClientAPI contains the set of methods on the EnvironmentsClient type.
type EnvironmentsClientAPI interface {
	Claim(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment labservices.Environment) (result labservices.Environment, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result labservices.EnvironmentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, expand string) (result labservices.Environment, err error)
	List(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationEnvironmentPage, err error)
	ResetPassword(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, resetPasswordPayload labservices.ResetPasswordPayload) (result labservices.EnvironmentsResetPasswordFuture, err error)
	Start(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result labservices.EnvironmentsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result labservices.EnvironmentsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment labservices.EnvironmentFragment) (result labservices.Environment, err error)
}

var _ EnvironmentsClientAPI = (*labservices.EnvironmentsClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, userName string, userParameter labservices.User) (result labservices.User, err error)
	Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, userName string) (result labservices.UsersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, userName string, expand string) (result labservices.User, err error)
	List(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string, filter string, top *int32, orderby string) (result labservices.ResponseWithContinuationUserPage, err error)
	Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, userName string, userParameter labservices.UserFragment) (result labservices.User, err error)
}

var _ UsersClientAPI = (*labservices.UsersClient)(nil)