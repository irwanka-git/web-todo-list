FROM golang:1.22
LABEL maintainer="IRWAN KURNIAWAN <irwanka.emaiil@gmail.com>"
 
WORKDIR /app   
COPY go/ ./ 

ENV TZ=Asia/Jakarta
RUN go install github.com/cosmtrek/air@latest 
RUN go mod tidy
RUN air init 
CMD ["air"]