# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Aplicación Monolítica con Balanceo y Datos Distribuidos (BD y archivos)

# 1. breve descripción de la actividad
La idea es desplegar un CMS Wordpress usando contenedores de Docker, con dominio y certificado SSL.

Hay que usar una instancia de nginx como load balancer, otra para la BDD y una ultima como FileServer de los datos distribuidos. 

La base de datos puede ser en Docker y el FileServer sera con NFS.
## 1.1. Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

- Desplegar un CMS Wordpress con contenedores
- Usar nginx como balanceador de cargas
- Crear una instancia de base de datos 
- Crear una instancia de archivos distribuidos
- Implementar un NFSServer en el FileServer
- Alta disponibilidad usando varios nodos
- Permitir el trafico de HTTPS
- Tener certificado SSL propio
- Tener 2 instancias de procesamiento Wordpress detras del Load balancer
- Tener 1 instancia de bases de datos mysql.

## 1.2. Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

N/A

# 2. información general de diseño de alto nivel

![Copia de tele p1 proyecto (1)](https://user-images.githubusercontent.com/65835577/227089038-dc7f636d-aa69-4860-b844-fc4979912b97.png)


- Arquitectura: Monolitica con varios nodos para alta disponibilidad
- Mejores practicas: Implementacion de contenedores, balanceo de cargas para alta disponibilidad con varios nodos, SSL con certificados propios, NFS para archivos distribuidos, uso de variables de entorno. Sistema de directorios y carpetas, notacion de archivos y carpetas.



# 3. Descripción del ambiente de EJECUCIÓN

- Contenedores en GCP con Ubuntu 22.24
- Wordpress image de docker
- Nginx image de docker
- MySQL image de docker
- nfs-kernel-server para el NFS
- OpenSSL para los certificados

# IP o nombres de dominio en nube o en la máquina servidor.

- Nginx: 10.128.0.2
- Wordpress 1: 10.128.0.3
- Wordpress 2: 10.128.0.4
- NFS: 10.128.0.6
- MySQL: 10.128.0.5

</br>

## Descripción y como se configura los parámetros del proyecto

Para modificar los parametros debemos hacerlo directamente en cada archivo listado.

Los parametros modificables para el proyecto son los siguientes:

</br>

### Nginx

- El docker-compose esta en el Root de la instancia de GCP
- Aqui podemos modificar:
    *  Ports: Los puertos de entrada  
    *  Volumes: Posicion de los certificados y del archivo de configuracion de nginx
```
    ports:
        - "80:80"
        - "443:443"
    volumes:
        - ./sslcerts/wildcard.crt:/jdramirezl.online.crt 
        - ./sslcerts/wildcard.key:/jdramirezl.online.key
        - ./sslcerts/wildcard.csr:/jdramirezl.online.csr 
        - ./nginx.conf:/etc/nginx/nginx.conf
```

</br>

- El archivo de configuracion de nginx esta en la carpeta Root de la instancia de GCP
- Aqui podemos modificar:
    * upstream backend: Las IP de redireccion
    * En el primer server: El puerto a escuchar y el dominio
    * En el segundo server: El puerto a escuchar y el dominio
    * En el segundo server: La ruta de los certificados

```
    upstream backend { # IP redirect with Round Robin
        server 10.128.0.3:80; # Wordpress 1
        server 10.128.0.4:80; # Wordpress 2
    }

    # Server HTTP
    listen 80; # Port 80
    server_name jdramirezl.online reto3.jdramirezl.online; # Domain name and alias

    # Server HTTPS
    listen 443 ssl; # Port 443 with SSL
    server_name jdramirezl.online reto3.jdramirezl.online; # Domain name and alias

    # SSL Certificates
    ssl_certificate /jdramirezl.online.crt;
    ssl_certificate_key /jdramirezl.online.key;
```
</br>

### Wordpress

- El docker-compose esta en el Root de la instancia de GCP
- Aqui podemos modificar:
    * Ports: Los puertos de entrada 
    * Environment: Las variables de entorno para la conexion con MySQL
    * Volume: La ruta del archivo de configuracion de Wordpress, que a la vez es un volumen de NFS
```
    ports:
        - "80:80" # Open port 80

    environment: # Set environment variables for the MySQL container
        WORDPRESS_DB_HOST: 10.128.0.5:3306
        WORDPRESS_DB_USER: jdramirezl
        WORDPRESS_DB_PASSWORD: 123
        WORDPRESS_DB_NAME: db

    volumes: # Mount the wp-content nfs shared folder from the host to the container
        - /mnt/wp-content:/var/www/html/wp-content
```
</br>

### MySQL
- El docker-compose esta en el Root de la instancia de GCP
- Aqui podemos modificar:
    * Ports: El puerto de entrada
    * environment: Las credenciales de la base de datos
```
    environment:
      MYSQL_ROOT_PASSWORD: 1234 
      MYSQL_DATABASE: db
      MYSQL_USER: jdramirezl
      MYSQL_PASSWORD: 123
    
    ports:
      - "3306:3306"
```
</br>

### NFS


- Aqui modificamos el archivo de configuracion de NFS (/etc/exports)
- Aqui podemos modificar:
    * La ruta de la carpeta a compartir
    * Las IP de los nodos que pueden acceder a la carpeta
    * Las opciones de la carpeta (Permiso de lectura y escritura, sincronizacion, no_subtree_check, etc)
```
    /mnt/wp-content 10.128.0.3(rw,sync,no_subtree_check) # IP 1
    /mnt/wp-content 10.128.0.4(rw,sync,no_subtree_check) # IP 2
```

## Como se lanza el servidor.

1. Ingresar a cada instancia de GCP mediante SSH
2. Correr `sudo docker-compose up -d` en cada una

## Una mini guia de como un usuario utilizaría el software o la aplicación

Con el servidor ya corriendo:

1. Ingresar a la pagina web reto3.jdramirezl.online
2. Usar Wordpress como de costumbre

## Resultados
![wordpress funcionando](https://user-images.githubusercontent.com/65835577/227088996-b96b7720-84d7-40d5-9695-b6e4eed67695.png)


# referencias:
[ssl ](https://www.namecheap.com/support/knowledgebase/article.aspx/9419/33/installing-an-ssl-certificate-on-nginx/)
[NFS](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nfs-mount-on-ubuntu-20-04)
[Nginx](https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04-es)
[load balancing](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/)
[Wordpress](https://www.digitalocean.com/community/tutorials/how-to-install-wordpress-with-docker-compose-es)
[MySQL](https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04-es)
[google cloud](https://cloud.google.com/vpc/docs/create-use-multiple-interfaces)
[URL para la informacion del proyecto](https://interactivavirtual.eafit.edu.co/d2l/le/content/122343/viewContent/601635/View)

#### versión README.md -> 1.0 (2023-marzo)

