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
				checkout scm
			}
		}
		stage('Build') {
			steps {
				sh 'go build ./...'
			}
		}
		stage('Test') {
			steps {
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
