FROM golang

RUN go get -x github.com/asalih/guardian
WORKDIR /go/src/github.com/asalih/guardian
RUN make

ENV GUARDIAN_ENV=LIVE

RUN sed -i 's,POSTGRES_PASSWORD,7eb12540045a4bc0b474a7efb91ebd39,g' appsettings.live.json && \
    sed -i 's,POSTGRES_HOST,db,g' appsettings.live.json && \
    sed -i 's,POSTGRES_USER,guardian,g' appsettings.live.json && \
    sed -i 's,POSTGRES_DB,guardiandb,g' appsettings.live.json

RUN cp *.json workdir/ && \
    mkdir -p workdir/crs && \
    cp crs/*.data workdir/crs/ && \
    cp crs/*.conf workdir/crs

WORKDIR /go/src/github.com/asalih/guardian/workdir
CMD ["./guardian"]

EXPOSE 80 443