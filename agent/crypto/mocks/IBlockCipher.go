// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import log "github.com/aws/amazon-ssm-agent/agent/log"
import mock "github.com/stretchr/testify/mock"

// IBlockCipher is an autogenerated mock type for the IBlockCipher type
type IBlockCipher struct {
	mock.Mock
}

// DecryptWithAESGCM provides a mock function with given fields: cipherText
func (_m *IBlockCipher) DecryptWithAESGCM(cipherText []byte) ([]byte, error) {
	ret := _m.Called(cipherText)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(cipherText)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(cipherText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncryptWithAESGCM provides a mock function with given fields: plainText
func (_m *IBlockCipher) EncryptWithAESGCM(plainText []byte) ([]byte, error) {
	ret := _m.Called(plainText)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(plainText)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(plainText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCipherTextKey provides a mock function with given fields:
func (_m *IBlockCipher) GetCipherTextKey() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GetKMSKeyId provides a mock function with given fields:
func (_m *IBlockCipher) GetKMSKeyId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateEncryptionKey provides a mock function with given fields: _a0, ciphertextkey, sessionId
func (_m *IBlockCipher) UpdateEncryptionKey(_a0 log.T, ciphertextkey []byte, sessionId string) error {
	ret := _m.Called(_a0, ciphertextkey, sessionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T, []byte, string) error); ok {
		r0 = rf(_a0, ciphertextkey, sessionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
