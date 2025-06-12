pipeline {
    agent any
    stages {
        stage('Debug') {
            steps {
                sh 'git branch -a'
                sh 'git status'
                sh 'git remote -v'
            }
        }
        stage('Checkout') {
            steps { git url: 'https://github.com/Phantomvv1/Forms_site', credentialsId: '46f1d746-664e-4009-9e7f-bd01c49ee4fc' }
        }
        stage('Build') {
            steps {
		sh 'docker-compose build'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker-compose up'
            }
        }
    }
}
