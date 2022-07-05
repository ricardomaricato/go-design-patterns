## Facade

Vamos tentar resumir o que aprendemos sobre o padrão de projeto de fachada, para construir uma fachada. Bem, a ideia de
construir uma fachada é fornecer alguma API simplificada sobre um conjunto de componentes. E esses componentes podem ser
bastante complicados. Eles podem ter muitos detalhes e peças, basicamente apenas um sistema muito complicado. E também
podemos opcionalmente expor esses internos através da fachada para que a fachada possa tentar escondê-los o máximo que
puder. Mas às vezes podemos querer expor esses detalhes internos para que, se tivermos um utilizador avançado, alguém 
que realmente queira entender o que acontece, eles também possam manipular esses detalhes de implementação. Então, às
vezes teria APIs simples e compreensíveis, mas também permitiria que os utilizadores escalassem o uso de APIs mais 
complexas. Então, por exemplo, teria funções que aceitam parâmetros adicionais onde pode especificar as opções avançadas,
digamos, e cabe ao cliente usá-las ou não, ou simplesmente deixar tudo e o nível compreensível ou ir em profundidade e 
personalizar o sistema para o conteúdo do seu coração.
