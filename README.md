
# orionxctools 

Este paquete fue creado para facilitar las consultas a la api de orionx.

Paso a paso
============

* Primero debes obtener un KEY en [API](https://orionx.io/developers/keys). Hacer click en el botón CREATE KEY.
* Dar los permisos correspondientes a la Key.
* Agregar la Secret Key y la Api Key en el código de ejemplo.

#### Ejemplo

La forma simple de usar orionxtools es:

```go
  package main

  // Importar paquete orionxctools
  import "github.com/dhinojosac/orionxclient/orionxctools"

  const (
    API_KEY = "put your API_KEY here"
    KEY     = "put your KEY here"
  )

  func main() {

    queries := []string{
      `{market(code:"CHACLP"){
        lastTrade{
          price
        }
        }
      }`,
      `{cha:marketOrderBook (marketCode:"CHACLP") {
          spread
          mid
        }}`,
    }

    for _, query := range queries {
      orionxctools.MakeRequest(KEY, API_KEY, query)
    }
  }  
```

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

