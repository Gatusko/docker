pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    dockerImage = docker.build("miApp:${env.BUILD_NUMBER}")
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