-- SQLite
SELECT 
    p.id_product,
    t.name,
    p.value_per_meter,
    p.total_value,
    p.thickness,
    p.cor
FROM 
    product p
JOIN 
    product_type t on p.id_type = t.id_type



select * from product
