pipeline {
    agent any
    
    enviroment {
        GIT_SSH_KEY = credentials('4227d049-af06-4906-ab84-71b312604ddc')
    }
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


