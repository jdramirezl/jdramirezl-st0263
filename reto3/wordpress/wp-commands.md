## Pull Dockerfile 
1  sudo apt-get update
2  sudo apt-get install docker.io -y
3  sudo usermod -aG docker $USER
4  exit
5  docker pull wordpress

## Create docker-compose.yml
6  touch docker-compose.yml
8  sudo apt install docker-compose
13  vim docker-compose.yml

## Create NFS share
33  df -h
35  sudo apt install nfs-common
36  sudo mkdir /mnt/wp-content
37  sudo chmod 777 /mnt/wp-content/
38  sudo mount 10.128.0.6:/mnt/wp-content /mnt/wp-content

## Check NFS share
39  ls /mnt/wp-content/

## Run docker-compose
41  sudo docker-compose up

