CREATE TABLE e_2020_sample(
    shirt int8,
    cardigan int8,
    chiffon int8,
    pants int8,
    heels int8,
    socks int8
)

INSERT INTO e_2020_sample (shirt, cardigan, chiffon, pants, heels, socks)
VALUES (5, 20, 36, 10, 10, 20)
RETURNING *
