-- Explique o conceito de normalização e por que ele é usado.

-- A normalização é uma técnica usada para eliminar redundâncias e inconsistências em um banco de dados.
-- Exemplo:

-- Tabela de filmes:
-- id | nome | ano | premios
-- 1 | filme1 | 2023 | 5
-- 2 | filme2 | 2023 | 3

-- Tabela de atores:
-- id | nome | filme_favorito (nome do filme)
-- 1 | ator1 | filme1
-- 2 | ator2 | filme2

-- Aqui temos uma redundância, pois o nome do filme está sendo repetido na tabela de atores. O ideal seria algo tipo:

-- Tabela de filmes:
-- id | nome | ano | premios
-- 1 | filme1 | 2023 | 5
-- 2 | filme2 | 2023 | 3

-- Tabela de atores:
-- id | nome | filme_favorito (id do filme)
-- 1 | ator1 | 1
-- 2 | ator2 | 2


-- Adicione um filme à tabela de filmes.

INSERT INTO movies (created_at, updated_at, title, rating, release_date, length, awards, genre_id) VALUES (NOW(), NOW(), 'Filme 1', 8.5, '2024-05-06', 180, 10, 1);
-- movie id: 22

-- Adicione um gênero à tabela de gêneros.

INSERT INTO genres (created_at, updated_at, name, ranking, active) VALUES (NOW(), NOW(), 'Gênero 1', 100, 1);
-- genre id: 14

-- Associe o filme do Ex 2. ao gênero criado no Ex 3.

UPDATE movies SET genre_id = 14 WHERE id = 22;

-- Modifique a tabela de atores para que pelo menos um ator tenha o filme adicionado no Ex.2 como favorito.

-- Using actor id 1, 2 and 3
UPDATE actors SET favorite_movie_id = 22 WHERE id IN (1, 2, 3);

-- Crie uma cópia de tabela temporária da tabela de filmes.

CREATE TEMPORARY TABLE temp_movies AS SELECT * FROM movies;

-- Elimine dessa tabela temporária todos os filmes que ganharam menos de 5 prêmios.

DELETE FROM temp_movies WHERE awards < 5;

-- Obtenha a lista de todos os gêneros que possuem pelo menos um filme.

SELECT g.name, COUNT(m.id) AS number_of_movies FROM genres g JOIN movies m ON g.id = m.genre_id GROUP BY g.name HAVING COUNT(m.id) > 0;

-- Obtenha a lista de atores cujo filme favorito ganhou mais de 3 prêmios.

-- Join actors and movies, filter by awards > 3, use concat for the actor name (first_name and last_name)
SELECT CONCAT(a.first_name, ' ', a.last_name) AS actor_name, m.title AS favorite_movie, m.awards AS number_of_awards FROM actors a JOIN movies m ON a.favorite_movie_id = m.id WHERE m.awards > 3;

-- Use o plano de explicação para analisar as consultas em Ex.6 e 7.

EXPLAIN SELECT g.name, COUNT(m.id) AS number_of_movies FROM genres g JOIN movies m ON g.id = m.genre_id GROUP BY g.name HAVING COUNT(m.id) > 0;

EXPLAIN SELECT CONCAT(a.first_name, ' ', a.last_name) AS actor_name, m.title AS favorite_movie, m.awards AS number_of_awards FROM actors a JOIN movies m ON a.favorite_movie_id = m.id WHERE m.awards > 3;

-- O que são índices? Para que servem?

-- Índices são estruturas de dados (possivelmente uma árvore binária) que permitem a busca rápida de informações em um banco de dados.
-- Eles servem para melhorar o desempenho das consultas, pois permitem que o banco de dados encontre os dados mais rapidamente, sem precisar percorrer toda a tabela. (pergunta: a busca cai para O(log n)? ou O(1)? ou depende do tipo de índice?)

-- Crie um índice no nome na tabela de filmes.

CREATE INDEX idx_movies_title ON movies (title);

-- Verifique se o índice foi criado corretamente.

SHOW INDEX FROM movies;