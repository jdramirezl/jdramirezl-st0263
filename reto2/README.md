# Info de la materia: ST0263 Topicos Especiales en Telematica
#
# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co
#
# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com


# nombre del proyecto, lab o actividad
Procesos comunicantes por API REST, RPC y MOM

# 1. breve descripción de la actividad
Este proyecto implementa un sistema distribuido con dos microservicios y un Gateway que funciona como Entrypoint (API), siguiendo el patrón BFF. 

El microservicio 1 se conecta por gRPC y el microservicio 2 por RabbitMQ. Todo el sistema es orquestado por Docker.

## 1.1. Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
- Implementar dos microservicios que ofrezcan servicios al Gateway.
- Los microservicios deben comunicarse por un middleware RPC y un middleware MOM.
- Los microservicios deben soportar concurrencia
- Usar gRPC para RPC.
- Usar RabbitMQ para MOM.
- Implementar consulta sobre los archivos de los procesos
- Implementar los servicios de listar archivos y buscar uno o más archivos.
- Implementar Round Robin para la seleccion del Microservicio a usar.
- Cada proceso debe tener un archivo de configuración dinamicamente leido.
- El archivo de configuración debe contener la IP, puerto y el directorio para buscar los archivos.
- Implementar un Gateway que exponga un API REST.

## 1.2. Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
- Ninguno

# 2. información general de diseño de alto nivel, arquitectura, patrones, mejores prácticas utilizadas.
La arquitectura general se ve a continuacion:
<p align="center">
<img src="https://user-images.githubusercontent.com/65835577/221764520-b30d16c1-7b11-4c62-9304-ac6f9b9f5740.png" alt="description of image" width="900" height="400" />
</p>

- Arquitectura: Microservicios. Cada microservicio se encarga de una funcionalidad
- Patrón: BFF (Backend for Frontend). Una API que se comunica con los microservicios y presenta los datos al frontend. 
- Los microservicios se comunican con el gateway con gRPC y RabbitMQ. 
    *   Los modelos de gRPC se definen en el archivo /internal/*.proto de gateway y micro1.
    *   Los mensajes de RabbitMQ son en formato JSON y son (Un)Marshalled en en cada llegada y salida.
- Se uso Golang como lenguaje de programación y Docker para orquestar el sistema.
- Mejores practicas: Uso de variables de entorno. Sistema de directorios y carpetas.


# 3. Descripción del ambiente de desarrollo y técnico: lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

- Docker: version 20.10.12, build 20.10 12-0ubuntu2~20.04.1
- Golang: go version go1.20.1 linux/amd64
- Protocol Buffer Compiler rotobuf: libprotoc 3.6.1
- RabbitMQ Docker Image: rabbitmq:3.8-management-alpine
- Librerias y paquetes:
    *   github.com/google/uuid v1.3.0
    *   github.com/rabbitmq/amqp091-go v1.7.0
	*   golang.org/x/net v0.7.0
	*   google.golang.org/grpc v1.53.0
	*   google.golang.org/protobuf v1.28.1
	*   github.com/golang/protobuf v1.5.2 // indirect
	*   golang.org/x/sys v0.5.0 // indirect
	*   golang.org/x/text v0.7.0 // indirect
	*   google.golang.org/genproto v0.0.*   
    *   google.golang.org/grpc v1.53.0


## como se compila y ejecuta.
El proyecto se puede compilar y ejecutar desde docker.
1. Descargar e instalar `Git`.
2. Crear una cuenta de Github
3. Clonar el proyecto desde Github mediante `git clone https://github.com/jdramirezl/jdramirezl-st0263/`
4. Descargar e instalar Docker.
5. Desde una terminal, dirigirse al directorio que tiene el archivo `docker-compose.yml`: `reto2`
6. Correr `docker-compose up --build` para construir y ejecutar el proyecto.

## detalles del desarrollo.
- Lenguaje de programación: Golang
- Control de versiones: Git
- 

## detalles técnicos
- Arquitectura/Patron: Microservicios y BFF
- Plataforma de nube: AWS
	* Servicios de nube usados: EC2 Ubuntu, Elastic IPs
- Herramientas de orquestación: Docker
- Comunicacion entre microservicios: gRPC y RabbitMQ

## descripción y como se configura los parámetros del proyecto
Cada servicio y el Gateway se pueden configurar a traves de un archivo .env que tienen en su carpeta `/config/`

Para modificarlos se hace necesario cambiar los valores despues el `=` que esta frente a cada variables 

Para cada uno se ve asi:

*Gateway*
```markdown
# Puerto y dirección IP donde el API Gateway estará escuchando
API_PORT=8080
API_IP=localhost

# Puerto y dirección IP donde el servidor GRPC está escuchando
GRPC_PORT=50051
GRPC_IP=grpc-server

# Puerto y dirección IP donde el servidor RabbitMQ está escuchando
RABBIT_PORT=5672
RABBIT_IP=rabbitmq

# Nombre de la cola en RabbitMQ donde se publicarán las solicitudes
RABBIT_QUEUE=requests
```

*Microservicio 1: RPC*
```markdown
# Dirección IP y puerto donde el servidor GRPC está escuchando
IP=grpc-server
PORT=50051

# Directorio donde se leeran/buscaran los archivos
DIRECTORY=/app/files
```

*Microservicio 2: RabbitMQ*
```markdown
# Dirección IP y puerto donde el servidor RabbitMQ está escuchando
IP=rabbitmq
PORT=5672

# Directorio donde se leeran/buscaran los archivos
DIRECTORY=/app/files

# Nombre de la cola en RabbitMQ donde se consumirán las solicitudes
QUEUE_NAME=requests
```

## Organización del código por carpetas o descripción de algún archivo.
![treeeee](https://user-images.githubusercontent.com/65835577/221765124-fe16adb8-194e-43c2-91d8-4801ac5ee25e.png)



# 4. Descripción del ambiente de EJECUCIÓN (en producción) lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.
El ambiente en produccion cuenta con las mismas versiones, lenguajes y librerias del ambiente de desarrollo.

# IP o nombres de dominio en nube o en la máquina servidor.
Se usa una IP elastica de AWS para acceder a la maquina EC2
- Gateway: 54.225.221.91:80

## descripción y como se configura los parámetros del proyecto
Las variables de ambiente son las mismas que las presentadas en desarrollo.

Para modificarlas se hace en el mismo directorio `config` mencionado anteriormente pero conectados a la maquina remota.

Una vez conectado, dirigirse al directorio `/home/ubuntu/project/reto2` donde encontraremos los archivos del proyecto.

En las carpetas de `gateway`, `micro1` y `micro2` encontraremos los archivos `.env` en el directiorio `config` que contienen las variables de ambiente.

Para las conexiones externas y entre servicios debemos tomar en cuenta que puertos se estan exponiendo y cuales no en el security group de la instancia de EC2.

En este como la salida a HTTP se hace por el puerto 80 entonces se hace necesario cambiar este valor en el archivo `.env` de `gateway`

## como se lanza el servidor.
1. Ingresar a AWS
2. Acceder a la seccion de EC2
3. Dar click en la instancia llamada `gateway`
4. Correr la instancia desde `Instance State` -> `Start`
5. Conectarse por alguna de las opciones que lista AWS (recomendado: `EC2 Instance Connect`)
6. Una vez conectado, dirigirse al directorio `/home/ubuntu/project/reto2` donde encontraremos los archivos del proyecto
7. Correr `docker-compose up`
8. Esperar a que se construyan y ejecuten los contenedores

## una mini guia de como un usuario utilizaría el software o la aplicación
Para conectarse al API Gateway se debe hacer una peticion HTTP a la IP elastica de AWS y al puerto 80. Ej:
- `http://54.225.221.91:80`

Aqui, mediante un `GET`, podemos acceder a dos endpoints: `/list` y `/search` 
- `http://54.225.221.91:80/list` -> lista todos los archivos que se encuentran en el directorio `/files` (o el que se indique) de cada microservicio
- `http://54.225.221.91:80/search?name=nombre_archivo` -> busca el archivo con el nombre `nombre_archivo` en el directorio `/files` (o el que se indique) de cada microservicio
	* Si usamos Postman podemos poner, en la seccion Params, `name` en el campo key y el nombre del archivo en el campo `value` y hacer la peticion

Como respuesta se obtiene un JSON con la informacion de los archivos encontrados.

# referencias:
- https://itnext.io/bff-pattern-with-go-microservices-using-rest-grpc-87d269bc2434?gi=2b74f1fa117c
- https://www.rabbitmq.com/tutorials/tutorial-one-go.html
- http://www.inanzzz.com/index.php/post/iamo/creating-a-rabbitmq-consumer-example-with-golang
- https://levelup.gitconnected.com/creating-a-minimal-rabbitmq-client-using-go-cbcec1470950
- https://grpc.io/docs/protoc-installation/
- https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
- https://www.youtube.com/watch?v=BdzYdN_Zd9Q&ab_channel=TutorialEdge

#### versión README.md -> 1.0 (2023-marzo)
