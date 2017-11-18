docker build -t cara . # build here..
docker run -d \ # run detached
	-v `pwd`:/root/cara \ # sync volumes
	-p 7777:7777 \ # expose 7777
	cara # run the 'cara' image

