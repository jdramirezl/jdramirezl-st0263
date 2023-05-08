
# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Reto 5 Parte 2
# GESTIÓN DE ARCHIVOS EN HDFS Y S3 PARA BIG DATA

# Descripcion de la actividad
Aprender a subir archivos al sistema no persistente de HDFS y el sistema persistente S3 a traves de HUE y SSH

## Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
* Copiar (gestión) de archivos hacia AWS S3 vía HUE.
* Copiar (gestión) de archivos hacia el AWS S3 vía SSH.
* Copiar (gestión) de archivos hacia el HDFS vía HUE.
* Copiar (gestión) de archivos hacia el HDFS vía SSH.

## Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
N/A

# Descripcion del ambiente de ejecucion
* AWS EMR 6.3.1
* HUE
* SSH 
* Ubuntu 22.04

# Ejecucion de la actividad

Antes de empezar vamos a hacer dos cosas:

1. Para subir los datos en HUE primero crearemos un directorio donde almacenaremos estos datasets. En mi caso es `/user/jdramirezl/datasets`.

2. Conseguiremos los datasets en este EC2 primario. Para esto yo descargue los datos en mi maquina local y los copie con el comando ` scp -i ~/reto5-keypair.pem -r ~/datasets  hadoop@ec2-34-237-52-15.compute-1.amazonaws.com:~`

## Copia por SSH
### HDFS

Teniendo los archivos descargados en nuestra maquina local, ahora los podremos subir con ` hdfs dfs -copyFromLocal ~/datasets/* /user/jdramirezl/datasets/` ya que queria hacer la copia recursiva


### S3

Como el EC2 Primario se encuentra en la red de EMR podemos usar la libreria `aws s3`. Con esto entonces hacemos `aws s3 cp ~/datasets/ s3://reto5-notebooks --recursive`

## Copia por HUE
### HDFS

Primero ingresamos a la pestana `Files`

Como quedamos en el usuario Hadoop nos moveremos al directorio con nuestro username (En mi caso `jdramirezl`).

En el boton `new` podemos crear un nuevo directorio. En mi caso yo lo llame `datasets-hue`

Hue no nos permite subir carpetas, por lo que tendremos que crear cada directorio individualmente

Y luego en cada directorio subir los archivos manualmente.

### S3

Primero ingresamos a la pestana `S3`

Ahi ya estaremos adentro del Bucket. En mi caso `reto5-notebooks`. Aqui crearemos un folder para guardar los datos, en mi caso `datasets-hue`

Aqui tendremos que repetir la misma operacion como con HDFS, donde creamos cada directorio individualmente y le subimos sus respectivos archivos


# Resultados
Podemos encontrar los resultados por medio de las imagenes subidas anteriormente, pero tambien se adjuntan los archivos en el bucket de s3 dentro de AWS




# Referencias
[Enunciado del reto](https://github.com/st0263eafit/st0263-231/blob/main/bigdata/lab5-2-hdfs-s3.txt)
[Explicacion HDFS en HUE y SSH](https://github.com/st0263eafit/st0263-231/tree/main/bigdata/01-hdfs)
[Dataset subido](https://github.com/st0263eafit/st0263-231/tree/main/bigdata/datasets)
[copy files to s3 bucket through ssh](https://www.middlewareinventory.com/blog/ec2-s3-copy/)
[AWS S3 Invalid length for parameter Key - Subir archivos recursivos a S3](https://storiesbynazreen.medium.com/quick-debug-aws-s3-invalid-length-for-parameter-key-4e07359b396d)
