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
)

const (
	clientID = "india-client"
	deviceID = "india-device"
)

func main() {
	println("\n*** Running ", clientID, "\n")

	print("Initialising ", clientID, "... ")
	sdkConfig := configuration.SDKConfig{
		ZMQClient: configuration.ZMQClient{
			Endpoint: "tcp://172.16.0.11:5556",
			SecretKey: "RV/AxAZ.2&1[Ca$ha%w2o8S(NX@wqrSwhOTM+U}E",
			PublicKey: "0]#!mQZ]D?GBfjjzYw%{xLQl&FY.C.[bl!Ja=r4p",
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
		print("Initialisation request failed: " + result.String())
	}
	println("Done")

	fmt.Printf("Registering device (id: %s)... ", deviceID)
	result = zmqclient.DeviceRegister(deviceID, "{}")
	if result.Failure() {
		print("Registration request failed: " + result.String())
	}
	println("Done")

	println("\n*** Completed  ", clientID, "\n")
}