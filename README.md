# Taller Go: Variables y punteros

## Integrantes

- Mathias Calvo Fernandez, carné 2025067626
- Felipe Bejarano Molina, 2024801138

## Lenguaje

Go

## Reparto del trabajo

- Mathias: ejercicios 1, 3, 5 y 6.
- Felipe: ejercicios 2, 4, 7 y 8.

## Cómo ejecutar

1. Instalar Go desde https://go.dev/dl/
2. Abrir una terminal en la carpeta del taller.
3. Ejecutar el comando `go run .`

## Resultados de las pruebas

Ejercicio 1, intercambio

- Caso a = 8, b = -3. Esperado: a = -3, b = 8. Obtenido: a = -3, b = 8. Correcto.

Ejercicio 2, Estadisticas de un arreglo

- [4, -2, 0, 7, -5]. Esperado: suma = 4. Obtenido: suma = 4. Correcto.
- [4, -2, 0, 7, -5]. Esperado: mínimo = -5. Obtenido: mínimo = -5. Correcto.
- [4, -2, 0, 7, -5]. Esperado: máximo = 7. Obtenido: máximo = 7. Correcto.
- [4, -2, 0, 7, -5]. Esperado: negativos = 2. Obtenido: negativos = 2. Correcto.
- [4, -2, 0, 7, -5]. Esperado: recorrido inverso = -5 7 0 -2 4. Obtenido: -5 7 0 -2 4. 

Explicación:

- La función `suma` recorre el arreglo acumulando todos sus elementos.
- La función `minimo` busca el menor valor presente en el arreglo.
- La función `maximo` busca el mayor valor presente en el arreglo.
- La función `contarNegativos` cuenta cuántos elementos son menores que cero.
- La función `recorridoInverso` imprime los elementos desde la última posición hasta la primera sin modificar el arreglo original.

Ejercicio 3, copia independiente

- Al modificar la copia, el original [4, -2, 0, 7, -5] se conserva. Correcto.

Ejercicio 4,

- Entrada: [4, -2, 0, 7, -5].
- Esperado: [-5, -2, 0, 4, 7].
- Obtenido: [-5, -2, 0, 4, 7].

Ejercicio 5, búsqueda de pares

- [-4, -1, 0, 2, 5] con objetivo 1. Esperado: un par válido es (-4, 5). Obtenido: un par válido es (-4, 5). Correcto.
- [-4, -1, 0, 2, 5] con objetivo 8. Esperado: no existe. Obtenido: no existe. Correcto.
- [0, 0] con objetivo 0. Esperado: (0, 0). Obtenido: (0, 0). Correcto.
- [3] con objetivo 6. Esperado: no existe. Obtenido: no existe. Correcto.

Ejercicio 6, 3Sum

- [-1, 0, 1, 2, -1, -4]. Esperado: 2 ternas. Obtenido: 2 ternas. Correcto.
- [0, 1, 1]. Esperado: 0 ternas. Obtenido: 0 ternas. Correcto.
- [0, 0, 0]. Esperado: 1 terna (0, 0, 0). Obtenido: 1 terna (0, 0, 0). Correcto.
- [0, 0, 0, 0]. Esperado: 1 terna (0, 0, 0). Obtenido: 1 terna (0, 0, 0). Correcto.
- [-2, 0, 0, 2, 2]. Esperado: 1 terna (-2, 0, 2). Obtenido: 1 terna (-2, 0, 2). Correcto.

Ejercicio 7, Maximos desde ambos extremos

Entrada:

height = [3, 0, 2, 0, 4]

Esperado:

leftMax = [3, 3, 3, 3, 4]
rightMax = [4, 4, 4, 4, 4]

Obtenido:

leftMax = [3, 3, 3, 3, 4]
rightMax = [4, 4, 4, 4, 4]



Ejercicio 8, Trappin Water

| Entrada | Resultado esperado | Explicación |
|----------|----------|-------------|
| [5] | 0 | Solo existe una columna. No hay paredes a ambos lados que permitan retener agua. |
| [0, 1, 2, 3] | 0 | Las alturas aumentan continuamente. El agua siempre puede escapar hacia la izquierda, por lo que no se acumula. |
| [3, 2, 1, 0] | 0 | Las alturas disminuyen continuamente. El agua siempre puede escapar hacia la derecha, por lo que no se acumula. |
| [2, 2, 2] | 0 | Todas las columnas tienen la misma altura. No existen depresiones donde pueda almacenarse agua. |
| [3, 0, 3] | 3 | Las columnas de altura 3 en los extremos forman una cavidad. La columna central puede almacenar 3 unidades de agua. |
| [3, 0, 2, 0, 4] | 7 | Existen dos depresiones entre las paredes. La posición 1 almacena 3 unidades, la posición 2 almacena 1 unidad y la posición 3 almacena 3 unidades, para un total de 7. |
| [4, 2, 0, 3, 2, 5] | 9 | Las columnas altas de los extremos permiten retener agua en varias posiciones intermedias. La suma del agua almacenada en cada posición es 9 unidades. |
| [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1] | 6 | Caso clásico del problema. Varias depresiones pequeñas almacenan agua en distintos puntos del arreglo. El volumen total retenido es de 6 unidades. |


- Se añadió el caso de prueba [2, 0, 2], las columnas de los extremos tienen altura 2 y la columna central tiene altura 0. El nivel del agua queda limitado por la menor pared de ambos lados, que en este caso es 2. Por tanto, la posición central almacena 2 unidades de agua y el volumen total atrapado es 2

## Evitar ternas repetidas en 3Sum

Lo primero que hacemos es ordenar una copia del arreglo. Así los valores iguales quedan juntos y es más fácil ver cuándo se repiten.

Después fijamos el primer número (s[i]) y con dos punteros buscamos los otros dos. Si el primer número es igual al anterior, nos lo saltamos, porque ya habríamos sacado las mismas ternas con el anterior.

Cuando encontramos una terna, movemos left y right, y luego dejamos avanzar left mientras el valor sea igual al que acaba de pasar, y lo mismo con right. Así no se repite la misma combinación.

## Por qué el agua depende del menor de los dos máximos

En cada una de las posiciones el agua queda limitada por la pared más baja entre el maximo a la izquierda y el maximo a la derecha. Por eso se utiliza min(leftMax[i], rightMax[i]). Si uno de los lados es mas bajo, el agua se iria por ese lado y no podria alcanzar una altura mayor.

## Alcance de variables

En ejercicio1.go, la variable tmp dentro de swap solo existe mientras la función corre. Cuando swap termina, tmp desaparece, así que no hay problema si otra función también usa un nombre tmp.

Otro ejemplo está en threeSum. Las variables left y right se declaran dentro del for de i, así que cada vuelta empieza con sus propios punteros nuevos.

## Manejo de memoria, Go frente a C

En C uno pide la memoria con malloc, tiene que revisar que no sea NULL y después liberarla con free. Si se olvida el free, queda una fuga de memoria.

En Go se usa make para crear los slices y no hay que liberar nada a mano. El recolector de basura borra lo que ya no se usa. La ventaja es que uno se olvida menos de liberar memoria y evita muchos errores. La desventaja es que no controla exactamente cuándo se libera la memoria, y el recolector gasta algunos recursos.

En este taller, copyArray usa make para crear la copia, y en ningún momento se libera manualmente.
