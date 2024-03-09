#!/usr/bin/env bash
source "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/base.env" set
source "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/$ENVIRONMENT.env" set
: "${IMAGE_LABEL:=latest}"
if [ -z "$IMAGE_TAG" ]; then
  IMAGE_TAG=`git log -n 1 --pretty=format:"%H"`
fi

if [ "$ENVIRONMENT" == "prod" ]; then
  IMAGE_TAG="main_$(git log -n 1 --pretty=format:"%H")"
else
  IMAGE_TAG="dev_$(git log -n 1 --pretty=format:"%H")"
fi
echo "using image tag: $IMAGE_TAG"  

ECR_REPOSITORY_URI=`aws ecr describe-repositories --repository-names "$REPO_NAME" 2>/dev/null | jq -r  .repositories[0].repositoryUri `

if [ -z "$ECR_REPOSITORY_URI" ]; then
  ECR_REPOSITORY_URI=`aws ecr create-repository --repository-name "$REPO_NAME" | jq -r .repository.repositoryUri`
  aws ecr put-lifecycle-policy --repository-name "$REPO_NAME" --lifecycle-policy-text "$(cat ./aws/ecr-lifecycle-policy-rules.json)"
fi
echo "repository URI: $ECR_REPOSITORY_URI"


aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin "$ACCOUNT_NUMBER.dkr.ecr.$REGION.amazonaws.com"
echo "docker build \
  -t $ECR_REPOSITORY_URI:$IMAGE_TAG \
  -t $ECR_REPOSITORY_URI:$IMAGE_LABEL ."

docker build \
  -t $ECR_REPOSITORY_URI:$IMAGE_TAG \
  -t $ECR_REPOSITORY_URI:$IMAGE_LABEL .

docker push $ECR_REPOSITORY_URI:$IMAGE_TAG
docker push $ECR_REPOSITORY_URI:$IMAGE_LABEL