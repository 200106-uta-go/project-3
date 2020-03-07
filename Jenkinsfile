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
				sh 'su apt-get install build-essential -y'
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
				sh 'go test ./... -v'
			}
		}
	}
	post {
		failure {
			discordSend description: "Jenkins Pipeline Build", footer: "Build Failed", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "https://discordapp.com/api/webhooks/685969377430077573/a0Nno_j58sYJS3QScZIE7v45GZYfG8JDSbJT112cWFiuDfh_eu53XSI7u4JC5XO6Lgf0", image: "https://engineering.taboola.com/wp-content/uploads/2018/05/tidhar-post-featured-665x408.jpg"
		}
		always {
				echo 'Job complete, deleting directory...'
				deleteDir()
		}
	}
}
