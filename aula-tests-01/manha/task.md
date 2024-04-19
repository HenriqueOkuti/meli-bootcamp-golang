# Questions

## Exercício 1

*Quais são as diferenças entre a caixa branca e a caixa preta?*

R: Essencialmente, a diferença está relacionada com o que está sendo testado em sí:

- Testes de caixa preta são testes feitos com base apenas na entrada e saída, sem se preocupar com a implementação do código em si, verificando apenas se os parâmetros de entrada e saída estão dentro do previsto 
- Testes de caixa branca são focados em testar a implementação do código, visando detectar possíveis vulnerabilidades (contexto específico)

## Exercício 2

*O que é um teste funcional?*

R: Um teste funcional é aquele que testa o funcionamento dos requisitos da aplicação. Exemplo: Uma API possui uma rota de autenticação, essa rota de fato autentica algo?

## Exercício 3

*O que é um teste de integração?*

R: Testes de integração focam em testar a comunicação entre as camadas, pegando o exemplo da questão anterior:

- Ao solicar uma autenticação, a camada responsável por receber a solicitação consegue se comunicar com a camada responsável por gerar o token de autenticação? (exemplo simplificado)

## Exercício 4

*Indique as dimensões prioritárias de qualidade no MELI*

R:

- Funcionalidade
- Eficiência
- Confiança
- Segurança
- Manutenibilidade
