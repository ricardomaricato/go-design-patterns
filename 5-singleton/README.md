## Singleton

Tudo bem, então vamos falar sobre o que vimos sobre o assustador padrão de design Singleton tão fácil, uma instanciação
do singleton é possível usando sinc uma vez, e isso resolve todos os nossos problemas relacionados à instanciação
preguiçosa, bem como a segurança do thread, a propósito. Portanto, a segurança do thread também é tratada pelos sinc. 
Mas queremos aderir ao princípio de inversão de dependência para que o nosso código permaneça testável e outros enfeites.
E a idéia aqui é que em vez de depender diretamente do Singleton e usar esse Singleton diretamente no seu código, o que
quer fazer é que o seu Singleton implemente alguma interface e então depende da interface. E a razão pela qual é 
importante é porque pode substituir o implementador dessa interface. Pode substituir o Singleton por, digamos, algum
manequim de teste que pode usar nos seus testes sem, digamos, se for um verdadeiro singleton entrar num banco de dados
ativo, não quer isso no seu teste. Deseja fornecer alternativas. E com essa abordagem, com a abordagem de depender de 
abstrações, obtém exatamente isso. Então, a lição, eu acho, desta sessão, esta seção do curso é que o singleton não é 
um padrão assustador, mas tem que ter cuidado. Tem que ter cuidado, porque se você usá-lo diretamente, se você usar o
Singleton e depender diretamente do singleton, então é uma dependência muito forte. É uma dependência que pode vir a 
prejudicá-lo no final. Portanto, o princípio de inversão de dependência é particularmente relevante quando se fala sobre
o singleton.
