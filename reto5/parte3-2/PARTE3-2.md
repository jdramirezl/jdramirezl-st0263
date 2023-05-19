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
    ![abrir pyspark](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/9cc7f415-5e51-4a56-8b70-ca06feb271fc)

- Siendo asi tomamos el codigo de `wordcount` que se encuentra en el github de los retos y corremos linea por linea.

  - Podemos ir viendo el output en la linea del print que va mostrando ciertas palabras y sus cuentas
  - Podemos ver que el input esta en hdfs en la carpeta `user/hadoop/gutenberg-small` y tomamos todos los archivos txt con el operador `wildcard`
  - El output lo guardamos en `$USER/tmp/wcout1`
    ![ejecutando hdfs interactivo](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/1fd432b1-b326-42db-a0d0-7a210370faa5)

- A diferencia del codigo visto con MRJob, aqui corremos funciones aplicadas directamente sobre los archivos leidos en memoria. Ya que finalmente estos son parte de `sc`

  - flatMap: Aplicamos funcion sobre las lineas (Aqui split)
  - map: Ap;icamos funcion sobre los objetos (Aqui mapear cada palabra con 1)
  - reduceByKey: Reducir objetos dados bajo una funcion (Aqui, una funcion anonima que dadas dos entradas las suma y devuelve)
  - sortBy: Organizar bajo un criterio (Aqui dado un par en tupla, organizamos en reversa con el `-` bajo el segundo elemento a[1])
  - take: Tomamos una muestra de todo el output

- Cada paso que ejecutamos nos da una serie de steps que representan el proceso de carga del codigo enviado
  ![steps spark](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/80d493b0-424f-4784-af50-68f352284870)

- Una vez terminado todo cerramos la sesion con `C-d` y mostramos la salida con un `ls` del hdfs
  ![resultado hdfs interactivo](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/c0ba43f5-c99f-4a60-883f-6f2e26244c15)
  ![resultado 2 hdfs interactivo](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/6e780557-fe54-4331-8b08-d65307c7d17e)

- Para revisar los resultados vamos a guardar en la maquina local del nodo el output de uno de estos archivos de `wcout1` y lo leemos con vim
  ![guardamos resultado y abrimos en vim](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/fb8802f6-501d-4c6e-ae4c-42936146e169)

- Una vez adentro, podemos ver el wordcout ejecutado
  ![resultado vim](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/4238fd73-b416-49d1-b233-6c6190ed7c6c)

- Adentro de HUE podemos revisar que la carpeta `tmp/wcout1` existe y tiene los mismos archivos vistos por consola
  ![resultados en hue hdfs](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/c868c884-c33a-42d5-8dbf-1a54a1997562)

## Wordcount interactivo por S3

- Igual que en el punto anterior (Y sabiendo que los datasets ya estan en S3) abrimos pyspark y comenzamos ejecucion
  ![codigo y salida s3 ](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/1350c541-b684-43b6-b38f-4701eef3218b)

- El codigo a pesar de ser igual internamente cambia con el input y output

  - Sacamos los datos de S3 a traves de `S3://reto5_notebooks/.../*.txt`
  - Los guardamos en el mismo formato de `S3://reto5_notebooks/.../tmp/wcout1`

- Podemos revisar que la salida es correcta entrando al bucket y verificando que la carpeta se creo y por dentro tiene los archivos de resultado
  ![resultado s3 aws](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/9c9d2fc7-8065-451a-8cca-ccb8fb8253d4)

- Tambien verificamos en el S3 de HUE
  ![resultados s3 hue](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/0d326b2f-6a69-4dc7-9480-a98f47d6df11)

## Wordcount por jupyter por S3

- Como hemos vistos en labs anteriores, con el cluster de EMR ya corriendo podemos entrar a JupyterHub y crear un nuevo notebook con el kernel de `PySpark`. Una vez adentro podemos iniciar el kernel con `spark` y `sc`
  ![JUPYTER RESULTADO PUNTO 3](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/7f1feefa-92dc-4840-b540-80ff7b98f157)

- A partir de esto ya podemos leer el input con la funcion `textFile` y pasando la ruta desde S3 (En este caso `S3://reto5_notebooks/.../gutenberg-small/*.txt`)

- Realizamos el mismo codigo de ambos puntos anteriores sin imprimir

- Realizamos un output a S3 pero a una carpeta llamada `jupyter`
  ![resultados s3 aws](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/1513df7c-fd20-491e-84dc-a91cdf13b84e)

## Explicacion del notebook `Data_processing_using_PySpark`

![resultados comentar ipynb](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/8977ed6b-9455-48af-9598-e978d5474a99)

- Este codigo nos permite entender como podemos manipular datos (datasets, etc) a traves de pyspark en el kernel de jupyer.
- Es manipulacion habitual de datos como se hace con scipy, pandas o numpy pero todo a traves de pyspark
- Encontramos maneras de agregar, quitar y derviar columnas. Tambien de encontrar medidas de centralizacion, etc.
- **EN EL JUPYTER EN ESTE MISMO DIRECTORIO SE ENCUENTRA LA EXPLICACION LINEA POR LINEA**
- Podemos encontrar el resutlado del notebook creado en persistencia en S3
  ![resultado aws data processsing](https://github.com/jdramirezl/jdramirezl-st0263/assets/65835577/965e5816-8bab-4060-a3f0-945d0496579d)

# Referencias

[Enunciado del proyecto](https://github.com/st0263eafit/st0263-231/blob/main/bigdata/lab5-3-mrjob-spark.txt)
