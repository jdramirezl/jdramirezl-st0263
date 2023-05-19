# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Reto 5 Parte 3 - 2

# Wordcount en Apache Spark EN AWS EMR

# Descripcion de la actividad

Realizar ejercicios de MAP-REDUCE en EMR con SPARK

## Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

- Ejecutar el wordcount por linea de comando 'pyspark' INTERACTIVO en EMR con datos en HDFS vía ssh en el nodo master.

- Ejecutar el wordcount por linea de comando 'pyspark' INTERACTIVO en EMR con datos en S3 (tanto de entrada como de salida)  vía ssh en el nodo master.

- Ejecutar el wordcount en JupyterHub Notebooks EMR con datos en S3 (tanto datos de entrada como de salida) usando un clúster EMR.

- Ejecutar el notebook `Data_processing_using_PySpark` con datos desde S3

- Explicar el notebook: Funcionamiento y salida

## Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

N/A

# Descripcion del ambiente de ejecucion

- AWS EMR 6.3.1
- Spark
- SSH
- Ubuntu 22.04

# Ejecucion de la actividad

Como siempre el primer paso es conectarnos por SSH a la instancia. Este paso ya lo tenemos gracias a las partes 1 y 2 del mismo reto.

Como los archivos ya estan en persistencia con S3, no los tenemos que pasar ahi, pero en HDFS tendremos que volverlos a subir.

Ya haciendo esto podemos comenzar con cada punto

## Wordcount interactivo con HDFS

- Al entrar y verificar que los archivos estan en HDFS corremos `pyspark`

  - Este nos va a permitir correr interactivamente cada linea

- Siendo asi tomamos el codigo de `wordcount` que se encuentra en el github de los retos y corremos linea por linea.

  - Podemos ir viendo el output en la linea del print que va mostrando ciertas palabras y sus cuentas
  - Podemos ver que el input esta en hdfs en la carpeta `user/hadoop/gutenberg-small` y tomamos todos los archivos txt con el operador `wildcard`
  - El output lo guardamos en `$USER/tmp/wcout1`

- Cada paso que ejecutamos nos da una serie de steps que representan el proceso de carga del codigo enviado

- Una vez terminado todo cerramos la sesion con `C-d` y mostramos la salida con un `ls` del hdfs

- Para revisar los resultados vamos a guardar en la maquina local del nodo el output de uno de estos archivos de `wcout1` y lo leemos con vim

- Una vez adentro, podemos ver el wordcout ejecutado

- Adentro de HUE podemos revisar que la carpeta `tmp/wcout1` existe y tiene los mismos archivos vistos por consola

## Wordcount interactivo por S3

- Igual que en el punto anterior (Y sabiendo que los datasets ya estan en S3) abrimos pyspark y comenzamos ejecucion

- El codigo a pesar de ser igual internamente cambia con el input y output

  - Sacamos los datos de S3 a traves de `S3://reto5_notebooks/.../*.txt`
  - Los guardamos en el mismo formato de `S3://reto5_notebooks/.../tmp/wcout1`

- Podemos revisar que la salida es correcta entrando al bucket y verificando que la carpeta se creo y por dentro tiene los archivos de resultado

- Tambien verificamos en el S3 de HUE

## Wordcount por jupyter por S3

## Explicacion del notebook `Data_processing_using_PySpark`

### Ejecucion del notebook en `jupyterhub` por S3

# Conclusiones

# Referencias
