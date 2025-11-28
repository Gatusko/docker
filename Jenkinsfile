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
                echo 'Deploying....'
            }
        }
    }
}