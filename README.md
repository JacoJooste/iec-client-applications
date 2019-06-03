# iec-client-applications

Prepare the docker image that the client will be built and run in

    docker build . -t iec-client-applications:latest

Build the client application

    docker run --rm -it -v $(pwd):/root/iec iec-client-applications:latest bash -c "./build.sh"

Run the client application

    docker run --rm -it -v $(pwd)/dist:/root/iec iec-client-applications:latest bash -c "export LD_LIBRARY_PATH=/usr/local/lib && ./goclient"
