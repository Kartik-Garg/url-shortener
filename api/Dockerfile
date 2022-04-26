FROM golang:alpine as builder
#getting golang docker image, getting alpine image

#creating a folder where it will be stored everytime the image is built
RUN mkdir /build
#adding it into the above created build folder
ADD . /build/
#setting the work directory
WORKDIR /build
#similar to go build command used for compilation, creates executable file
RUN go build -o main.

#stage 2
FROM alpine

RUN adduser -S -D -H -h /app appuser
#making user as appuser
USER appuser
#copy from current to app folder
COPY . /app
#copy from build/main to app folder , from tag here represents the stage
COPY --from=builder /build/main /app
#change working or current directory
WORKDIR /app

#expose the port
#port where the container will listen to
EXPOSE 3000

#run the following command at the end of the docker image
#basically executing the main.go file
CMD [ "./main" ]

