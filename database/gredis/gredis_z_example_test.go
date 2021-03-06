// Copyright 2020 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gredis_test

import (
	"fmt"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
)

func Example_autoMarshalUnmarshalMap() {
	var (
		err    error
		result *gvar.Var
		key    = "user"
		data   = g.Map{
			"id":   10000,
			"name": "john",
		}
	)
	_, err = g.Redis().Do("SET", key, data)
	if err != nil {
		panic(err)
	}
	result, err = g.Redis().DoVar("GET", key)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Map())
}

func Example_autoMarshalUnmarshalStruct() {
	type User struct {
		Id   int
		Name string
	}
	var (
		err    error
		result *gvar.Var
		key    = "user"
		user   = &User{
			Id:   10000,
			Name: "john",
		}
	)

	_, err = g.Redis().Do("SET", key, user)
	if err != nil {
		panic(err)
	}
	result, err = g.Redis().DoVar("GET", key)
	if err != nil {
		panic(err)
	}

	var user2 *User
	if err = result.Struct(&user2); err != nil {
		panic(err)
	}
	fmt.Println(user2.Id, user2.Name)
}
