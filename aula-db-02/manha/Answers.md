# Bancos de dados relacionais

## Aula 02 - Manhã

### Primeira Parte

1. Como é chamado um JOIN em um banco de dados

R: Join

2. Nomeie e explique 2 tipos de JOIN

R: 

Inner Join: Retorna a interseção das tabelas
Full Outer Join: Retorna tudo

3. Para que serve o GROUP BY?

R: Agrupar informações. Exemplo: Agrupar atores de um mesmo filme

4. Para que serve o HAVING?

R: Filtrar informações. Exemplo: Após agrupar os atores de um mesmo filme, selecionar apenas os que tiverem >2 awards.

5. Indique o tipo de Join

a. Interseção Tables A e B (descrição da figura)

R: Inner Join

b. Table A, incluindo interseção A com B (descrição da figura)

R: Left Join

6. Escreva uma consulta genérica 

a. Table B, incluindo interseção A com B (descrição da figura)

R: Right Join

SELECT * FROM table_A RIGHT JOIN table_B on table_A.something = table_B.something

b. Table A e Table B, incluindo a interseção A com B (descrição da figura)

R: Full Outer Join;

SELECT * FROM table_A FULL OUTER JOIN table_B on table_A.something = table_B.something;


### Exercício 2

No arquivo queries.sql
