
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

## Copia por SSH
### HUE

Para subir los datos en HUE primero crearemos un directorio donde almacenaremos estos datasets. En mi caso es `/user/jdramirezl/datasets`.


Habiendo hecho esto ahora conseguiremos los datasets en este EC2 primario. Para esto yo descargue los datos en mi maquina local y los copie con el comando ` scp -i ~/reto5-keypair.pem -r ~/datasets  hadoop@ec2-34-237-52-15.compute-1.amazonaws.com:~`


Teniendo los archivos descargados en nuestra maquina local, ahora los podremos subir con ` hdfs dfs -copyFromLocal ~/datasets/* /user/jdramirezl/datasets/` ya que queria hacer la copia recursiva






### S3
## Copia por HDFS
### HUE
### S3
