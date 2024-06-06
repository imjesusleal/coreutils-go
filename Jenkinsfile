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
                        withCredentials([usernamePassword(credentialsId: 'e4b9a7a1-5ad0-49b7-977b-58ae06b9cf26', usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD')]) {
                            sh '''
                                git config --global user.name "${GIT_USERNAME}"
                                git config --global user.password "${GIT_PASSWORD}"
                                git push --set-upstream origin main
                                '''
                        }
                        sh 'git tag -a v1.1.2 -m "Release v1.1.2"'
                            sh 'git push origin v1.1.2'
                    }
                }
            }
        }
}


