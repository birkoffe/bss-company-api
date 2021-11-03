// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/bss_company_api/v1/bss_company_api.proto

package bss_company_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Company) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompanyValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CompanyValidationError is the validation error returned by Company.Validate
// if the designated constraints aren't met.
type CompanyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompanyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompanyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompanyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompanyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompanyValidationError) ErrorName() string { return "CompanyValidationError" }

// Error satisfies the builtin error interface
func (e CompanyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompany.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompanyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompanyValidationError{}

// Validate checks the field values on DescribecompanyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribecompanyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCompanyId() <= 0 {
		return DescribecompanyV1RequestValidationError{
			field:  "CompanyId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribecompanyV1RequestValidationError is the validation error returned by
// DescribecompanyV1Request.Validate if the designated constraints aren't met.
type DescribecompanyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribecompanyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribecompanyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribecompanyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribecompanyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribecompanyV1RequestValidationError) ErrorName() string {
	return "DescribecompanyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribecompanyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribecompanyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribecompanyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribecompanyV1RequestValidationError{}

// Validate checks the field values on DescribecompanyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribecompanyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribecompanyV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribecompanyV1ResponseValidationError is the validation error returned by
// DescribecompanyV1Response.Validate if the designated constraints aren't met.
type DescribecompanyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribecompanyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribecompanyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribecompanyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribecompanyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribecompanyV1ResponseValidationError) ErrorName() string {
	return "DescribecompanyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribecompanyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribecompanyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribecompanyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribecompanyV1ResponseValidationError{}
