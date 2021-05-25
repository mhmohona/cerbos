// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: response/v1/response.proto

package responsev1

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

	sharedv1 "github.com/cerbos/cerbos/internal/genpb/shared/v1"
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

	_ = sharedv1.Effect(0)
)

// Validate checks the field values on CheckResourceSetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceSetResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	for key, val := range m.GetResourceInstances() {
		_ = val

		// no validation rules for ResourceInstances[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponseValidationError{
					field:  fmt.Sprintf("ResourceInstances[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceSetResponseValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CheckResourceSetResponseValidationError is the validation error returned by
// CheckResourceSetResponse.Validate if the designated constraints aren't met.
type CheckResourceSetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponseValidationError) ErrorName() string {
	return "CheckResourceSetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponseValidationError{}

// Validate checks the field values on CheckResourceBatchResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceBatchResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceBatchResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceBatchResponseValidationError is the validation error returned
// by CheckResourceBatchResponse.Validate if the designated constraints aren't met.
type CheckResourceBatchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchResponseValidationError) ErrorName() string {
	return "CheckResourceBatchResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchResponseValidationError{}

// Validate checks the field values on PlaygroundResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlaygroundId

	switch m.Outcome.(type) {

	case *PlaygroundResponse_Failure:

		if v, ok := interface{}(m.GetFailure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundResponseValidationError{
					field:  "Failure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PlaygroundResponse_Success:

		if v, ok := interface{}(m.GetSuccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundResponseValidationError{
					field:  "Success",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundResponseValidationError is the validation error returned by
// PlaygroundResponse.Validate if the designated constraints aren't met.
type PlaygroundResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundResponseValidationError) ErrorName() string {
	return "PlaygroundResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundResponseValidationError{}

// Validate checks the field values on CheckResourceSetResponse_ActionEffectMap
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_ActionEffectMap) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Actions

	return nil
}

// CheckResourceSetResponse_ActionEffectMapValidationError is the validation
// error returned by CheckResourceSetResponse_ActionEffectMap.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_ActionEffectMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) ErrorName() string {
	return "CheckResourceSetResponse_ActionEffectMapValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_ActionEffectMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_ActionEffectMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_ActionEffectMapValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceSetResponse_Meta) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetResourceInstances() {
		_ = val

		// no validation rules for ResourceInstances[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponse_MetaValidationError{
					field:  fmt.Sprintf("ResourceInstances[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceSetResponse_MetaValidationError is the validation error
// returned by CheckResourceSetResponse_Meta.Validate if the designated
// constraints aren't met.
type CheckResourceSetResponse_MetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_MetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_MetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_MetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_MetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_MetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_MetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_MetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_MetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_MetaValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta_EffectMeta
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_Meta_EffectMeta) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MatchedPolicy

	return nil
}

// CheckResourceSetResponse_Meta_EffectMetaValidationError is the validation
// error returned by CheckResourceSetResponse_Meta_EffectMeta.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_Meta_EffectMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_Meta_EffectMetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta_EffectMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_Meta_EffectMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_Meta_EffectMetaValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta_ActionMeta
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_Meta_ActionMeta) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetActions() {
		_ = val

		// no validation rules for Actions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponse_Meta_ActionMetaValidationError{
					field:  fmt.Sprintf("Actions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceSetResponse_Meta_ActionMetaValidationError is the validation
// error returned by CheckResourceSetResponse_Meta_ActionMeta.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_Meta_ActionMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_Meta_ActionMetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta_ActionMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_Meta_ActionMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_Meta_ActionMetaValidationError{}

// Validate checks the field values on
// CheckResourceBatchResponse_ActionEffectMap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CheckResourceBatchResponse_ActionEffectMap) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResourceId

	// no validation rules for Actions

	return nil
}

// CheckResourceBatchResponse_ActionEffectMapValidationError is the validation
// error returned by CheckResourceBatchResponse_ActionEffectMap.Validate if
// the designated constraints aren't met.
type CheckResourceBatchResponse_ActionEffectMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) ErrorName() string {
	return "CheckResourceBatchResponse_ActionEffectMapValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchResponse_ActionEffectMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchResponse_ActionEffectMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchResponse_ActionEffectMapValidationError{}

// Validate checks the field values on PlaygroundResponse_Error with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundResponse_Error) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for File

	// no validation rules for Error

	return nil
}

// PlaygroundResponse_ErrorValidationError is the validation error returned by
// PlaygroundResponse_Error.Validate if the designated constraints aren't met.
type PlaygroundResponse_ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundResponse_ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundResponse_ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundResponse_ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundResponse_ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundResponse_ErrorValidationError) ErrorName() string {
	return "PlaygroundResponse_ErrorValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundResponse_ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundResponse_Error.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundResponse_ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundResponse_ErrorValidationError{}

// Validate checks the field values on PlaygroundResponse_ErrorList with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundResponse_ErrorList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundResponse_ErrorListValidationError{
					field:  fmt.Sprintf("Errors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundResponse_ErrorListValidationError is the validation error returned
// by PlaygroundResponse_ErrorList.Validate if the designated constraints
// aren't met.
type PlaygroundResponse_ErrorListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundResponse_ErrorListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundResponse_ErrorListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundResponse_ErrorListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundResponse_ErrorListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundResponse_ErrorListValidationError) ErrorName() string {
	return "PlaygroundResponse_ErrorListValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundResponse_ErrorListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundResponse_ErrorList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundResponse_ErrorListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundResponse_ErrorListValidationError{}

// Validate checks the field values on PlaygroundResponse_EvalResult with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundResponse_EvalResult) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Action

	// no validation rules for Effect

	// no validation rules for Policy

	return nil
}

// PlaygroundResponse_EvalResultValidationError is the validation error
// returned by PlaygroundResponse_EvalResult.Validate if the designated
// constraints aren't met.
type PlaygroundResponse_EvalResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundResponse_EvalResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundResponse_EvalResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundResponse_EvalResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundResponse_EvalResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundResponse_EvalResultValidationError) ErrorName() string {
	return "PlaygroundResponse_EvalResultValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundResponse_EvalResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundResponse_EvalResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundResponse_EvalResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundResponse_EvalResultValidationError{}

// Validate checks the field values on PlaygroundResponse_EvalResultList with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *PlaygroundResponse_EvalResultList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundResponse_EvalResultListValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundResponse_EvalResultListValidationError is the validation error
// returned by PlaygroundResponse_EvalResultList.Validate if the designated
// constraints aren't met.
type PlaygroundResponse_EvalResultListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundResponse_EvalResultListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundResponse_EvalResultListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundResponse_EvalResultListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundResponse_EvalResultListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundResponse_EvalResultListValidationError) ErrorName() string {
	return "PlaygroundResponse_EvalResultListValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundResponse_EvalResultListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundResponse_EvalResultList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundResponse_EvalResultListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundResponse_EvalResultListValidationError{}
