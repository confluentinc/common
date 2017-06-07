#!/usr/bin/env groovy

node('docker-openjdk7-wily') {
  stage('Preparation') {
    checkout scm
  }
  stage('Build') {
    sh "mvn --batch-mode -Pjenkins clean install dependency:analyze site"
  }
  stage('Deploy') {
    withMaven(
      globalMavenSettingsConfig: 'jenkins-maven-global-settings',
    ) {
      sh "mvn --batch-mode -Pjenkins -D${env.deployOptions} deploy -DskipTests"
    }
  }
  stage('Notify') {
    junit '**/target/surefire-reports/TEST-*.xml'
    step([$class: 'hudson.plugins.checkstyle.CheckStylePublisher', pattern: '**/target/checkstyle-result.xml', unstableTotalAll:'0'])
    step([$class: 'hudson.plugins.findbugs.FindBugsPublisher', pattern: '**/findbugsXml.xml'])
    archive 'target/*.jar'

    switch (currentBuild.currentResult) {
      case 'SUCCESS':
        if (currentBuild.previousBuild != null && currentBuild.previousBuild.currentResult != 'SUCCESS') {
          slackSend(channel: '#clients-eng', color: 'good', message: "${env.JOB_NAME} - #[${env.BUILD_NUMBER}] Success <${env.BUILD_URL}|(Open)>", teamDomain: 'confluent')
        }
        break;
      case 'UNSTABLE':
        slackSend(channel: '#clients-eng', color: 'YELLOW', message: "${env.JOB_NAME} - #[${env.BUILD_NUMBER}] Unstable <${env.BUILD_URL}|(Open)>", teamDomain: 'confluent')
        break;
      case 'FAILURE':
      default:
        slackSend(channel: '#clients-eng', color: 'bad', message: "${env.JOB_NAME} - #[${env.BUILD_NUMBER}] Failure <${env.BUILD_URL}|(Open)>", teamDomain: 'confluent')
        break;
    }

    if ($env.BRANCH_NAME.contains('/pull/')) {
      if (currentBuild.currentResult == 'SUCCESS') {
        githubNotify description: 'tests passed',  status: currentBuild.currentResult
      } else {
        githubNotify description: 'tests failed',  status: currentBuild.currentResult
      }
    }

  }
}