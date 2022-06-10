## Bridge

Tudo bem, vamos tentar resumir o que aprendemos sobre o padrão de projeto da ponte, então a ideia aqui é que, para a ponte,
dissociamos a abstração da implementação. Eu sei que soa um pouco científico, tipo o que fazemos aqui? Mas isso é 
realmente o que acontece nos bastidores. Então, separamos algumas abstrações da implementação e criamos uma hierarquia 
paralela para isso. E assim evitamos toda a explosão de complexidade e tanto a abstração quanto a implementação. 
Eles podem existir como hierarquias separadas, o que significa que tem uma hierarquia que também usa outra, em oposição
a essa abordagem de produto vetorial. Pode pensar no padrão de ponte realmente como uma forma mais forte de encapsulamento.