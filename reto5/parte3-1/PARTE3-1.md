# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Reto 5 Parte 3 - 1

# MAP/REDUCE

# Descripcion de la actividad

Ejercicios básicos de MapReduce con MRJOB en python

## Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

- Para el dataset: dataempleados.csv

  1. El salario promedio por Sector Económico (SE)
  2. El salario promedio por Empleado
  3. Número de SE por Empleado que ha tenido a lo largo de la estadística

- Para el dataset: datempresas.csv

  1. Por acción, dia-menor-valor, día-mayor-valor
  2. Listado de acciones que siempre han subido o se mantienen estables.
  3. DIA NEGRO: Saque el día en el que la mayor cantidad de acciones tienen el menor valor de acción (DESPLOME), suponga una inflación independiente del tiempo.

- Para el dataset: datapeliculas.csv

  1. Número de películas vista por un usuario, valor promedio de calificación
  2. Día en que más películas se han visto
  3. Día en que menos películas se han visto
  4. Número de usuarios que ven una misma película y el rating promedio
  5. Día en que peor evaluación en promedio han dado los usuarios
  6. Día en que mejor evaluación han dado los usuarios
  7. La mejor y peor película evaluada por genero

## Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

N/A

# Descripcion del ambiente de ejecucion

- WSL - Ubuntu 22.04
- Windows 10
- Python3.10
  - MRJob

# Ejecucion de la actividad

Para ejecutar localmente los ejercicios se hace necesario instalar MRJob, una libreria para ejecutar map-reduce en python. La instalamos con el comando `pip install mrjob`

Despues de eso descargamos los datasets necesarios para el ejercicio. Estos son `datapeliculas`, `dataempresas` y `dataempleados`

A partir de esto comenzamos a resolver los ejercicios.

- Los codigos estan en la carpeta `./codigos`
- Los resultados estan en la carpeta `./outputs`
- Los datasets estan en la carpeta `./datasets`
