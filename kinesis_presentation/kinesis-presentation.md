<!-- $theme: default -->

# Kinesis, lambda, golang e dynamodb
## André Leoni
Enjoei.com.br


---

# Objetivo

Adquirir conhecimento básico de estrutura serverless...

A idéia é abrir a visão para novas soluções de projetos :D

---


# O que é stream?

Um fluxo de dados de processamento.

É comum ser utilizado por aplicações de stream de video, como Youtube, mas pode ser utilizado para processamento de dados também.

Microsoft utiliza stream no pacote office para evitar a perda de documentos. No qual as alterações são salvas em um stream, e caso seja executado a recuperação, é efetuado as alterações salvas em ordem até chegar ao estado atual.

![](/Users/andreleoni/Documents/kinesis%20presentation/o%20que%20e%20stream%20.png)

---

# Kinesis

Serviço de stream da AWS, para criação e processamento de stream.

![](/Users/andreleoni/Documents/kinesis%20presentation/kinesis%20home.png)


---

# Lambda

Serviço da AWS que possibilita funções serverless que executam qualquer coisa...

Elas tendem a ser bem rápidas

---

# GOLang

Linguagem do Pedro

Recentemente tem surgido vários posts explicando o porque GOlang tem sido a melhor opção para funções serverless, como o lambda

É muito bom para trabalhar com paralelismo e concorrência

---

# Dynamodb

Banco de dados não relacional da amazon

É focado em performance, garante até 10ms de tempo de resposta para qualquer requisição

Você não escala memória/processamento, e sim quantidade de leitura/escrita por segundo

O identificador único de um objeto é a junção do em partition_key e sort_key, sendo que o sort_key é também utlizado para ordenação

---

# CloudWatch

Serviço para criação de alertas para os logs das aplicações


---

# Mão na massa

1 - Criaremos um kinesis
`test_enjoei_kinesis`
2 - Criaremos um Dynamodb
`test_enjoei_dynamo`
3 - Criaremos uma função lambda
`test_enjoei_lambda`
4 - Compilaremos e enviaremos a função compilada para o dynamo
	`GOOS=linux go build -o lab lab.go`
5 - Vincularemos a função ao kinesis
6 - Enviaremos dados para o kinesis
`ruby create_batch_kinesis.rb`
7 - Faremos uma consulta para visualizar o registro criado
`ruby get_items_dynamo.rb`
8 - Conferiremos o log no cloudwatch