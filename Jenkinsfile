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
        sh 'go build'
      }
    }
    
    stage('Testing') {
      steps {
        sh 'go test'
      }
    }
    
    stage('Execute') {
      steps {
        sh 'go run main.go'
      }
    }
  }
}