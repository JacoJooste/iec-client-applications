#!/usr/bin/env bash
set -e

#
# Copyright 2019 ForgeRock AS
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#  http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.
#

docker run --network training_iec_net --rm -it -v $(pwd)/dist/india:/root/iec iec-client-applications:latest bash -c "export LD_LIBRARY_PATH=/usr/local/lib && ./india"
docker run --network training_iec_net --rm -it -v $(pwd)/dist/echo:/root/iec iec-client-applications:latest bash -c "export LD_LIBRARY_PATH=/usr/local/lib && ./echo"
