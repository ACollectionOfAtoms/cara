FROM bamos/openface AS cara
RUN wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
RUN tar -C /usr/local/ -xzf go1.9.2.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/root/workspace/"
ENV OPENFACE_PATH="/root/openface/"

RUN go get github.com/rs/cors
COPY . /root/workspace/src/github.com/ACollectionOfAtoms/cara/
EXPOSE 7777
