--1
SHOW TABLES;
--2
SELECT title FROM film WHERE length > 120;
--3
SELECT title FROM film WHERE rating='PG-13' ORDER BY length LIMIT 4; 
--4
SELECT DISTINCT film.title, language.name AS 'language' FROM film INNER JOIN language ON film.language_id=language.language_id WHERE film.description LIKE '%drama%';
--5
SELECT DISTINCT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Family'AND film.description LIKE '%Documentary%';
--6
SELECT DISTINCT film.title FROM film JOIN film_category ON film.film_id = film_category.film_id JOIN category ON film_category.category_id = category.category_id WHERE category.name = 'Children' AND film.rating <>'PG-13';
--7
SELECT rating, COUNT(*) FROM film GROUP BY rating ORDER BY rating;
--8
SELECT DISTINCT film.title FROM rental JOIN inventory ON rental.inventory_id = inventory.inventory_id JOIN film ON inventory.film_id = film.film_id WHERE rental.rental_date BETWEEN '2005-05-31 00:00:00' AND '2005-06-30 23:59:59' ORDER BY film.title DESC;
--9
SELECT DISTINCT actor.first_name, actor.last_name FROM actor JOIN film_actor ON actor.actor_id = film_actor.actor_id JOIN film ON film_actor.film_id = film.film_id WHERE film.special_features LIKE '%Deleted%';
--10
SELECT customer.first_name, customer.last_name FROM customer JOIN rental ON customer.customer_id = rental.customer_id JOIN payment payment1 ON rental.rental_id = payment1.rental_id JOIN staff staff1 ON staff1.staff_id = payment1.staff_id JOIN payment payment2 ON rental.rental_id = payment2.rental_id JOIN staff staff2 ON payment2.staff_id = staff2.staff_id WHERE staff1.staff_id <> staff2.staff_id;
--11

--12
SELECT DISTINCT 
    CONCAT(a1.first_name, ' ', a1.last_name) AS actor1,
    CONCAT(a2.first_name, ' ', a2.last_name) AS actor2
FROM 
    film_actor fa1
JOIN 
    film_actor fa2 ON fa1.film_id = fa2.film_id AND fa1.actor_id < fa2.actor_id
JOIN 
    actor a1 ON fa1.actor_id = a1.actor_id
JOIN 
    actor a2 ON fa2.actor_id = a2.actor_id
WHERE 
    (a1.actor_id, a2.actor_id) IN (
        SELECT 
            fa.actor_id AS actor1_id, 
            fb.actor_id AS actor2_id
        FROM 
            film_actor fa
        JOIN 
            film_actor fb ON fa.film_id = fb.film_id AND fa.actor_id < fb.actor_id
        GROUP BY 
            actor1_id, actor2_id
        HAVING 
            COUNT(DISTINCT fa.film_id) > 1
    );
--13

--14

--15

--16

--17

--18

--19
