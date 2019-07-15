# Client attestation simulation

AM and IEC Service must be started from the training environment with the latest snapshot binaries.

Prepare the docker image that the client will be built and run in

    docker build . -t iec-client-applications:latest

Build the client application

    docker run --rm -it -v $(pwd):/root/iec iec-client-applications:latest bash -c "./build-demo.sh"

Modify the [IEC configuration](http://am.iec.com:8080/identitymanager/#iec) by adding the following keys to "zmq_server.client_public_keys"

    "m.DH?j1Y}g0I!V)6v]a*n{!up@YO/*UVdKW@Ji54",
    "0]#!mQZ]D?GBfjjzYw%{xLQl&FY.C.[bl!Ja=r4p"

Restart the IEC Service by running the following in the iot-gateway terminal:

    stop_iec_service && start_iec_service

Run the client application

    ./run-demo.sh
