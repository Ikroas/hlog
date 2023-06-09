// Copyright (c) 2012 - Cloud Instruments Co., Ltd.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package hlog

import (
	"errors"
)

var (
	errNodeMustHaveChildren   = errors.New("node must have children")
	errNodeCannotHaveChildren = errors.New("node cannot have children")
)

type unexpectedChildElementError struct {
	baseError
}

func newUnexpectedChildElementError(msg string) *unexpectedChildElementError {
	custmsg := "Unexpected child element: " + msg
	return &unexpectedChildElementError{baseError{message: custmsg}}
}

type missingArgumentError struct {
	baseError
}

func newMissingArgumentError(nodeName, attrName string) *missingArgumentError {
	custmsg := "Output '" + nodeName + "' has no '" + attrName + "' attribute"
	return &missingArgumentError{baseError{message: custmsg}}
}

type unexpectedAttributeError struct {
	baseError
}

func newUnexpectedAttributeError(nodeName, attr string) *unexpectedAttributeError {
	custmsg := nodeName + " has unexpected attribute: " + attr
	return &unexpectedAttributeError{baseError{message: custmsg}}
}
