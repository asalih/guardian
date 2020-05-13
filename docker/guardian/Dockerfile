FROM golang

RUN go get -x github.com/asalih/guardian
WORKDIR /go/src/github.com/asalih/guardian
RUN make

ENV GUARDIAN_ENV=LIVE

RUN "sed -i 's,POSTGRES_PASSWORD,7eb12540045a4bc0b474a7efb91ebd39,g' appsettings.live.json"
RUN "sed -i 's,POSTGRES_HOST,localhost,g' appsettings.live.json"
RUN "sed -i 's,POSTGRES_USER,guardian,g' appsettings.live.json"
RUN "sed -i 's,POSTGRES_DB,guardiandb,g' appsettings.live.json"

CMD ["./guardian"]

EXPOSE 80 443