pipeline {
  agent {docker{image 'golang:alpine'}}

  environment {
    XDG_CACHE_HOME = "/tmp/.cache/go-build"
    CGO_ENABLED = 0
  }

  stages{
    stage('Build'){
      steps {
          sh 'mkdir -p ${GOPATH}/src/hydrophonics'
          sh 'cp -r ${WORKSPACE}/* /go/src/hydrophonics'
          sh 'cd /go/src/hydrophonics'
          sh 'go clean -cache'
          sh 'go build '
      }
    }

    stage('Test'){
      steps {
          sh 'mkdir -p ${GOPATH}/src/hydrophonics'
          sh 'cp -r ${WORKSPACE}/* /go/src/hydrophonics'
          sh 'cd /go/src/hydrophonics'
          sh 'go clean -cache'
          sh 'go test -v'
      }
    }

    stage('Review'){
      steps {
        sh 'echo TODO'
      }
    }
  }
}