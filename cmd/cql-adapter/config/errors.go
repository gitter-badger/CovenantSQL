/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import "github.com/pkg/errors"

var (
	// ErrEmptyAdapterConfig defines empty adapter config.
	ErrEmptyAdapterConfig = errors.New("empty adapter config")
	// ErrRequireServerCertificate defines error of empty server certificate.
	ErrRequireServerCertificate = errors.New("require server certificate")
	// ErrInvalidStorageConfig defines error on incomplete storage config.
	ErrInvalidStorageConfig = errors.New("invalid storage config")
	// ErrInvalidCertificateFile defines invalid certificate file error.
	ErrInvalidCertificateFile = errors.New("invalid certificate file")
)
