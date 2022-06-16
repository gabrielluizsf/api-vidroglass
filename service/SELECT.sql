SELECT 
i.quantidade, i.valor, i.desconto,
i.metragem_produto, p.espessura, p.cor, tp.nome
FROM item i 
join produto p on p.id_produto = i.id_produto
join tipo_produto tp on tp.id_tipo_produto = p.id_tipo
where i.id_nota = 1