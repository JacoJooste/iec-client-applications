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
)

func main() {
	sdkConfig := configuration.SDKConfig{
		ZMQClient: configuration.ZMQClient{
			Endpoint: "tcp://172.16.0.11:5556",
			SecretKey: "zZZfS7BthsFLMv$]Zq{tNNOtd69hfoBsuc-lg1cM",
			PublicKey: "uH&^{aIzDw5<>TRbHcu0q#(zo]uLl6Wyv/1{/^C+",
			ServerPublicKey: "9m27tKf3aoNWQ(G-f[>W]gP%f&+QxPD:?mX*)hdJ",
			MessageResponseTimeoutSec: 5,
		},
		ClientConfig: configuration.ClientConfig{
			ID: "simple-go-client",
		},
		Logging: configuration.Logging{
			Enabled: true,
			Debug: true,
			Logfile: "simplego.log",
		},
	}
	result := zmqclient.Initialise(&sdkConfig)
	if result.Failure() {
		panic(result.String())
	}

	result = zmqclient.DeviceRegister("simple-gopher", "")
	if result.Failure() {
		panic(result.String())
	}

}