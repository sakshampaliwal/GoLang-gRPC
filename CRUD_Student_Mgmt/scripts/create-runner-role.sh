#!/usr/bin/env bash
source "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/base.env" set
source "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/$ENVIRONMENT.env" set
ROLE_STACK_NAME="$PROJECT_NAME-runner-role"
STACK_EXISTS=`aws cloudformation describe-stacks --stack-name $ROLE_STACK_NAME 2>/dev/null`
COMMAND=update
if [ -z "$STACK_EXISTS" ]; then
  COMMAND=create
fi
aws cloudformation $COMMAND-stack --no-cli-pager \
  --stack-name $ROLE_STACK_NAME \
  --template-body file://aws/github-runner-role.yaml \
  --capabilities CAPABILITY_NAMED_IAM \
  --parameters \
  "ParameterKey=Environment,ParameterValue=$ENVIRONMENT" \
  "ParameterKey=ProjectName,ParameterValue=$PROJECT_NAME" \
  "ParameterKey=Region,ParameterValue=$REGION"\
  "ParameterKey=AccountNumber,ParameterValue=$ACCOUNT_NUMBER"\
  "ParameterKey=OidcProviderName,ParameterValue=$OIDC_PROVIDER_NAME"\
  "ParameterKey=GhRepoName,ParameterValue=$GH_REPO_NAME"\
  --tags Key=ProjectName,Value=$PROJECT_NAME --region $REGION
aws cloudformation wait stack-$COMMAND-complete --stack-name $ROLE_STACK_NAME
if [ $? -eq 0 ]; then
   echo "Stack $ROLE_STACK_NAME has been successfully $COMMAND"
else
  echo "stack $COMMAND failed."
  echo "open https://$REGION.console.aws.amazon.com/cloudformation/home?region=$REGION#/stacks in the browser to investigate."
fi