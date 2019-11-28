# KNN Distribuido - Tarea2
---

Se busca cumplir con el objetivo de aplivar el algoritmo de KNN de manera distribuida atraves de una red P2P en diferentes puertos para fragmentar el procesamiento.

Se utiliza el dataset de Iris.

## Como compilar:

- Entrar a `serverknngo.go` en la linea 86 editar el data set entre `iris.csv` o `iristest.csv`.

- `go run serverknngo.go`

- Definir caso de prueba:
    - Primer puerto => 8060.
    - Segundo puerto => 9060.
    - Remota 1 => 8060.
    - Remota 2 => 9060.
    - Insertar el valor de k, en cualquier puerto, **valor de ejemplo 3**.
    - KNN magic done, el output debe mostrar los vecinos mas cercanos para los test de prueba para los elementos escogidos de manera aleatoria.


## Conclusiones

Para este trabajo la mayor complejidad existio en hacer distribuido la aplicacion del algoritmo de KNN. Dentro de lo propuesto se busco repartir las tareas de la aplicacion de la euclidiana en los diferentes puertos.