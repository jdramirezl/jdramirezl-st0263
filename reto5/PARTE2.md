
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

1. Para subir los datos en HUE primero crearemos un directorio donde almacenaremos estos datasets. En mi caso fue con `hdfs dfs -mkdir /user/jdramirezl/datasets`.

2. Conseguiremos los datasets en este EC2 primario. Para esto yo descargue los datos en mi maquina local y los copie con el comando ` scp -i ~/reto5-keypair.pem -r ~/datasets  hadoop@ec2-34-237-52-15.compute-1.amazonaws.com:~`

![data cp y](https://user-images.githubusercontent.com/65835577/236962593-50747b83-2d04-4098-a0b6-b74fb6712e30.png)

## Copia por SSH
### HDFS

Teniendo los archivos descargados en nuestra maquina local, ahora los podremos subir con ` hdfs dfs -copyFromLocal ~/datasets/* /user/jdramirezl/datasets/` ya que queria hacer la copia recursiva. Este fue el resultado

![ssh hdfs](https://user-images.githubusercontent.com/65835577/236962725-1827c83c-d569-44f6-882d-f1bdb81abf9b.png)

Y el resultado visualizado en HUE

![ssh hdfs evidencia](https://user-images.githubusercontent.com/65835577/236962869-b7697fb6-e28b-489e-a04e-cdbf32dfc745.png)


### S3

Como el EC2 Primario se encuentra en la red de EMR podemos usar la libreria `aws s3`. Con esto entonces hacemos `aws s3 cp ~/datasets/ s3://reto5-notebooks --recursive`. Este es el resultado en el bucket con los objetos copiados

![ssh s3](https://user-images.githubusercontent.com/65835577/236962837-ca22d800-1fd5-42f9-9136-73f153082b3e.png)

## Copia por HUE
### HDFS

Primero ingresamos a la pestana `Files`

![hue upload files](https://user-images.githubusercontent.com/65835577/236962880-525b6fc7-e658-45a2-8a7b-68323ec07a8e.png)

Como quedamos en el usuario Hadoop nos moveremos al directorio con nuestro username (En mi caso `jdramirezl`).

![entramos a jdramirezl](https://user-images.githubusercontent.com/65835577/236962896-d34cf7f2-0f6a-40c4-a922-7550fcf38c99.png)

En el boton `new` podemos crear un nuevo directorio. En mi caso yo lo llame `datasets-hue`

![nuevo folder gue](https://user-images.githubusercontent.com/65835577/236962922-b500cd06-41c1-42ed-9515-a3052048f97e.png)

![carpeta creada hue](https://user-images.githubusercontent.com/65835577/236962909-41ecc031-74d0-4f69-a9f2-d3614a369c27.png)

Hue no nos permite subir carpetas, por lo que tendremos que crear cada directorio individualmente

![carpetas creadas hue hdfs](https://user-images.githubusercontent.com/65835577/236962936-eb63b44e-ab93-42e1-9c09-5cadc2514c06.png)

Y luego en cada directorio subir los archivos manualmente.

![archivos subidos hue hdfs](https://user-images.githubusercontent.com/65835577/236962953-e3574866-8358-48f6-b4c9-45dc6444212f.png)

### S3

Primero ingresamos a la pestana `S3`

![pestana s3](https://user-images.githubusercontent.com/65835577/236962976-1fcd113c-c80f-4fd2-b5a2-50cdb8ef2808.png)

Ahi ya estaremos adentro del Bucket. En mi caso `reto5-notebooks`. Aqui crearemos un folder para guardar los datos, en mi caso `datasets-hue`

![datasets hue s3 ](https://user-images.githubusercontent.com/65835577/236962994-626f44c6-368a-40b9-be42-00399d50b677.png)

Aqui tendremos que repetir la misma operacion como con HDFS, donde creamos cada directorio individualmente y le subimos sus respectivos archivos

![S3 DATASTS HUE LLENADOS](https://user-images.githubusercontent.com/65835577/236963036-ab35c81a-1f51-4ea4-939e-13d695f3eea3.png)

Y si revisamos en el bucket, si se creo la carpeta con archivos

![s3  hue carpeta creada](https://user-images.githubusercontent.com/65835577/236963085-1b122ab2-5e45-442d-a399-67eaf0e9f1b6.png)



# Resultados
Podemos encontrar los resultados por medio de las imagenes subidas anteriormente, pero tambien se adjuntan los archivos en el bucket de s3 dentro de AWS ([URL a los archivos del bucket](https://reto5-notebooks.s3.amazonaws.com/datasets-hue/))

![s3 llenado ](https://user-images.githubusercontent.com/65835577/236963089-a1a52b14-2158-492f-9187-d6cb085277e0.png)



# Referencias
[Enunciado del reto](https://github.com/st0263eafit/st0263-231/blob/main/bigdata/lab5-2-hdfs-s3.txt)
[Explicacion HDFS en HUE y SSH](https://github.com/st0263eafit/st0263-231/tree/main/bigdata/01-hdfs)
[Dataset subido](https://github.com/st0263eafit/st0263-231/tree/main/bigdata/datasets)
[copy files to s3 bucket through ssh](https://www.middlewareinventory.com/blog/ec2-s3-copy/)
[AWS S3 Invalid length for parameter Key - Subir archivos recursivos a S3](https://storiesbynazreen.medium.com/quick-debug-aws-s3-invalid-length-for-parameter-key-4e07359b396d)
[Como hacer un bucket de S3 publico](https://bobbyhadz.com/blog/aws-s3-allow-public-read-access)
