// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"fmt"
	"strings"
)

// EmptyError represents an error that occurred during parsing.
type EmptyError struct{}

// Error returns the string representation of the error.
func (e *EmptyError) Error() string {
	return fmt.Sprint("Your SQL query is empty")
}

// QueryError represents an error that occured when switching access.
type QueryError struct{}

// Error returns the string representation of the error.
func (e *QueryError) Error() string {
	return fmt.Sprint("You don't have permission to perform this query type")
}

// BlankError represents an error that occured when switching access.
type BlankError struct{}

// Error returns the string representation of the error.
func (e *BlankError) Error() string {
	return fmt.Sprint("You need to specify a namespace and a database to use")
}

// PermsError represents an error that occured when switching access.
type PermsError struct {
	Resource string
}

// Error returns the string representation of the error.
func (e *PermsError) Error() string {
	return fmt.Sprintf("You don't have permission to access the '%s' resource", e.Resource)
}

// GroupError occurs when a 'group' expression is invalid.
type GroupError struct {
	found interface{}
}

// Error returns the string representation of the error.
func (e *GroupError) Error() string {
	return fmt.Sprintf("Found '%v' but field is not an aggregate function, and is not present in GROUP expression", e.found)
}

// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Found    string
	Expected []string
}

// Error returns the string representation of the error.
func (e *ParseError) Error() string {
	if len(e.Found) > 1000 {
		return fmt.Sprintf("Found `%s...` but expected `%s`", e.Found[:1000], strings.Join(e.Expected, ", "))
	}
	return fmt.Sprintf("Found `%s` but expected `%s`", e.Found, strings.Join(e.Expected, ", "))
}
