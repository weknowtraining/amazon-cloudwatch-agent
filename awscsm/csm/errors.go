// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package csm

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServiceException":  newErrorInternalServiceException,
	"InvalidRequestException":   newErrorInvalidRequestException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}
