pipeline {
	agent { docker { image 'golang' } } 
	stages {
		stage('Checkout') {
			checkout scm
		}
		stage('Build') {
			steps {
				sh 'go version'
			}
		}
	}
}
