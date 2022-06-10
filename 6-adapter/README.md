## Adapter

OK, então vamos tentar resumir o que aprendemos sobre o padrão de projeto do adaptador, então implementar um adaptador
é bem fácil. Basicamente, determina a API que tem e determina a API que precisa para que a coisa toda realmente funcione.
E então cria algum componente que agrega ou tem um ponteiro para o adaptively. E então tem a coisa toda fornecendo os 
dados apropriados. E às vezes, sim, às vezes verá uma situação em que tem representações intermediárias e elas podem se 
acumular. E, neste caso, precisa usar o cache e outras otimizações para garantir que a quantidade de dados temporários
que gera ao fornecer o adaptador seja gerenciável e não ultrapasse os limites.