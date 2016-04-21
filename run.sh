#!/usr/bin/env bash

export AWS_ACCESS_KEY_ID=$WERCKER_LAMBDA_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY=$WERCKER_LAMBDA_SECRET_KEY

PUBLISH=""
if [[ $WERCKER_LAMBDA_PUBLISH == "false" ]]; then
  $PUBLISH="--publish false"
fi

if [[ -z $WERCKER_LAMBDA_REGION ]]; then
  WERCKER_LAMBDA_REGION="us-east-1"
fi

LAMBDA="${WERCKER_STEP_ROOT}/lambda-deploy --functionname ${WERCKER_LAMBDA_FUNCTION_NAME} --region ${WERCKER_LAMBDA_REGION} --filepath ${WERCKER_LAMBDA_FILEPATH} ${PUBLISH}"
debug "$LAMBDA"
update_output=$($LAMBDA)

if [[ $? -ne 0 ]];then
    echo "${update_output}"
    fail 'Lambda update failed';
else
    echo "${update_output}"
    success 'Lambda update succeeded';
fi
