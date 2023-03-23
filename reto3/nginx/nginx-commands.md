## Update apt-get
1  sudo apt-get update

## Install Docker
2  sudo apt-get install docker.io -y
5  sudo usermod -aG docker $USER
6  exit

## Pull nginx image
7  docker pull nginx

## Create docker-compose.yml
8  touch docker-compose.yml
9  vim docker-compose.yml 

## Create nginx.conf
10  touch nginx.conf
11  vim nginx.conf 

## Install docker-compose
15  sudo apt install docker-compose

## Create the access.log file for nginx
51  mkdir /var/log/nginx
54  cd nginx/
56  sudo touch access.log

## Create the certificate folder for nginx
190  mkdir sslcerts
191  cd sslcerts/

## Create the certificates for nginx
192  openssl genrsa -out wildcard.key 2048
193  openssl req -new -key wildcard.key -out wildcard.csr
194  openssl x509 -req -days 365 -in wildcard.csr -signkey wildcard.key -out wildcard.crt