CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t gonitor/gonitor-websocket .
docker tag  gonitor/gonitor-websocket gonitor/gonitor-websocket:0.0.0