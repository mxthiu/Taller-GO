# Taller Go: Variables y punteros

## Integrantes

- Mathias Calvo Fernandez, carné 2025067626
- [Nombre completo 2], [carné]

## Lenguaje

Go

## Reparto del trabajo

- Mathias: ejercicios 1, 3, 5 y 6.
- [Nombre 2]: ejercicios 2, 4, 7 y 8.

## Cómo ejecutar

1. Instalar Go desde https://go.dev/dl/
2. Abrir una terminal en la carpeta del taller.
3. Ejecutar el comando `go run main.go`

## Resultados de las pruebas

Ejercicio 1, intercambio

- Caso a = 8, b = -3. Esperado: a = -3, b = 8. Obtenido: a = -3, b = 8. Correcto.

Ejercicio 3, copia independiente

- Al modificar la copia, el original [4, -2, 0, 7, -5] se conserva. Correcto.

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

## Evitar ternas repetidas en 3Sum

Lo primero que hacemos es ordenar una copia del arreglo. Así los valores iguales quedan juntos y es más fácil ver cuándo se repiten.

Después fijamos el primer número (s[i]) y con dos punteros buscamos los otros dos. Si el primer número es igual al anterior, nos lo saltamos, porque ya habríamos sacado las mismas ternas con el anterior.

Cuando encontramos una terna, movemos left y right, y luego dejamos avanzar left mientras el valor sea igual al que acaba de pasar, y lo mismo con right. Así no se repite la misma combinación.

## Por qué el agua depende del menor de los dos máximos

[Lo completa el compañero que hace el ejercicio 8.]

## Alcance de variables

En main.go, la variable tmp dentro de swap solo existe mientras la función corre. Cuando swap termina, tmp desaparece, así que no hay problema si otra función también usa un nombre tmp.

Otro ejemplo está en threeSum. Las variables left y right se declaran dentro del for de i, así que cada vuelta empieza con sus propios punteros nuevos.

## Manejo de memoria, Go frente a C

En C uno pide la memoria con malloc, tiene que revisar que no sea NULL y después liberarla con free. Si se olvida el free, queda una fuga de memoria.

En Go se usa make para crear los slices y no hay que liberar nada a mano. El recolector de basura borra lo que ya no se usa. La ventaja es que uno se olvida menos de liberar memoria y evita muchos errores. La desventaja es que no controla exactamente cuándo se libera la memoria, y el recolector gasta algunos recursos.

En este taller, copyArray usa make para crear la copia, y en ningún momento se libera manualmente.

## Observaciones


