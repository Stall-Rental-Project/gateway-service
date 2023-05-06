pipeline {
    agent any
    environment {
                VERSION                = 'latest'
                PROJECT                = 'mhmarket'
                DOCKERHUB_CREDENTIALS  = credentials('dockerhub')
                ENVIRONMENT            = 'dev'
                SERVICE                = 'gateway'
    }
    stages {
        stage("Docker build") {
            steps {
                sh "docker build --network=host --tag 211020/${PROJECT}-${SERVICE}:${VERSION} ."
            }
        }
        stage("Docker Push") {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh "docker push 211020/${PROJECT}-${SERVICE}:${VERSION}"
            }
        }
        stage("Deploy") {
             steps {
                sh "docker-compose pull"
                sh "docker-compose down | echo IGNORE"
                sh "docker-compose up -d"
               }
        }
    }

}
