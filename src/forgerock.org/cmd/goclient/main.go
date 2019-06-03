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
	"stash.forgerock.org/iot/identity-edge-controller-core/zmqclient"
	"stash.forgerock.org/iot/identity-edge-controller-core/configuration"
	"fmt"
	"stash.forgerock.org/iot/identity-edge-controller-core/json"
)

const (
	clientID = "go-client"
	deviceID = "gopher"
)

func main() {
	println("\n*** Running goclient\n")

	print("Initialising sdk... ")
	sdkConfig := configuration.SDKConfig{
		ZMQClient: configuration.ZMQClient{
			Endpoint: "tcp://127.0.0.1:5556",
			SecretKey: "zZZfS7BthsFLMv$]Zq{tNNOtd69hfoBsuc-lg1cM",
			PublicKey: "uH&^{aIzDw5<>TRbHcu0q#(zo]uLl6Wyv/1{/^C+",
			ServerPublicKey: "9m27tKf3aoNWQ(G-f[>W]gP%f&+QxPD:?mX*)hdJ",
			MessageResponseTimeoutSec: 5,
		},
		ClientConfig: configuration.ClientConfig{
			ID: clientID,
		},
		Logging: configuration.Logging{
			Enabled: true,
			Debug: true,
			Logfile: "goclient.log",
		},
	}
	result := zmqclient.Initialise(&sdkConfig)
	if result.Failure() {
		panic("Initialisation request failed: " + result.String())
	}
	println("Done")

	fmt.Printf("Registering device (id: %s)... ", deviceID)
	result = zmqclient.DeviceRegister(deviceID, "{}")
	if result.Failure() {
		panic("Registration request failed: " + result.String())
	}
	println("Done")

	fmt.Printf("Requesting configuration for device (id: %s)... ", deviceID)
	config, result := zmqclient.DeviceConfiguration(deviceID)
	if result.Failure() {
		panic("Configuration request failed: " + result.String())
	}
	println("Done")
	println("Received configuration: " + config)

	fmt.Printf("Requesting tokens for device (id: %s)... ", deviceID)
	tokens, result := zmqclient.DeviceTokens(deviceID)
	if result.Failure() {
		panic("Token request failed: " + result.String())
	}
	println("Done")
	println("Received tokens: " + tokens)

	print("Executing 'Hello World' custom command... ")
	helloResponse, result := zmqclient.CustomCommand(deviceID, "HELLO_WORLD", nil)
	if result.Failure() {
		panic("Custom command request failed: " + result.String())
	}
	println("Done")
	var helloJSON json.Object
	if json.Parse(&helloJSON, []byte(helloResponse)) && helloJSON.ContainsString("response") {
		println("Received response: " + helloJSON.GetString("response"))
	} else {
		println("Failed to extract AM response: " + helloResponse)
	}

	fmt.Printf("Requesting a user code for device (id: %s)... ", deviceID)
	userCode, result := zmqclient.UserCode(deviceID)
	if result.Failure() {
		panic("User code request failed: " + result.String())
	}
	println("Done")
	var userCodeJSON json.Object
	if json.Parse(&userCodeJSON, []byte(userCode)) && userCodeJSON.ContainsString("user_code") &&
		userCodeJSON.ContainsString("verification_url") {
		println(fmt.Sprintf("Extracted: {user_code: %s, verification_url: %s}", userCodeJSON.GetString("user_code"),
			userCodeJSON.GetString("verification_url")))
	} else {
		println("Failed to extract User Code values: " + userCode)
	}

	println("\n*** Completed goclient\n")
}