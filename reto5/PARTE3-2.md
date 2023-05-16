# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Reto 5 Parte 3 - 2

# Wordcount en Apache Spark EN AWS EMR

# Descripcion de la actividad

Realizar ejercicios de MAP-REDUCE en EMR con SPARK

## Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

- ejecutar el wordcount por linea de comando 'pyspark' INTERACTIVO en EMR con datos en HDFS vía ssh en el nodo master.

- ejecutar el wordcount por linea de comando 'pyspark' INTERACTIVO en EMR con datos en S3 (tanto de entrada como de salida)  vía ssh en el nodo master.

- ejecutar el wordcount en JupyterHub Notebooks EMR con datos en S3 (tanto datos de entrada como de salida) usando un clúster EMR.

- Replique, ejecute y EXPLIQUE el notebook: Data_processing_using_PySpark.ipynb con los datos respectivos en S3
  ejecutelo en AWS EMR notebooks jupyterhub o en AWS EMR zeppelin.

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

- Cada paso que ejecutamos nos da una serie de steps que representan el proceso de carga del codigo enviado

## Wordcount interactivo por S3

## Wordcount por jupyter por S3

## Explicacion del notebook `Data_processing_using_PySpark`

### Ejecucion del notebook en `jupyterhub` por S3

# Conclusiones

# Referencias
