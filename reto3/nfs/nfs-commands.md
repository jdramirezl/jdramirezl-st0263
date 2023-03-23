## Update apt-get
1  sudo apt-get update

## Install NFS
14  sudo apt-get install nfs-kernel-server

## Create a directory to share and give it the right permissions
114  sudo mkdir /mnt/wp-content
119  sudo chmod 777 /mnt/wp-content/

## Create a file to test the NFS share
120  cd /mnt/wp-content/
121  touch a.txt
122  vim a.txt 

## Edit /etc/exports to add the directory to share and the clients that can access it
126  sudo vim /etc/exports 

## Restart the NFS server to apply the changes
127  sudo systemctl restart nfs-kernel-server
