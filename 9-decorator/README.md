## Decorator

Tudo bem, vamos tentar resumir o que aprendemos sobre o padrão de projeto do decorador, então um decorador incorpora o 
objeto decorado e então adiciona quaisquer campos ou métodos utilitários para aumentar os recursos dos objetos. 
E o decorador é frequentemente usado para emular algo que chamaríamos herança múltipla em outras linguagens no sentido 
de que basicamente adquire os comportamentos e os campos não apenas de um único objeto, mas de vários objetos. 
É isso que a incorporação faz por nós. Mas pode haver problemas como, por exemplo, se você incorporar dois objetos que
têm o mesmo campo, por exemplo, um campo com o mesmo nome, então acaba tendo que gerir a consistência entre esses 
objetos diferentes porque esses dois campos serão únicos. Isso não será mesclado num único campo. E então eu mostrei a
APAGA-lo problemas que surgem e como pode resolvê-los.
