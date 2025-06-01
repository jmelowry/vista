#!/bin/zsh

echo "Testing API endpoints..."

echo "\n1. Testing /repos endpoint:"
curl -s http://localhost:8080/repos | jq

echo "\n2. Testing /repo/ecr-main endpoint:"
curl -s http://localhost:8080/repo/ecr-main | jq

echo "\n3. Testing /repo/ecr-main/resources endpoint:"
curl -s http://localhost:8080/repo/ecr-main/resources | jq

echo "\n4. Testing /repo/ecr-main/resource/my-app endpoint:"
curl -s http://localhost:8080/repo/ecr-main/resource/my-app | jq

echo "\n5. Testing /repo/dockerhub/resources endpoint:"
curl -s http://localhost:8080/repo/dockerhub/resources | jq

echo "\n6. Testing /repo/dockerhub/resource/nginx endpoint:"
curl -s http://localhost:8080/repo/dockerhub/resource/nginx | jq

echo "\nTests completed."
