-- Mostre o título e o nome do gênero de todas as séries.

SELECT title as "título", name as "gênero" FROM series INNER JOIN genres ON series.genre_id = genres.id;

-- Mostre o título dos episódios, o nome e sobrenome dos atores que trabalham em cada um deles.

-- Option A
SELECT title as "título do episódio", CONCAT(first_name, ' ', last_name) as "nome completo" FROM episodes INNER JOIN actor_episode ON episodes.id = actor_episode.episode_id INNER JOIN actors ON actor_episode.actor_id = actors.id ORDER BY title ASC;

-- Option B
SELECT title as "título do episódio", GROUP_CONCAT(CONCAT(first_name, ' ', last_name) SEPARATOR ', ') as "nome completo" FROM episodes INNER JOIN actor_episode ON episodes.id = actor_episode.episode_id INNER JOIN actors ON actor_episode.actor_id = actors.id GROUP BY title ORDER BY title ASC;

-- Mostre o título de todas as séries e o número total de temporadas que cada uma delas possui.

-- Option A
SELECT series.title as "título", COUNT(*) as "número de temporadas" FROM series INNER JOIN seasons ON series.id = seasons.serie_id GROUP BY series.title ORDER BY series.title ASC;

-- Option B
SELECT series.title as "título", (SELECT COUNT(*) FROM seasons WHERE seasons.serie_id = series.id) as "número de temporadas" FROM series ORDER BY series.title ASC;

-- Mostre o nome de todos os gêneros e o número total de filmes de cada um, desde que seja maior ou igual a 3.

-- Option A
SELECT genres.name as "gênero", COUNT(*) as "número de filmes" FROM genres INNER JOIN movies ON genres.id = movies.genre_id GROUP BY genres.name HAVING COUNT(*) >= 3 ORDER BY genres.name ASC;

-- Mostre apenas o nome e sobrenome dos atores que atuam em todos os filmes de Star Wars e que estes não se repitam.

-- Since there are no Star Wars movies in the database, I will use Harry Potter movies instead.

-- Option A
SELECT CONCAT(first_name, ' ', last_name) as "nome completo" FROM actors INNER JOIN actor_movie ON actors.id = SELECT CONCAT(first_name, ' ', last_name) as "nome completo" FROM actors INNER JOIN actor_movie ON actors.id = actor_movie.actor_id INNER JOIN movies ON actor_movie.movie_id = movies.id WHERE movies.title LIKE '%Harry Potter%' GROUP BY CONCAT(first_name, ' ', last_name) HAVING COUNT(*) = (SELECT COUNT(*) FROM movies WHERE movies.title LIKE '%Harry Potter%') ORDER BY CONCAT(first_name, ' ', last_name) ASC;

-- Option B
SELECT CONCAT(a.first_name, ' ', a.last_name) as "nome completo" FROM actors a INNER JOIN actor_movie am ON a.id = am.actor_id INNER JOIN movies m ON am.movie_id = m.id WHERE m.title LIKE '%Harry Potter%' GROUP BY CONCAT(a.first_name, ' ', a.last_name) HAVING COUNT(*) = (SELECT COUNT(*) FROM movies WHERE title LIKE '%Harry Potter%') ORDER BY CONCAT(a.first_name, ' ', a.last_name) ASC;