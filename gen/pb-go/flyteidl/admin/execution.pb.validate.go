// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/execution.proto

package admin

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ExecutionCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionCreateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionCreateRequestValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionCreateRequestValidationError is the validation error returned by
// ExecutionCreateRequest.Validate if the designated constraints aren't met.
type ExecutionCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionCreateRequestValidationError) ErrorName() string {
	return "ExecutionCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionCreateRequestValidationError{}

// Validate checks the field values on ExecutionCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Urn

	return nil
}

// ExecutionCreateResponseValidationError is the validation error returned by
// ExecutionCreateResponse.Validate if the designated constraints aren't met.
type ExecutionCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionCreateResponseValidationError) ErrorName() string {
	return "ExecutionCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionCreateResponseValidationError{}

// Validate checks the field values on Execution with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Execution) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ExecutionId

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionValidationError is the validation error returned by
// Execution.Validate if the designated constraints aren't met.
type ExecutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionValidationError) ErrorName() string { return "ExecutionValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionValidationError{}

// Validate checks the field values on ExecutionList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExecutionList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetExecutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionListValidationError{
					field:  fmt.Sprintf("Executions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExecutionListValidationError is the validation error returned by
// ExecutionList.Validate if the designated constraints aren't met.
type ExecutionListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionListValidationError) ErrorName() string { return "ExecutionListValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionListValidationError{}

// Validate checks the field values on LiteralMapBlob with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LiteralMapBlob) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Data.(type) {

	case *LiteralMapBlob_Values:

		if v, ok := interface{}(m.GetValues()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralMapBlobValidationError{
					field:  "Values",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralMapBlob_Uri:
		// no validation rules for Uri

	}

	return nil
}

// LiteralMapBlobValidationError is the validation error returned by
// LiteralMapBlob.Validate if the designated constraints aren't met.
type LiteralMapBlobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LiteralMapBlobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LiteralMapBlobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LiteralMapBlobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LiteralMapBlobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LiteralMapBlobValidationError) ErrorName() string { return "LiteralMapBlobValidationError" }

// Error satisfies the builtin error interface
func (e LiteralMapBlobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLiteralMapBlob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LiteralMapBlobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LiteralMapBlobValidationError{}

// Validate checks the field values on ExecutionClosure with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExecutionClosure) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetComputedInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "ComputedInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Phase

	if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.OutputResult.(type) {

	case *ExecutionClosure_Outputs:

		if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionClosureValidationError{
					field:  "Outputs",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExecutionClosure_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionClosureValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExecutionClosureValidationError is the validation error returned by
// ExecutionClosure.Validate if the designated constraints aren't met.
type ExecutionClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionClosureValidationError) ErrorName() string { return "ExecutionClosureValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionClosureValidationError{}

// Validate checks the field values on ExecutionMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mode

	// no validation rules for Principal

	// no validation rules for Nesting

	return nil
}

// ExecutionMetadataValidationError is the validation error returned by
// ExecutionMetadata.Validate if the designated constraints aren't met.
type ExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionMetadataValidationError) ErrorName() string {
	return "ExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionMetadataValidationError{}

// Validate checks the field values on ExecutionSpec with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExecutionSpec) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for LaunchPlanUrn

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionSpecValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExecutionSpecValidationError is the validation error returned by
// ExecutionSpec.Validate if the designated constraints aren't met.
type ExecutionSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionSpecValidationError) ErrorName() string { return "ExecutionSpecValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionSpecValidationError{}
