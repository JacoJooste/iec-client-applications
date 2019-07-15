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
	"github.com/dgrijalva/jwt-go"
	"strings"
)

const (
	clientID = "echo-client"
	peerClientID = "india-client"
	peerDeviceID = "india-device"
)

func main() {
	println("\n*** Running ", clientID, "\n")

	print("Initialising ", clientID, "... ")
	sdkConfig := configuration.SDKConfig{
		ZMQClient: configuration.ZMQClient{
			Endpoint: "tcp://172.16.0.11:5556",
			SecretKey: "Y5&eKAL^%M7YMWHDh#!mdclMAxzW-&.4vSMH4?mq",
			PublicKey: "m.DH?j1Y}g0I!V)6v]a*n{!up@YO/*UVdKW@Ji54",
			ServerPublicKey: "9m27tKf3aoNWQ(G-f[>W]gP%f&+QxPD:?mX*)hdJ",
			MessageResponseTimeoutSec: 5,
		},
		ClientConfig: configuration.ClientConfig{
			ID: clientID,
		},
		Logging: configuration.Logging{
			Enabled: true,
			Debug: true,
			Logfile: clientID + ".log",
		},
	}
	result := zmqclient.Initialise(zmqclient.UseDynamicConfig(sdkConfig))
	if result.Failure() {
		panic("Initialisation request failed: " + result.String())
	}
	println("Done")

	fmt.Printf("Requesting configuration for device (id: %s)... ", peerDeviceID)
	config, result := zmqclient.DeviceConfiguration(peerDeviceID)
	if result.Failure() {
		panic("Configuration request failed: " + result.String())
	}
	println("Done")
	println("Received configuration: " + config)

	fmt.Printf("Requesting tokens for device (id: %s)... ", peerDeviceID)
	tokens, result := zmqclient.DeviceTokens(peerDeviceID)
	if result.Failure() {
		panic("Token request failed: " + result.String())
	}
	println("Done")
	var tokenObject json.Object
	if !json.Parse(&tokenObject, []byte(tokens)) || !tokenObject.ContainsString("id_token") {
		panic("Failed to unpack ID Token from device tokens")
	}

	parts := strings.Split(tokenObject.GetString("id_token"), ".")
	claims, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		panic("Failed to parse ID Token: " + err.Error())
	}
	println("Received ID token: ", string(claims))

	fmt.Printf("Requesting configuration for client (id: %s)... ", peerClientID)
	config, result = zmqclient.CustomCommand(peerClientID, "GET_CLIENT_CONFIG", nil)
	if result.Failure() {
		panic("Configuration request failed: " + result.String())
	}
	println("Done")
	println("Received configuration: " + config)

	fmt.Printf("Requesting tokens for client (id: %s)... ", peerClientID)
	tokens, result = zmqclient.DeviceTokens(peerClientID)
	if result.Failure() {
		panic("Token request failed: " + result.String())
	}
	println("Done")
	if !json.Parse(&tokenObject, []byte(tokens)) || !tokenObject.ContainsString("id_token") {
		panic("Failed to unpack ID Token from device tokens")
	}

	parts = strings.Split(tokenObject.GetString("id_token"), ".")
	claims, err = jwt.DecodeSegment(parts[1])
	if err != nil {
		panic("Failed to parse ID Token: " + err.Error())
	}
	println("Received ID token: ", string(claims))

	println("\n*** Completed  ", clientID, "\n")
}