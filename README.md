
# orionxctools 

Este paquete fue creado para facilitar las consultas a la api de orionx.

Paso a paso
============

* Primero debes obtener un KEY en [API](https://orionx.com/developers/keys). Hacer click en el botón CREATE KEY.
* Dar los permisos correspondientes a la Key.
* Agregar la Secret Key y la Api Key en el código de ejemplo.
* Obtener la librería usando el comando ```go get github.com/dhinojosac/orionxctools```

#### Ejemplo
Agregar las credenciales como variables de entorno **ORIONX_KEY** y **ORIONX_API_KEY**.
Correr el ejemplo *main.go* que se encuentra en la carpeta test.


Si las request fueron exitosas, la salida es:
```
Status: 200 OK
Body: {"data":{"market":{"lastTrade":{"price":760}}}}
Status: 200 OK
Body: {"data":{"cha":{"spread":17,"mid":771}}}
```

Autor
=====

* Diego Hinojosa Cordova [dhinojosac](http://github.com/dhinojosac)

