#!/usr/bin/env groovy
docker_oraclejdk8 {
    // This is not strictly required in this repo, but is easiest here so we only maintain this Kafka upstream trigger in one
    // place to trigger the rest of the pipeline. Note that we need to maintain the references to the jobs carefully since they
    // change across version branches.
    upstreamProjects = ['kafka-trunk', 'confluentinc/confluent-docker-utils']
    dockerRegistry = '368821881613.dkr.ecr.us-west-2.amazonaws.com/'
    dockerRepos = ['confluentinc/cp-base']
    slackChannel = '#clients-eng'
    dockerPush = true
}
