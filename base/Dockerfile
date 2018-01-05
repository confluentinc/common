#
# Copyright 2017 Confluent Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM azul/zulu-openjdk-debian:8

ARG PROJECT_VERSION
ARG ARTIFACT_ID

# Remember where we came from
LABEL io.confluent.docker.git.repo="confluentinc/common"

ARG COMMIT_ID=unknown
LABEL io.confluent.docker.git.id=$COMMIT_ID

ARG BUILD_NUMBER=-1
LABEL io.confluent.docker.build.number=$BUILD_NUMBER

MAINTAINER tools@confluent.io
LABEL io.confluent.docker=true

# This affects how strings in Java class files are interpreted.  We want UTF-8 and this is the only locale in the
# base image that supports it
ENV LANG="C.UTF-8"

# Set the classpath for JARs required by `cub`
ENV CUB_CLASSPATH='"/usr/share/java/cp-base/*"'

COPY requirements.txt .

RUN apt update \
    && apt install -y bash curl git wget netcat-openbsd python python-pip python-setuptools python-enum34 \
    && pip install --install-option="--prefix=/usr/local" --upgrade -rrequirements.txt \
    && apt remove --purge -y git python-pip \
    && apt autoremove -y \
    && rm -rf /var/lib/apt/lists/*     \
    && rm requirements.txt \
    && mkdir -p /etc/confluent/docker

ADD target/${ARTIFACT_ID}-${PROJECT_VERSION}-package/share/doc/* /usr/share/doc/${ARTIFACT_ID}/
ADD target/${ARTIFACT_ID}-${PROJECT_VERSION}-package/share/java/${ARTIFACT_ID}/* /usr/share/java/${ARTIFACT_ID}/
