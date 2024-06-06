pipeline {
    agent any
    
    triggers {
        githubPush()
    }

    environment {
        GITHUB_REPO = 'usuario/repositorio'
        GITHUB_CREDENTIALS_ID = 'github-credentials-id'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                sh 'make clean build'
            }
        }
        stage('Create GitHub Release') {
            steps {
                script {
                    def releaseName = "v${env.VERSION}"

                    withCredentials([usernamePassword(credentialsId: env.GITHUB_CREDENTIALS_ID, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                        githubRelease createRelease: true,
                                      releaseName: releaseName,
                                      tagName: releaseName,
                                      repoName: env.GITHUB_REPO,
                                      token: env.PASSWORD,
                                      body: "Release ${releaseName}"
                    }
                }
            }
        }
    }
}

