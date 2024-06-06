pipeline {
    agent any
    
    environment {
        SSH_PRIVATE_KEY = credentials('4227d049-af06-4906-ab84-71b312604ddc')
    }

    stages {
        stage('Clonar Repositorio') {
            steps {
                git credentialsId: '', url: 'git@github.com:imjesusleal/coreutils-go.git'
            }
        }
        stage('Build') {
            steps {
                sh 'make clean build release'
            }
        }
        stage('Crear Release') {
            steps {
                // Crear un release en GitHub
                script {
                    sh 'git tag -a v1.1.2 -m "Release v1.1.2"'
                    sh 'git push origin v1.1.2'
                }
            }
        }
    }
}
