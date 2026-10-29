package main

// Ejercicio 1: intercambio de valores con punteros
func swap(a *int, b *int) {
	tmp := *a // tmp vive solo dentro de la función
	*a = *b
	*b = tmp
}
