FROM golang:1.18-alpine as buildStage

WORKDIR /app 

#copy from current working directory to working directory in docer image 
COPY . .

#bulding our app to a single binary executable file  specify directory where main entry point is. in this case, ./ or main.go 
RUN GOOS=linux CGO_ENABLED=0 go build -o goSimpleBank ./

##alpine makes the docker image smaller , the smaller the better 😀
FROM alpine:latest 
WORKDIR /app

COPY --from=buildStage /app/goSimpleBank .
COPY app.env .
# COPY --from=buildStage /app/goSimpleBank /app

EXPOSE  8080
#command to executable that was built earlier
CMD ["/app/goSimpleBank"]