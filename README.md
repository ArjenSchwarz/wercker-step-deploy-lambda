# Wercker step for deploying to Lambda functions

[![wercker status](https://app.wercker.com/status/b97259cecb4dd1a4d20806fd2b2934a8/m "wercker status")](https://app.wercker.com/project/bykey/b97259cecb4dd1a4d20806fd2b2934a8)

This step will easily allow you to deploy your lambda functions using Wercker.

# AWS Configuration

You will require access keys for an AWS IAM user with sufficient privileges to be set up. A user having only the below policy is sufficient, and therefore recommended.

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1461228078000",
            "Effect": "Allow",
            "Action": [
                "lambda:PublishVersion",
                "lambda:UpdateFunctionCode"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```

# Parameters

## access_key

The AWS access key from the above mentioned user. Set this up as a variable in your deployment target.

## secret_key

The AWS secret key from the above mentioned user. Set this up as a variable in your deployment target.

## function_name

The name of the Lambda function you want to update.

## filepath

The path of the zip file you want to push up to the Lambda function. This will usually be a path like "$WERCKER_SOURCE_DIR/lambda.zip"

## region

The AWS region where your Lambda function is to be found. Defaults to "us-east-1"

## publish

A flag whether to publish the uploaded function or not. Defaults to true, disable that behavior with `publish: false`.

# Example

```yml
deploy:
    steps:
        - arjen/lambda:
            access_key: $AWS_ACCESS_KEY
            secret_key: $AWS_SECRET_KEY
            function_name: myfunction
            filepath: $WERCKER_SOURCE_DIR/myfunction.zip
```
