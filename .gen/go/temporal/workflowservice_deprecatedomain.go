// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package temporal

import (
	"errors"
	"fmt"
	"go.temporal.io/temporal/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_DeprecateDomain_Args represents the arguments for the WorkflowService.DeprecateDomain function.
//
// The arguments for DeprecateDomain are sent and received over the wire as this struct.
type WorkflowService_DeprecateDomain_Args struct {
	DeprecateRequest *shared.DeprecateDomainRequest `json:"deprecateRequest,omitempty"`
}

// ToWire translates a WorkflowService_DeprecateDomain_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_DeprecateDomain_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.DeprecateRequest != nil {
		w, err = v.DeprecateRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DeprecateDomainRequest_Read(w wire.Value) (*shared.DeprecateDomainRequest, error) {
	var v shared.DeprecateDomainRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_DeprecateDomain_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_DeprecateDomain_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_DeprecateDomain_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_DeprecateDomain_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DeprecateRequest, err = _DeprecateDomainRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_DeprecateDomain_Args
// struct.
func (v *WorkflowService_DeprecateDomain_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.DeprecateRequest != nil {
		fields[i] = fmt.Sprintf("DeprecateRequest: %v", v.DeprecateRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_DeprecateDomain_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_DeprecateDomain_Args match the
// provided WorkflowService_DeprecateDomain_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_DeprecateDomain_Args) Equals(rhs *WorkflowService_DeprecateDomain_Args) bool {
	if !((v.DeprecateRequest == nil && rhs.DeprecateRequest == nil) || (v.DeprecateRequest != nil && rhs.DeprecateRequest != nil && v.DeprecateRequest.Equals(rhs.DeprecateRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "DeprecateDomain" for this struct.
func (v *WorkflowService_DeprecateDomain_Args) MethodName() string {
	return "DeprecateDomain"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_DeprecateDomain_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_DeprecateDomain_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.DeprecateDomain
// function.
var WorkflowService_DeprecateDomain_Helper = struct {
	// Args accepts the parameters of DeprecateDomain in-order and returns
	// the arguments struct for the function.
	Args func(
		deprecateRequest *shared.DeprecateDomainRequest,
	) *WorkflowService_DeprecateDomain_Args

	// IsException returns true if the given error can be thrown
	// by DeprecateDomain.
	//
	// An error can be thrown by DeprecateDomain only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for DeprecateDomain
	// given the error returned by it. The provided error may
	// be nil if DeprecateDomain did not fail.
	//
	// This allows mapping errors returned by DeprecateDomain into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// DeprecateDomain
	//
	//   err := DeprecateDomain(args)
	//   result, err := WorkflowService_DeprecateDomain_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from DeprecateDomain: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*WorkflowService_DeprecateDomain_Result, error)

	// UnwrapResponse takes the result struct for DeprecateDomain
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if DeprecateDomain threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := WorkflowService_DeprecateDomain_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_DeprecateDomain_Result) error
}{}

func init() {
	WorkflowService_DeprecateDomain_Helper.Args = func(
		deprecateRequest *shared.DeprecateDomainRequest,
	) *WorkflowService_DeprecateDomain_Args {
		return &WorkflowService_DeprecateDomain_Args{
			DeprecateRequest: deprecateRequest,
		}
	}

	WorkflowService_DeprecateDomain_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.ServiceBusyError:
			return true
		case *shared.DomainNotActiveError:
			return true
		case *shared.ClientVersionNotSupportedError:
			return true
		default:
			return false
		}
	}

	WorkflowService_DeprecateDomain_Helper.WrapResponse = func(err error) (*WorkflowService_DeprecateDomain_Result, error) {
		if err == nil {
			return &WorkflowService_DeprecateDomain_Result{}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.BadRequestError")
			}
			return &WorkflowService_DeprecateDomain_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.InternalServiceError")
			}
			return &WorkflowService_DeprecateDomain_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.EntityNotExistError")
			}
			return &WorkflowService_DeprecateDomain_Result{EntityNotExistError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.ServiceBusyError")
			}
			return &WorkflowService_DeprecateDomain_Result{ServiceBusyError: e}, nil
		case *shared.DomainNotActiveError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.DomainNotActiveError")
			}
			return &WorkflowService_DeprecateDomain_Result{DomainNotActiveError: e}, nil
		case *shared.ClientVersionNotSupportedError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DeprecateDomain_Result.ClientVersionNotSupportedError")
			}
			return &WorkflowService_DeprecateDomain_Result{ClientVersionNotSupportedError: e}, nil
		}

		return nil, err
	}
	WorkflowService_DeprecateDomain_Helper.UnwrapResponse = func(result *WorkflowService_DeprecateDomain_Result) (err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}
		if result.DomainNotActiveError != nil {
			err = result.DomainNotActiveError
			return
		}
		if result.ClientVersionNotSupportedError != nil {
			err = result.ClientVersionNotSupportedError
			return
		}
		return
	}

}

// WorkflowService_DeprecateDomain_Result represents the result of a WorkflowService.DeprecateDomain function call.
//
// The result of a DeprecateDomain execution is sent and received over the wire as this struct.
type WorkflowService_DeprecateDomain_Result struct {
	BadRequestError                *shared.BadRequestError                `json:"badRequestError,omitempty"`
	InternalServiceError           *shared.InternalServiceError           `json:"internalServiceError,omitempty"`
	EntityNotExistError            *shared.EntityNotExistsError           `json:"entityNotExistError,omitempty"`
	ServiceBusyError               *shared.ServiceBusyError               `json:"serviceBusyError,omitempty"`
	DomainNotActiveError           *shared.DomainNotActiveError           `json:"domainNotActiveError,omitempty"`
	ClientVersionNotSupportedError *shared.ClientVersionNotSupportedError `json:"clientVersionNotSupportedError,omitempty"`
}

// ToWire translates a WorkflowService_DeprecateDomain_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_DeprecateDomain_Result) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.DomainNotActiveError != nil {
		w, err = v.DomainNotActiveError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.ClientVersionNotSupportedError != nil {
		w, err = v.ClientVersionNotSupportedError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_DeprecateDomain_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DomainNotActiveError_Read(w wire.Value) (*shared.DomainNotActiveError, error) {
	var v shared.DomainNotActiveError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_DeprecateDomain_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_DeprecateDomain_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_DeprecateDomain_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_DeprecateDomain_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.DomainNotActiveError, err = _DomainNotActiveError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.ClientVersionNotSupportedError, err = _ClientVersionNotSupportedError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if v.DomainNotActiveError != nil {
		count++
	}
	if v.ClientVersionNotSupportedError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("WorkflowService_DeprecateDomain_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_DeprecateDomain_Result
// struct.
func (v *WorkflowService_DeprecateDomain_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}
	if v.DomainNotActiveError != nil {
		fields[i] = fmt.Sprintf("DomainNotActiveError: %v", v.DomainNotActiveError)
		i++
	}
	if v.ClientVersionNotSupportedError != nil {
		fields[i] = fmt.Sprintf("ClientVersionNotSupportedError: %v", v.ClientVersionNotSupportedError)
		i++
	}

	return fmt.Sprintf("WorkflowService_DeprecateDomain_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_DeprecateDomain_Result match the
// provided WorkflowService_DeprecateDomain_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_DeprecateDomain_Result) Equals(rhs *WorkflowService_DeprecateDomain_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}
	if !((v.DomainNotActiveError == nil && rhs.DomainNotActiveError == nil) || (v.DomainNotActiveError != nil && rhs.DomainNotActiveError != nil && v.DomainNotActiveError.Equals(rhs.DomainNotActiveError))) {
		return false
	}
	if !((v.ClientVersionNotSupportedError == nil && rhs.ClientVersionNotSupportedError == nil) || (v.ClientVersionNotSupportedError != nil && rhs.ClientVersionNotSupportedError != nil && v.ClientVersionNotSupportedError.Equals(rhs.ClientVersionNotSupportedError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "DeprecateDomain" for this struct.
func (v *WorkflowService_DeprecateDomain_Result) MethodName() string {
	return "DeprecateDomain"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_DeprecateDomain_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
