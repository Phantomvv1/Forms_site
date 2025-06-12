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
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
		sh 'docker compose build'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker compose up'
            }
        }
    }
}
