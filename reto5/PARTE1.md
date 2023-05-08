# Info de la materia: ST0263 Topicos Especiales en Telematica

# Estudiante: Julian David Ramirez Lopera, jdramirezl@eafit.edu.co

# Profesor: Edwin Nelson Montoya, emontoya@eafit.brightspace.com

# Reto 5 Parte 1
# Montaje de EMR

# Descripcion de la actividad
Realizar el montaje de un servicio de EMR (Elastic Map Reduce) en AWS para usar servicios de big data.

## Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
- Crear un cluster EMR en AWS

## Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
N/A

# Descripcion del ambiente de ejecucion
- AWS EMR 6.3.1 con
    * Hadoop
    * Jupyterhub
    * Hive
    * Sqoop
    * Zeppelin
    * Tez 
    * JupyterEnterpriseGateway
    * Hue
    * Spark
    * Livy
    * HCatalog

 # IP o nombres de dominio en nube

 # Descripcion de la configuracion del proyecto

## Crear el cluster
Para crear un EMR nos dirigimos a la pestana de `Amazon EMR` en AWS. Una vez aqui presionamos el boton `Create Cluster`. Aqui tendremos que dar las siguientes opciones

1. Nombrar el cluster

![Screenshot 2023-05-07 132504](https://user-images.githubusercontent.com/65835577/236702067-14ba0500-8318-4299-ad05-9da0d3336f3f.png)

2. Elegimos la version del EMR (6.3.1 en mi caso) y application bundle, que sera `Custom` pues nosotros queremos elegir las app especificas

![Screenshot 2023-05-07 132521](https://user-images.githubusercontent.com/65835577/236702087-24c56008-016d-4aff-b0f3-95ef8613fa43.png)

3. Para las instancias de EC2 elegimos `m4.xlarge` para usar con la cuenta de education. Ademas definimos un spot price con el modo `On-Demand`. Lo hacemos para el primary, core y task.

![Screenshot 2023-05-07 132539](https://user-images.githubusercontent.com/65835577/236702091-ca34fe5b-5729-42e6-b084-5c93e3cf3001.png)

4. Para la terminacion del cluster definimos `1 hora`

![Screenshot 2023-05-07 132600](https://user-images.githubusercontent.com/65835577/236702096-76d5c756-af3a-4770-8ec6-3f1fd714344b.png)

5. Definimos el bucket de S3 a usar en `Software Settings`. En las referencias se tiene el sitio desde el cual se copio la configuracion. Yo nombre mi bucket como `reto5-notebooks`
  * Tambien es necesario crear el Bucket en S3, no solo nombrarlo en el config. Se debe llamar igual.

![Screenshot 2023-05-07 132612](https://user-images.githubusercontent.com/65835577/236702099-db15438f-4c30-463a-b46c-01b6b3c37d53.png)

6. Por ultimo definimos una key para poder hacer SSH al cluster, en mi caso use uno propio `reto5-keypair`.

![image](https://user-images.githubusercontent.com/65835577/236702124-9b746dc2-c6bb-4bf0-946a-d115e4841f9b.png)


## Configuracion y permisos de red

1. Nos conectamos al cluster, especificamente a la instancia `EC2` primaria a traves de SSH. Hay opciones para Windows, Linux y Mac. En mi caso yo me conecte a traves de WSL.
    * Recordar entrar al `Security Group` de la instancia y permitir el trafico de SSH.

![ssh connect](https://user-images.githubusercontent.com/65835577/236702148-744f7bc0-f0a4-4fb0-b6e0-eb98b92e4257.png)
![conectssh](https://user-images.githubusercontent.com/65835577/236702154-96ecadd8-966e-42dc-99ba-d7c3199578b4.png)
![acceso al nodo master](https://user-images.githubusercontent.com/65835577/236702162-9fa53795-3e20-472f-bfa4-36a5fdc368cd.png)

2. En las pestana `Applications` del cluster revisamos los puertos de las aplicaciones `Hue`, `JupyterHub` y `Zeppelin`.

![ips apps](https://user-images.githubusercontent.com/65835577/236702173-a130e3b4-741a-4912-8701-77cc82876301.png)

3. Estos les tendremos que permitir el trafico en la pestana `Block Public Access` de la seccion izquierda. Aqui adentro podremos agregar las IP

![block public access](https://user-images.githubusercontent.com/65835577/236702179-2ad11b90-65f6-4355-86d0-80750ca65cdc.png)

4. Igualmente estos puertos los tendremos que poner en el Security group de la instancia privada. Para acceder a esta ingresamos a la informacion del cluster, bajamos a `Network and security` y damos click al SG del `Primary Node`.

![primary cluster security](https://user-images.githubusercontent.com/65835577/236702210-804f564d-8142-4e33-92be-3475dda987b9.png)

5. Una vez adentro agregamos una regla por puerto, poniendo este y un acceso publico con `0.0.0.0/0`

![new rules secirtury group](https://user-images.githubusercontent.com/65835577/236702195-7c61a9fc-6f61-4e14-abd1-55ce42fca319.png)

## Ingreso y configuracion de las Apps

- Volvemos a la informacion del cluster y accedemos a la pestana `Applications`. Aqui encontraremos las URL para los servicios instalados en el cluster. 
![ips apps](https://user-images.githubusercontent.com/65835577/236713959-ff21edcd-136c-415c-8a07-c570ed9179aa.png)

### Hue
1. Ingresamos a la URL de `Hue`

![hue url](https://user-images.githubusercontent.com/65835577/236713976-006d3280-a6d6-4f72-9ccf-9e639211164c.png)

2. Cuando entramos por primera vez debemos crear una cuenta, que en este ejercicio sera con user `Hadoop`.

![hue cuenta](https://user-images.githubusercontent.com/65835577/236713972-a69c9265-72f5-4c03-8e8a-b052f2d36b30.png)

3. Aqui tenemos muchas opciones disponibles como Spark, Hive, archivos HDFS y S3. A continuacion evidencia de que el bucket creado en la configuracion esta en Hue

![hue buckets](https://user-images.githubusercontent.com/65835577/236713980-2eb549d7-6b45-4f04-8cc3-118724bb655d.png)

### JupyterHub
1. Ingresamos a la URL de `JupyterHub`

![jupyter url](https://user-images.githubusercontent.com/65835577/236713985-0527912f-2aea-4782-8d26-e472f99cc7c2.png)

2. No tenemos que crear cuenta, ingresamos con las credenciales por defecto de user `jovyan` y contrasena `jupyter`

![sign in jupyter](https://user-images.githubusercontent.com/65835577/236714428-f1efe180-6836-4f39-b70b-576ce8ea973b.png)

3. Probamos creando un notebook PySpark y probamos que funcionen las variables de `spark` y `sc` copiandolas en las casillas del notebook.

![ne notebook jupyeter](https://user-images.githubusercontent.com/65835577/236714433-4adac260-31c2-48e0-8932-fab01dbc7130.png)

![spark funcionando jupyter](https://user-images.githubusercontent.com/65835577/236714438-72ec5be5-9d25-43d1-b989-e81ef8176051.png)

### Zeppelin
1. Ingresamos a la URL de `Zeppelin` y entramos sin usuario

![ze ppelin url](https://user-images.githubusercontent.com/65835577/236714001-e5c69488-00da-4c99-a604-fd3dd2951bc0.png)

2. Creamos un nuevo notebook en `Create new note`. Le damos nombre y le elegimos como `Default interpreter` a `Spark`

3. Ya adentro probamos las variables `%spark.pyspark` y `spark` para probar el correcto funcionamiento del interprete


# Resultados
Podemos evidenciar como en el bucket que habiamos creado inicialmente esta el archivo de Jupyter respectivo


# Referencias
[EMR Software config](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-jupyterhub-s3.html)
