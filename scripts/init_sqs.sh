#!/bin/bash

QUEUE_ENDPOINT=$1

AWS_ACCESS_KEY_ID=dummy AWS_SECRET_ACCESS_KEY=dummy aws sqs create-queue \
    --endpoint-url ${QUEUE_ENDPOINT} \
    --queue-name test \
    --attributes ReceiveMessageWaitTimeSeconds=10 \
    --region ap-northeast-1

AWS_ACCESS_KEY_ID=dummy AWS_SECRET_ACCESS_KEY=dummy aws sqs create-queue \
    --endpoint-url ${QUEUE_ENDPOINT} \
    --queue-name test \
    --attributes ReceiveMessageWaitTimeSeconds=10 \
    --region ap-northeast-3
