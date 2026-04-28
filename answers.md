##### Ejercicio 1 (Req/Rep)

b) i) El cliente hace el envio para luego bloquearse esperando la respuesta.

b) ii) El servidor recibe el mensaje, lo procesa y envia la respuesta. El cliente recibe la respuesta y se desbloquea.

b) iii) No, ya que no se recibe un mensaje de conexion rechazada, simplemente se queda esperando la respuesta.

b) iv) Por como esta implementado el protocolo, el cliente intenta conectarse el cual zeromq se encarga de reintentar la conexion hasta que el servidor responda.

c) 


##### Ejercicio 2 (Pub/Sub)

1) a) Si detengo solamente el subscriber, el publisher sigue mandando mensajes sin parar. 

   b) Los mensajes publicados mientras el subscriber esta detenido simplemente se pierden si es que no hay ningún otro subscriber escuchando en ese tópico

   c) Si volvemos a ejecutar el subscriber va a recibir los mensajes que manda el publisher en el mismo momento en el que corren ambos, todos aquellos mensajes enviados cuando el subscriber se encontraba detenido no serán recuperados por el subscriber. 

   d) Si detengo solo el publisher, el subscriber deja de recibir mensajes (dado que no hay ningún mensaje que recibir por parte del publisher) pero sigue esperando mensajes.

   e) Si volvemos a ejecutar el publisher luego de detenerlo, como el subscriber se habia quedado esperando algún mensaje simplemente comenzará a recibir los nuevos mensajes que envia el publisher. Un detalle es que la conexión ya queda preestablecida de la vez anterior, por lo que no vemos un mensaje indicando la conexión. 

   f) Lo único que nos llama la atención es lo comentado anteriormente acerca del mensaje de la conexión y que el mensaje numero 1 que envia el publisher no es mencionado por el subscriber. 

2)

   a) Si, funciona. Ambos subscribers muestran los mensajes recibidos.

   b) No no funciona. Dado que ya tenemos un publisher publicando en un puerto, cuando intentamos levantar el segundo obtenemos el error de que esa dirección ya esta ocupada. Esto es debido a como funcionan las funciones bind y connect 

   c) Si cambiamos el topico ya sea en cliente o servidor, el subscriber no será capaz de recibir los mensajes dado que estos se estan enviando por un canal "diferente" que por el que esta escuchando. 

   d) Si, un pub puede tener más de un tópico en ejecución

   e) Si, un subscriber puede atender muchos tópicos. 



##### Ejercicio 3 (Push/Pull)

1) a) La diferencia que encuentro es que en el patron Push/Pull tanto producer como worker se esperan mutuamente para transmitir mensajes, mientras que en pub/sub el publisher manda mensajes sin importar si hay alguien escuchando o no. 

    b) La primera diferencia que veo es que cuando tengo dos workers en vez de ambos recibir todos los mensajes como en pub/sub, se van "dividiendo" esos mensajes, por lo que un worker procesa los impares y el otro los pares (en el caso de tener 2 workers, de haber más se dividirian de igual forma entre ellos). En cuanto a ejecutar dos producers al mismo tiempo, se obtiene el mismo error de que el puerto se encuentra ocupado.
