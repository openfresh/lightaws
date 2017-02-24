#!/usr/bin/env bash

if [ -z "$GITHUB_TOKEN" ]; then
  echo "GITHUB_TOKEN is empty." 1>&2
  exit 1
fi

if [ ! -z "$CIRCLE_TAG" ]; then
  TAG=$CIRCLE_TAG
fi

echo "TAG=$TAG"

UPLOAD_URL=`curl -s -X POST \
  -H "Authorization: token $GITHUB_TOKEN" \
  https://api.github.com/repos/openfresh/lightaws/releases/tags/$TAG | grep upload_url | sed -e "s/\s*\"upload_url\": \"//" | sed -e "s/\",//" | sed -e "s/{?name,label}/?name=/"`

if [ $? -ne 0 ]; then
  echo "create release failed." 1>&2
  exit 1
fi

for file in `ls artifacts`
do
  curl --data-binary @artifacts/$file \
    -H "Content-Type: application/x-gzip" \
    -H "Authorization: token $GITHUB_TOKEN" $UPLOAD_URL$file

  if [ $? -ne 0 ]; then
    echo "upload artifact failed." 1>&2
    exit 1
  fi
done
