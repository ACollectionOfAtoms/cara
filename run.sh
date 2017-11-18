docker build -t cara .
docker run -d \
	-v `pwd`:/root/cara \
	cara

