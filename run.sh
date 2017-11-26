docker build -t cara . 
docker run -d \
	-v `pwd`:/root/workspace/src/github.com/ACollectionOfAtoms/cara \
	-p 7777:7777 \
	cara 
