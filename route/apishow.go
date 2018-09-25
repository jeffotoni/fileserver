/*
* Go Library (C) 2017 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     FileServer
* @package     main
* @author      @jeffotoni
* @size        23/09/2018
*
 */

///////// start api
package route

import (
	"fmt"
)

import (
	config "github.com/jeffotoni/fileserver/config"
	"github.com/jeffotoni/fileserver/pkg/gcolor"
)

// Screen Logo
func ShowScreenMain() {

	fmt.Println(config.LOGO_FILESERVER)
}

//
// Mounting the properties on the api screen
//
func ShowScreen() {

	//
	// Basic Authentication
	//
	Login := Protocol + Schema + ":" + ServerPort + "" + HandlerLogin

	//
	//
	//
	Ping := Protocol + Schema + ":" + ServerPort + "" + HandlerPing

	//
	//
	//
	Hello := Protocol + Schema + ":" + ServerPort + "" + HandlerHello

	//
	// Basic Authentication, handler
	//
	CreateUser := Protocol + Schema + ":" + ServerPort + "" + HandlerCreateUser

	//
	// Basic Authentication, handler
	//
	Upload := Protocol + Schema + ":" + ServerPort + "" + HandlerUpload

	//
	//
	//
	UploadRemove := Protocol + Schema + ":" + ServerPort + "" + HandlerUploadRemoveDefinitive

	//
	//
	//
	UploadRemoveTrash := Protocol + Schema + ":" + ServerPort + "" + HandlerUploadRemoveFileTrash

	//
	//
	//
	Confirmemail := Protocol + Schema + ":" + ServerPort + "" + HandlerConfirmEmail

	//
	//
	//
	CloseAccount := Protocol + Schema + ":" + ServerPort + "" + HandlerCloseAccount

	//
	// HandlerRestoreAccount
	//
	RestoreAccount := Protocol + Schema + ":" + ServerPort + "" + HandlerRestoreAccount

	//
	// HandlerDisableUser
	//
	DisableUser := Protocol + Schema + ":" + ServerPort + "" + HandlerDisableUser

	//
	// HandlerDisableUser
	//
	EnableUser := Protocol + Schema + ":" + ServerPort + "" + HandlerEnableUser

	//
	//
	//
	Download := Protocol + Schema + ":" + ServerPort + "" + HandlerDownload

	//
	//
	//
	sizeMb := MaxHeaderByte

	//
	//
	//
	SizeString := fmt.Sprint("Max bytes: ", sizeMb, " Giga")

	// [GIN-debug] POST   /v1/login                 --> main.loginEndpoint (3 handlers)

	stringSchema := "⇨ http server started on " + Schema + ":" + ServerPort

	stringLogin := gcolor.CyanCor("[Ukk-debug] POST       ") + Login + gcolor.YellowCor("                 --> Login(2 Handlers)")

	stringPing := gcolor.CyanCor("[Ukk-debug] POST       ") + Ping + gcolor.YellowCor("                  --> Ping(2 Handlers)")

	stringHello := gcolor.CyanCor("[Ukk-debug] POST       ") + Hello + gcolor.YellowCor("                 --> Hello(2 Handlers)")

	stringCreateUser := gcolor.CyanCor("[Ukk-debug] POST       ") + CreateUser + gcolor.RedCor("                       --> CreateUser(2 Handlers)")

	stringUpload := gcolor.CyanCor("[Ukk-debug] POST       ") + Upload + gcolor.YellowCor("                --> Upload(2 Handlers)")

	stringUploadRemove := gcolor.CyanCor("[Ukk-debug] DELETE     ") + UploadRemove + gcolor.RedCor("                       --> UploadRemoveDefinitive(2 Handlers)")

	stringUploadRemoveTrash := gcolor.CyanCor("[Ukk-debug] DELETE     ") + UploadRemoveTrash + gcolor.RedCor("          --> UploadRemoveTrash(2 Handlers)")

	stringConfirmemail := gcolor.CyanCor("[Ukk-debug] POST       ") + Confirmemail + gcolor.RedCor("         --> Confirmemail(2 Handlers)")

	stringCloseAccount := gcolor.CyanCor("[Ukk-debug] POST       ") + CloseAccount + gcolor.RedCor("         --> CloseAccount(2 Handlers)")

	stringRestoreAccount := gcolor.CyanCor("[Ukk-debug] POST       ") + RestoreAccount + gcolor.RedCor("       --> RestoreAccount(2 Handlers)")

	stringDisableUser := gcolor.CyanCor("[Ukk-debug] PUT        ") + DisableUser + gcolor.RedCor("               --> DisableUser(2 Handlers)")

	stringEnableUser := gcolor.CyanCor("[Ukk-debug] PUT        ") + EnableUser + gcolor.RedCor("                --> EnableUser(2 Handlers)")

	stringDownload := gcolor.CyanCor("[Ukk-debug] POST       ") + Download + gcolor.YellowCor("              --> Download(2 Handlers)")

	//
	// Showing on the screen
	//
	gcolor.Green.Cprintln(stringSchema)

	fmt.Println(stringPing)

	fmt.Println("")

	fmt.Println(stringLogin)

	fmt.Println(stringHello)

	fmt.Println("")

	fmt.Println(stringCreateUser)

	fmt.Println(stringConfirmemail)

	fmt.Println(stringCloseAccount)

	fmt.Println(stringRestoreAccount)

	fmt.Println(stringDisableUser)

	fmt.Println(stringEnableUser)

	fmt.Println("")

	fmt.Println(stringUpload)

	fmt.Println(stringUploadRemove)

	fmt.Println(stringUploadRemoveTrash)

	fmt.Println(stringDownload)

	gcolor.Yellow.Cprintln(SizeString)

}
