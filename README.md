# iec-client-applications

docker build . -t iec-client-applications:latest

docker run --rm -it -v $(pwd):/root/iec iec-client-applications:latest bash -c "./build.sh"

docker run --rm -it -v $(pwd):/root/iec iec-client-applications:latest bash -c "export LD_LIBRARY_PATH=/usr/local/lib && ./simplego"
