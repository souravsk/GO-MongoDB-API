#Using Golang as a base image
FROM golang:latest

#creating work dir
WORKDIR /app

#copy all the files to app folder which is on container
COPY . .

#Build the go projct
RUN go build -o main .

#Expose the application port
EXPOSE 4000

#Run the go app
CMD ["./main"]
