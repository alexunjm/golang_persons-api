# API de Personas

## Descripción

Api construida en golang basado en las clean architectures
El framework web de este bloque es [gin](https://gin-gonic.com/).

## Los principales patrones y estilos de arquitectura que guían este bloque son

### Arquitectura hexagonal

Arquitectura que fomenta que nuestro dominio sea el núcleo de todas las capas, también conocida como puertos y adaptadores en la cual el dominio define los puertos y en las capas superiores se definen los adaptadores para desacoplar el dominio. Se divide principalmente en tres capas, **aplicación**, **dominio** e **infraestructura**

- **Infraestructura**: Capa que tiene las responsabilidades de realizar los adaptadores a los puertos definidos en el domino, exponer web services, consumir web services, realizar conexiones a bases de datos, ejecutar sentencias DML, en general todo lo que sea implementaciones de cualquier framework
- **Aplicación**: capa encargada de enrutar los eventos entrantes de la capa de infraestructura hacía la capa del dominio, generalmente se conoce como una barrera transaccional la cual agrupa toda la invocación de un caso de uso, se pueden encontrar patrones como Fabricas, Manejadores de Comandos, Bus de eventos, etc
- **Dominio**: representa toda la lógica de negocio de la aplicación la cual es la razón de existir del negocio. Se busca evitar el anti-patron [https://martinfowler.com/bliki/AnemicDomainModel.html](https://martinfowler.com/bliki/AnemicDomainModel.html) y favorecer el principio [https://martinfowler.com/bliki/TellDontAsk.html](https://martinfowler.com/bliki/TellDontAsk.html) en esta capa se pueden encontrar los siguientes patrones agregados, servicios de dominio, entidades, objetos de valor, repositorios (puerto), etc.

Para obtener mas documentación sobre este tipo de arquitectura se recomienda [https://codely.tv/blog/screencasts/arquitectura-hexagonal-ddd/](https://codely.tv/blog/screencasts/arquitectura-hexagonal-ddd/)

### Patrón CQRS:

Patrón con el cual dividimos nuestro modelo de objetos en dos, un modelo para consulta y un modelo para comando (modificación de datos). Este patrón es recomendado cuando se va desarrollar lógica de negocio compleja porque nos ayuda a separar las responsabilidades y a mantener un modelo de negocio consistente.

- **Consulta**: modelo a través del cual se divide la responsabilidad para presentar datos en la interfaz de usuario, los objetos se modelan basado en lo que se va a presentar y no en la lógica de negocio, ejm: ver facturas, consultar clientes
- **Comando**: son todas las operaciones que cambian el estado del sistema, ejm: (facturar, aplicar descuento), este modelo se construye todo el modelo de objetos basado en la lógica de negocio de la aplicación

Para mayor documentación del patrón [https://martinfowler.com/bliki/CQRS.html](https://martinfowler.com/bliki/CQRS.html)

### Especificaciones técnicas:

- Framework usado Gin Web Framework
- El código se encuentra agrupado por el patrón [agregado](https://martinfowler.com/bliki/DDD_Aggregate.html); separado por carpeta module
- Uso de sql para el acceso a datos, en los archivos **.env** y **src/infrastructure/config/storage/mysqlstorage/mysql_db_connection.go** se encuentran los datos de conexión a la bd

## Arrancar la aplicación

```bash
# desarrollo
$ docker-compose -f dev.docker-compose.yml up

# modo producción
$ docker-compose up
```
