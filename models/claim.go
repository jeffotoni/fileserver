/*
* Ukkbox Go Library (C) 2017 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Ukkbox
* @package     main
* @author      @jeffotoni
* @size        01/06/2017
*
 */

package models

//
// Claim structure, where we will use
// to validate our token with jwt
//
import "github.com/dgrijalva/jwt-go"

//
// jwt
//
type Claim struct {
	User   string `json:"user"`
	Uid    string `json:"uid"`
	Uidwks string `json:"uidwks"`

	jwt.StandardClaims
}
