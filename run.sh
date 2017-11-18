docker build -t cara . 
docker run -d \
	-v `pwd`:/root/cara \
	-p 7777:7777 \
	cara 

