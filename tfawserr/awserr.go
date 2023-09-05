// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfawserr

import (
	"strings"

	smithy "github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/hashicorp/aws-sdk-go-base/v2/internal/errs"
)

// ErrCodeEquals returns true if the error matches all these conditions:
//   - err is of type smithy.APIError
//   - APIError.ErrorCode() equals one of the passed codes
func ErrCodeEquals(err error, codes ...string) bool {
	if apiErr, ok := errs.As[smithy.APIError](err); ok {
		for _, code := range codes {
			if apiErr.ErrorCode() == code {
				return true
			}
		}
	}
	return false
}

// ErrMessageContains returns true if the error matches all these conditions:
//   - err is of type smithy.APIError
//   - APIError.ErrorCode() equals code
//   - APIError.ErrorMessage() contains message
func ErrMessageContains(err error, code string, message string) bool {
	if apiErr, ok := errs.As[smithy.APIError](err); ok {
		return apiErr.ErrorCode() == code && strings.Contains(apiErr.ErrorMessage(), message)
	}
	return false
}

// ErrHTTPStatusCodeEquals returns true if the error matches all these conditions:
//   - err is of type smithyhttp.ResponseError
//   - ResponseError.HTTPStatusCode() equals one of the passed status codes
func ErrHTTPStatusCodeEquals(err error, statusCodes ...int) bool {
	if respErr, ok := errs.As[*smithyhttp.ResponseError](err); ok {
		for _, statusCode := range statusCodes {
			if respErr.HTTPStatusCode() == statusCode {
				return true
			}
		}
	}
	return false
}
