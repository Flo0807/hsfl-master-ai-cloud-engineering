DO $$
DECLARE
    word text[] := ARRAY['Lorem', 'ipsum', 'dolor', 'sit', 'amet', 'consetetur', 'sadipscing', 'elitr', 'sed', 'diam', 'nonumy', 'eirmod', 'tempor', 'ut', 'labore', 'voluptua', 'magna', 'aliquyam', 'erat', 'dolore'];
    i integer;
    content text;
BEGIN
    FOR i IN 1..100 LOOP
        content := array_to_string((ARRAY(SELECT unnest(word) LIMIT (5 + floor(random() * 14)::integer))), ' ');
        EXECUTE format('INSERT INTO posts (created_at, updated_at, title, content) VALUES (current_timestamp, current_timestamp, ''Post %s'', ''%s'')', i, content);
    END LOOP;
END $$;