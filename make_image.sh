docker build -t survey_backend .

#docker run --env=PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOLANG_VERSION=1.20.5 --env=GOPATH=/go --workdir=/go/src/survey_backend -p 8080:8080 --restart=no --runtime=runc -d --name survey_backend survey_backend