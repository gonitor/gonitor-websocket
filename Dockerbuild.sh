echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t gonitor/gonitor-websocket .
docker tag  gonitor/gonitor-websocket gonitor/gonitor-websocket:0.0.0
docker push gonitor/gonitor-websocket