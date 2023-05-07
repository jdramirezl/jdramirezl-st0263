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

Para crear un EMR nos dirigimos a la pestana de `Amazon EMR` en AWS. Una vez aqui presionamos el boton `Create Cluster`. Aqui tendremos que dar las siguientes opciones

1. Nombrar el cluster
2. Elegimos la version del EMR (6.3.1 en mi caso) y application bundle, que sera `Custom` pues nosotros queremos elegir las app especificas
3. Para las instancias de EC2 elegimos `m4.xlarge` para usar con la cuenta de education. Ademas definimos un spot price con el modo `On-Demand`. Lo hacemos para el primary, core y task.
4. Para la terminacion del cluster definimos `1 hora`
5. Definimos el bucket de S3 a usar en `Software Settings`. En las referencias se tiene el sitio desde el cual se copio la configuracion. Yo nombre mi bucket como `reto5-notebooks`
    * Tambien es necesario crear el Bucket en S3, no solo nombrarlo en el config. Se debe llamar igual.
6. Por ultimo definimos una key para poder hacer SSH al cluster, en mi caso use el predeterminado `Vockey`.

## 

1. Nos conectamos al cluster, especificamente a la instancia `EC2` primaria a traves de SSH. Hay opciones para Windows, Linux y Mac. En mi caso yo me conecte a traves de WSL.
    * Recordar entrar al `Security Group` de la instancia y permitir el trafico de SSH.
2. En las pestana `Applications` del cluster revisamos los puertos de las aplicaciones `Hue`, `JupyterHub` y `Zeppelin`.
3. Estos les tendremos que permitir el trafico en la pestana `Block Public Access` de la seccion izquierda. Aqui adentro podremos agregar las IP
4. Igualmente estos puertos los tendremos que poner en el Security group de la instancia privada. Para acceder a esta ingresamos a la informacion del cluster, bajamos a `Network and security` y damos click al SG del `Primary Node`.
5. Una vez adentro agregamos una regla por puerto, poniendo este y un acceso publico con `0.0.0.0/0`


# Referencias
[EMR Software config](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-jupyterhub-s3.html)
