// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
	"strings"

	"github.com/casbin/casbin/rbac"
	"github.com/casbin/casbin/util"
)

// Assertion represents an expression in a section of the model.
// For example: r = sub, obj, act
type Assertion struct {
	Key    string
	Value  string
	Tokens []string
	Policy [][]string
	RM     rbac.RoleManager
}

func (ast *Assertion) buildRoleLinks(rmc rbac.RoleManagerConstructor) {
	ast.RM = rmc()
	count := strings.Count(ast.Value, "_")
	for _, rule := range ast.Policy {
		if count < 2 {
			panic(errors.New("the number of \"_\" in role definition should be at least 2"))
		}
		if len(rule) < count {
			panic(errors.New("grouping policy elements do not meet role definition"))
		}

		if count == 2 {
			ast.RM.AddLink(rule[0], rule[1])
		} else if count == 3 {
			ast.RM.AddLink(rule[0], rule[1], rule[2])
		} else if count == 4 {
			ast.RM.AddLink(rule[0], rule[1], rule[2], rule[3])
		}
	}

	util.LogPrint("Role links for: " + ast.Key)
	ast.RM.PrintRoles()
}