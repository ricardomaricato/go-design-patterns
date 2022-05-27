## Builder

Tudo bem, vamos tentar resumir o que aprendemos sobre o padrão de projeto do construtor, então um construtor é um 
componente separado que é usado para construir um objeto e para tornar um construtor fluente, que é uma das coisas que
às vezes é útil. Você pode devolver o receptor. Então você tem um método e esse método pode realmente retornar o receptor.
Então retorne o ponteiro, que permite encadear para que você possa ter várias chamadas de métodos do construtor uma atrás 
da outra sem nenhum tipo de cerimônia, claro, vai formatador, meio que quebra um pouco as coisas porque tem que ter 
cuidado com as quebras de linha. Você quer fazer quebras de linha, mas precisa ter cuidado com elas. Assim, diferentes 
facetas de um objeto também podem ser construídas com vários construtores diferentes. Então, se você tem um objeto muito,
muito complicado e se você tem diferentes aspectos desse objeto sendo construídos como preocupações separadas, como 
aspectos separados que você pode querer separar, então você pode ter vários construtores que podem trabalhar em conjunto 
por meio de algum tipo de construção comum que eles estão agregando e podem trabalhar em conjunto para construir os 
diferentes aspectos desse objeto.

