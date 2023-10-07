SHOW TABLES;
SELECT title FROM film WHERE length > 120;
SELECT title FROM film WHERE rating='PG-13' ORDER BY length LIMIT 4; 
SELECT film.title, language.name AS 'language' FROM film INNER JOIN language ON film.language_id=language.language_id WHERE film.description LIKE '%drama%';
SELECT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Family'AND film.description LIKE '%Documentary%';
SELECT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Children' AND film.rating <>'PG-13';