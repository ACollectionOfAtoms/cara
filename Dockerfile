FROM bamos/openface AS cara
RUN pip install flask
COPY . /root/cara
EXPOSE 7777
