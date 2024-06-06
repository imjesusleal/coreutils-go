pipeline {
    agent any

        stages {
            stage('Build') {
                steps {
                    sh 'make clean build rename'
                }
            }

            stage('Crear Release') {
                steps {
                    script {
                        sh 'git tag -a v1.1.2 -m "Release v1.1.2"'
                            sh 'git push origin v1.1.2'
                    }
                }
            }
        }

}

