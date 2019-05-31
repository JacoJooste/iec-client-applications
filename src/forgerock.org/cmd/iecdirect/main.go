/*
 * Copyright 2019 ForgeRock AS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"stash.forgerock.org/iot/identity-edge-controller-core/am"
	"stash.forgerock.org/iot/identity-edge-controller-core/device"
	"stash.forgerock.org/iot/identity-edge-controller-core/json"
	"fmt"
	"stash.forgerock.org/iot/identity-edge-controller-core/logging"
	"os"
)

func main() {
	//ssoToken, result := am.Authenticate("Device", node.NodeDevice)
	//if result.Failure() {
	//	panic(result.String())
	//}
	//println("SSO Token: " + ssoToken)

	logging.Init(os.Stderr, os.Stdout)

	response, result := device.SendCommand("Narwhal", am.CommandRequest{
		Command: "GET_TOKENS",
		Params:  "",
	})
	if result.Failure() {
		panic(result.String())
	}
	println("Command response: " + response)

	var messageBody json.Object
	if json.Parse(&messageBody, []byte(response)) && messageBody.ContainsString("access_token") {
		token := messageBody.GetString("access_token")
		introspection := device.IntrospectToken(token)
		println("Token introspection: " + introspection)
		valid := device.ValidateToken(token)
		println(fmt.Sprintf("Token validation: %t", valid))
	}
}