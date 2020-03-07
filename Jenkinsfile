pipeline {
	agent { 
		node {
			label 'my-defined-label'
			customWorkspace '/var/jenkins_home/go/src/github.com/200106-uta-go/project-3'
		}
	 } 
	 tools {
		 go 'Go'
	 }
	stages {
		stage('Checkout') {
			steps {
				sh 'git init'
				echo 'Checking out code...'
				checkout scm
			}
		}
		stage('Build') {
			steps {
				echo 'Building project...'
				sh 'go get -d ./...'
				sh 'go build ./...'
			}
		}
		stage('Test') {
			steps {
				echo 'Testing project...'
				sh 'go test ./...'
			}
		}
	}
	post {
		always {
				deleteDir()
		}
	}
}
