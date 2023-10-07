SHOW TABLES;
SELECT title FROM film WHERE length > 120;
SELECT title FROM film WHERE rating='PG-13' ORDER BY length LIMIT 4; 
SELECT DISTINCT film.title, language.name AS 'language' FROM film INNER JOIN language ON film.language_id=language.language_id WHERE film.description LIKE '%drama%';
SELECT DISTINCT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Family'AND film.description LIKE '%Documentary%';
SELECT DISTINCT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Children' AND film.rating <>'PG-13';
SELECT rating, COUNT(*) FROM film GROUP BY rating ORDER BY rating;
SELECT DISTINCT film.title FROM rental JOIN inventory ON rental.inventory_id = inventory.inventory_id JOIN film ON inventory.film_id = film.film_id WHERE rental.rental_date BETWEEN '2005-05-31 00:00:00' AND '2005-06-30 23:59:59' ORDER BY film.title DESC;
SELECT DISTINCT actor.first_name, actor.last_name FROM actor JOIN film_actor ON actor.actor_id = film_actor.actor_id JOIN film ON film_actor.film_id = film.film_id WHERE film.special_features LIKE '%Deleted%';
