pipeline {
  agent any
  stages {
    stage('Clone repository') {
      steps {
        git 'https://github.com/PabloAlarcon1/SoftServe_Demo3.git'
      }
    }

    stage('Compile') {
      steps {
        sh 'export GO111MODULE=on'
        sh 'go build -o myapp'
      }
    }

    stage('Testing') {
      steps {
        sh 'go test'
      }
    }

    stage('Execute') {
      steps {
        sh './myapp'
      }
    }

  }
}