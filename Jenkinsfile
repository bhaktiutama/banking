pipeline {
    environment {
        CREDS = credentials('mysql-db-cred')
    }
    agent any
    stages {
        // stage('Preparation') {
        //     // Preparation command
        // }
        stage('DB Migration') {
            when {
                tag '*with-db*'
            }
            steps {
                sh 'migrate -path db -database "mysql://%CREDS_USR%:%CREDS_PSW%@tcp(localhost:3306)/test" -verbose up'
            }
        }
        // stage('Build') {
        //     //Build command     
        // }
        // stage('Send Slack') {
        //     // Send slack
        // }
    }
}