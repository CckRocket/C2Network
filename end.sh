sudo docker-compose -f docker-orderer.yaml down
sudo docker-compose -f docker-peer.yaml down
sudo docker stop $(sudo docker ps -a -q)
sudo docker rm $(sudo docker ps -a -q)
