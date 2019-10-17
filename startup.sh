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

AM_URL="http://am.iec.com:8080/openam"

auth_reqeust="curl -s \
    -X POST \
    -H \"Content-Type: application/json\" \
    -H \"Accept-API-Version: resource=2.0, protocol=1.0\" \
    -H \"X-OpenAM-Username: amadmin\" \
    -H \"X-OpenAM-Password: password\" \
    ${AM_URL}/json/authenticate"
response=$(eval ${auth_reqeust})
token=$(echo ${response} | python -c 'import sys, json; print json.load(sys.stdin)[sys.argv[1]]' tokenId)

echo "Create a user called \"Charlie\""

curl -s -X PUT \
    -H "Content-Type: application/json" \
    -H "iPlanetDirectoryPro: ${token}" \
    -H "Accept-API-Version: protocol=2.1,resource=4.0" \
    -H "If-None-Match: *" \
    -d '{"userPassword":"password"}' \
    ${AM_URL}/json/realms/root/realms/edge/users/Charlie >/dev/null
