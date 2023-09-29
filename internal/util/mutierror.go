/*
 *
 * Copyright 2023 @ Linying Assad <linying@apache.org>
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
 * /
 */

package util

import (
	"fmt"
	"strings"
)

// MultiError is a list of errors.
type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	errStrs := make([]string, 0, len(e.Errors))
	for _, err := range e.Errors {
		errStrs = append(errStrs, err.Error())
	}
	return strings.Join(errStrs, "; ")
}

// MergeErrors merges multiple errors into one.
func MergeErrors(errs ...error) *MultiError {
	errorList := make([]error, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			errorList = append(errorList, err)
		}
	}
	if len(errorList) == 0 {
		return nil
	}
	return &MultiError{Errors: errorList}
}

// MultiTaggedError is a list of errors with tags.
type MultiTaggedError struct {
	Errors map[string]error
}

func (e *MultiTaggedError) Error() string {
	errStrs := make([]string, 0, len(e.Errors))
	for tag, err := range e.Errors {
		errStrs = append(errStrs, fmt.Sprintf("[%s] %s", tag, err.Error()))
	}
	return strings.Join(errStrs, "; ")
}

// MergeErrorsWithTag merges multiple errors into one with tags.
func MergeErrorsWithTag(errors map[string]error) *MultiTaggedError {
	errMap := make(map[string]error, len(errors))
	for tag, err := range errors {
		if err != nil {
			errMap[tag] = err
		}
	}
	if len(errMap) == 0 {
		return nil
	}
	return &MultiTaggedError{Errors: errMap}

}
