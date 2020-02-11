// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/matchable_resource.proto

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

// define the regex for a UUID once up-front
var _matchable_resource_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TaskResourceSpec with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TaskResourceSpec) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Cpu

	// no validation rules for Gpu

	// no validation rules for Memory

	// no validation rules for Storage

	return nil
}

// TaskResourceSpecValidationError is the validation error returned by
// TaskResourceSpec.Validate if the designated constraints aren't met.
type TaskResourceSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskResourceSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskResourceSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskResourceSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskResourceSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskResourceSpecValidationError) ErrorName() string { return "TaskResourceSpecValidationError" }

// Error satisfies the builtin error interface
func (e TaskResourceSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskResourceSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskResourceSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskResourceSpecValidationError{}

// Validate checks the field values on TaskResourceAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskResourceAttributes) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDefaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskResourceAttributesValidationError{
				field:  "Defaults",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLimits()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskResourceAttributesValidationError{
				field:  "Limits",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskResourceAttributesValidationError is the validation error returned by
// TaskResourceAttributes.Validate if the designated constraints aren't met.
type TaskResourceAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskResourceAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskResourceAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskResourceAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskResourceAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskResourceAttributesValidationError) ErrorName() string {
	return "TaskResourceAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e TaskResourceAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskResourceAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskResourceAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskResourceAttributesValidationError{}

// Validate checks the field values on ClusterResourceAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterResourceAttributes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Attributes

	return nil
}

// ClusterResourceAttributesValidationError is the validation error returned by
// ClusterResourceAttributes.Validate if the designated constraints aren't met.
type ClusterResourceAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterResourceAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterResourceAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterResourceAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterResourceAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterResourceAttributesValidationError) ErrorName() string {
	return "ClusterResourceAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterResourceAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterResourceAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterResourceAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterResourceAttributesValidationError{}

// Validate checks the field values on ExecutionQueueAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionQueueAttributes) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ExecutionQueueAttributesValidationError is the validation error returned by
// ExecutionQueueAttributes.Validate if the designated constraints aren't met.
type ExecutionQueueAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionQueueAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionQueueAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionQueueAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionQueueAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionQueueAttributesValidationError) ErrorName() string {
	return "ExecutionQueueAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionQueueAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionQueueAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionQueueAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionQueueAttributesValidationError{}

// Validate checks the field values on ExecutionClusterLabel with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionClusterLabel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// ExecutionClusterLabelValidationError is the validation error returned by
// ExecutionClusterLabel.Validate if the designated constraints aren't met.
type ExecutionClusterLabelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionClusterLabelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionClusterLabelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionClusterLabelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionClusterLabelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionClusterLabelValidationError) ErrorName() string {
	return "ExecutionClusterLabelValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionClusterLabelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionClusterLabel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionClusterLabelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionClusterLabelValidationError{}

// Validate checks the field values on MatchingAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchingAttributes) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Target.(type) {

	case *MatchingAttributes_TaskResourceAttributes:

		if v, ok := interface{}(m.GetTaskResourceAttributes()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchingAttributesValidationError{
					field:  "TaskResourceAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchingAttributes_ClusterResourceAttributes:

		if v, ok := interface{}(m.GetClusterResourceAttributes()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchingAttributesValidationError{
					field:  "ClusterResourceAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchingAttributes_ExecutionQueueAttributes:

		if v, ok := interface{}(m.GetExecutionQueueAttributes()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchingAttributesValidationError{
					field:  "ExecutionQueueAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchingAttributes_ExecutionClusterLabel:

		if v, ok := interface{}(m.GetExecutionClusterLabel()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchingAttributesValidationError{
					field:  "ExecutionClusterLabel",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MatchingAttributesValidationError is the validation error returned by
// MatchingAttributes.Validate if the designated constraints aren't met.
type MatchingAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchingAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchingAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchingAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchingAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchingAttributesValidationError) ErrorName() string {
	return "MatchingAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e MatchingAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchingAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchingAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchingAttributesValidationError{}

// Validate checks the field values on MatchableResourceConfiguration with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchableResourceConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MatchableResourceConfigurationValidationError{
				field:  "Attributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Domain

	// no validation rules for Project

	// no validation rules for Workflow

	return nil
}

// MatchableResourceConfigurationValidationError is the validation error
// returned by MatchableResourceConfiguration.Validate if the designated
// constraints aren't met.
type MatchableResourceConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchableResourceConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchableResourceConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchableResourceConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchableResourceConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchableResourceConfigurationValidationError) ErrorName() string {
	return "MatchableResourceConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e MatchableResourceConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchableResourceConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchableResourceConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchableResourceConfigurationValidationError{}

// Validate checks the field values on ListMatchableResourcesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMatchableResourcesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResourceType

	return nil
}

// ListMatchableResourcesRequestValidationError is the validation error
// returned by ListMatchableResourcesRequest.Validate if the designated
// constraints aren't met.
type ListMatchableResourcesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMatchableResourcesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMatchableResourcesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMatchableResourcesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMatchableResourcesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMatchableResourcesRequestValidationError) ErrorName() string {
	return "ListMatchableResourcesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMatchableResourcesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMatchableResourcesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMatchableResourcesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMatchableResourcesRequestValidationError{}

// Validate checks the field values on ListMatchableResourcesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMatchableResourcesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConfigurations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMatchableResourcesResponseValidationError{
					field:  fmt.Sprintf("Configurations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListMatchableResourcesResponseValidationError is the validation error
// returned by ListMatchableResourcesResponse.Validate if the designated
// constraints aren't met.
type ListMatchableResourcesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMatchableResourcesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMatchableResourcesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMatchableResourcesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMatchableResourcesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMatchableResourcesResponseValidationError) ErrorName() string {
	return "ListMatchableResourcesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMatchableResourcesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMatchableResourcesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMatchableResourcesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMatchableResourcesResponseValidationError{}
