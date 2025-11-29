pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                script {
                    dockerImage = docker.build("test:${env.BUILD_NUMBER}")
                }
            }
        }
        stage('push to registry') {
            steps {
                script{
                    docker.withRegistry('http://registry:5000',''){
                        dockerImage.push()
                    }
                }
            }
        }
    }
}