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

- A diferencia del codigo visto con MRJob, aqui corremos funciones aplicadas directamente sobre los archivos leidos en memoria. Ya que finalmente estos son parte de `sc`

  - flatMap: Aplicamos funcion sobre las lineas (Aqui split)
  - map: Ap;icamos funcion sobre los objetos (Aqui mapear cada palabra con 1)
  - reduceByKey: Reducir objetos dados bajo una funcion (Aqui, una funcion anonima que dadas dos entradas las suma y devuelve)
  - sortBy: Organizar bajo un criterio (Aqui dado un par en tupla, organizamos en reversa con el `-` bajo el segundo elemento a[1])
  - take: Tomamos una muestra de todo el output

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

- Como hemos vistos en labs anteriores, con el cluster de EMR ya corriendo podemos entrar a JupyterHub y crear un nuevo notebook con el kernel de `PySpark`. Una vez adentro podemos iniciar el kernel con `spark` y `sc`

- A partir de esto ya podemos leer el input con la funcion `textFile` y pasando la ruta desde S3 (En este caso `S3://reto5_notebooks/.../gutenberg-small/*.txt`)

- Realizamos el mismo codigo de ambos puntos anteriores sin imprimir

- Realizamos un output a S3 pero a una carpeta llamada `jupyter`

## Explicacion del notebook `Data_processing_using_PySpark`

- Este codigo nos permite entender como podemos manipular datos (datasets, etc) a traves de pyspark en el kernel de jupyer.
- Es manipulacion habitual de datos como se hace con scipy, pandas o numpy pero todo a traves de pyspark
- Encontramos maneras de agregar, quitar y derviar columnas. Tambien de encontrar medidas de centralizacion, etc.
- **EN EL JUPYTER EN ESTE MISMO DIRECTORIO SE ENCUENTRA LA EXPLICACION LINEA POR LINEA**

# Conclusiones

# Referencias
