# GoDFA
GoDFA es un programa escrito en go para convertir un automata finito no determinista(NFA) a un automata finito determinista(DFA) equivalente.

## Entrada
El programa recibe un archivo de texto plano con la descripcion del NFA, la sintaxis del archivo es la siguiente.
* Todas las lineas en blanco y lineas que comiencen con el simbolo '#' seran ignoradas.
1. La primera linea debera contener los estados del automata separados por ','.
2. La segunda linea debera contener los simbolos en el alfabeto del automata separados por ','.
3. La tercera linea debera de contener el estado inicia del automata y este debera de ser alguno de los estados contenidos en la primera linea.
4. La cuarta linea debera de contener los estados finales del automata separados por ','. Todos estos estados deben de aparecer tambien en la primera linea del archivo.
5. Para las lineas siguientes cada una de estas representara una fila de la tabla de transiciónes del automata y tendran el siguiente formato.
    * El estado del automata seguido de una coma. seguido de una lista separada por ',' indicando a donde va la transición con cada uno de los simbolos del alfabeto en elorden en el que se encuentran en la segunda linea.
    * Si la transición de un estado con un simbolo va a mas de un estado estos deben de estar separados por espacios en blanco.
    * Si la transición de un estado con un simbolo es vacia esta se puede indicar con el simbolo '_' o simplemente con espacios en blanco.

## Compilacion 
El programa fue escrito usando la version 13.3 de [Go](https://golang.org/), para compilar el programa se requiere un compilador de este lenguaje compatible con esta version y ejecutar el siguiente comando dentro de la raiz del proyecto, esto producira un ejecutable llamado ``godfa``.
```` shell
$go build
````

## Ejecucion
Para ejecutar el programa se requiere indicar el archivo de entrada y salida para el programa usando los atributos ``-i`` y ``-o``, por ejemplo.
````shell
$godfa -i example.nfa -o salida.dfa
````

## Salida
La salida del programa sera un archivo de texto plano con la descripcion del DFA equivalente al NFA de entrada siguiendo la misma sintaxis que el archivo de entrada.

## Archivos
El proyecto esta separado en diferentes archivos y cada uno tiene la siguiente funcion.
* ``bitset.go`` contiene la implementacion de la estrutura de datos bitset.
* ``dfa.go`` contiene la descripcion de la estrutura dfa , el codigo para convertir el nfa de entrada y el codigo para generar la salida del programa.
* ``example.nfa`` es un archivo un ejemplo de archivo de entrada para el programa.
* ``hashset`` contiene la implementacion de la estructura hashset.
* ``main.go`` contiene funcion main del programa y lo necesario para leer los atributos de la ejecucion del programa
* ``nfa.go`` contiene la implementacion de la estructura nfa, el codigo para leer el archivo de entrada y crear una representacion del nfa.  