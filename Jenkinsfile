pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps { git url: 'https://github.com/Phantomvv1/Forms_site', credentialsId: 'Github-pat' }
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
