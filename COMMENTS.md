Este documento é um diário, aqui eu detalho cada dia e decisões que fui tomando conforme o desenvolvimento.

## Dia 1

Resolvi começar desenhando e criando o esqueleto do projeto. O objetivo era montar um projeto com o básico para iniciar um desenvolvimento produtivo. Fiz uma lista com as coisas que eu queria garantir.

![image](https://user-images.githubusercontent.com/3903012/95229585-1fee3500-07d7-11eb-9a80-d25ab907e25c.png)


### Web Framework

Era neste momento que eu iria escolher o framework, eu estava entre: Gin, Echo e Negroni. Uma das coisas que eu queria era uma facilidade para documentar a API. Para isto, eu achei o swaggo/swag. Dos três frameworks que tinha cogitado, somente o Negroni não possuia suporte. Fiquei entre o Gin e Echo, acabei escolhendo o Echo simplesmente por no momento ter mais familiaridade.

### A estrutura de pastas

Pensei em utilizar clean architeture. Mas repensei e achei melhor começar simples, se achar necessário eu refatoro no futuro. A decisão foi começar usando três camada.  A estrutura de pastas inicial ficou assim:

- `app`
    - `api`: Arquivos responsáveis pela Rest API, configuração de rotas
    - `common`: Arquivos de úteis, por exemplo logs e errors
    - `healthcheck`: Package por "modulo", dentro dele terá as três arquivos, 3 camadas (handler, service e repository).
        - handler
        - service
        - repository
- `docker`: Dockerfiles
- `docs`:  Arquivos do swagger

## Dia 2

Comecei o desenvolvimento pela accounts. Criei um arquivo para implementar um banco `in memory`, usando map mesmo. Assim, somente no final eu me preocupo com o banco.

Mudei um pouco a estrutura inicial. Eu criei um package `modules` e movi o `healthcheck` para ele. Assim as pastas dos modulos não ficam soltas. Estou planejando isolar cada modulo num package e assim mitigar as interdependências. Penso que essa técnica torna mais fácil quebrar esse app em microservices no futuro (se necessário).

Agora estrutura ficou assim:

- `app`
    - `modules`
        - `healthcheck`
        - `accounts`

Passeo a usar o conceito de `representation`.  Dessa maneira consigo ter mais controle sobre qual camada entrega o que. E também consigo ter flexibilidade usando o validator ([https://github.com/go-playground/validator](https://github.com/go-playground/validator)) somente nas structs das representations

Entrada:

`representation`→ `handler`→ `transforma em model`→ `service` → `repository`

Saída:

`repository` → `model` → `service` → `handler` → `transforma em representation`

## Dia 3

Desenvolvi o Login, resolvi criar um module separado só para ele. Sobre parte de retornar o token, sem dúvidas vou usar o JWT, visto que ele serve justamente para isso.

Durante o desenvolvimento fui tendo ideias de melhorias, validações novas. Para não interromper o desenvolvimento vou só anotar e se der tempo faço no final.

Algumas das ideias:

- Validar CPF
    - Não pode ter duas accounts com o mesmo CPF
    - Validar o numero do CPF
- No endpoint que retorna todas as accounts, retornar paginado.

## Dia 4 e 5

### Dilema 1: Maybe Repeat Yourself?

Comecei a desenvolver o Core da aplicação, a parte de transferências. Até o momento ainda não estou usando um banco de dados de verdade, estou usando o Repository pattern pra isso.

Cheguei num momento que estou na dúvida. Vejo que terei que pesar e perder de um lado para ganhar de outro.

A situação é a seguinte, eu organizei o projeto por modulos, está ficando assim:

```
.
├── app
    ├── modules
        ├── account
        ├── healthcheck
        ├── login
        └── transfer
```

A ideia é que cada um dos modulos tenha zero interdependencias. De modo que consiga separar eles em serviços no futuro. Então, pensando assim, poderiamos ter um serviço responsavel só pelo cadastro (accounting), outro só para o login e outro só para transferências.

Para fazer isso, estou implementando as três camadas dentro de cada pacote (handler, service e repository). O que acontece é que no repository da `transfer`  terei que implementar um método que já existe no repository da `account` . Seria o método `GetAccount`.  Eu iria ferir então o conceito DRY. Mas ao mesmo tempo, não vejo ganho em mudar a estrutura do projeto por causa disso. Penso que manter separado eu ganho na diminuição de dependencia e aumento da estabilidade ([https://skarlso.github.io/2019/04/21/efferent-and-afferent-metrics-in-go/](https://skarlso.github.io/2019/04/21/efferent-and-afferent-metrics-in-go/)).

**Conclusão:** Vou seguir assim por enquanto, se eu ver que vou precisar duplicar muito código, eu mudo a estratégia.

### Dilema 2: Transaction na camada de negócio?

Cheguei em outra questão de organização, analisando a regra de negócio percebi que preciso ter um controle de transações, senão uma das contas pode acabar perdendo dinheiro.

O dilema que eu estava era: Faz sentido a camada `service` conhecer o "conceito" de transação?  Ao mesmo tempo, eu não queria fazer um método na camada `repository` que tivesse toda a regra de negócio da transferência.

Li bastante sobre, aparentemente esse é mais um momento que você tem que escolher entre ferir a camadas em troca de segurança. Vi que existe alguns patterns (por exemplo SAGA pattern) que podem mitigar isso. A ideia do SAGA pattern é implementar o "Ctrl+Z" no caso de algo der errado. Porém eu pensei que seria overengineering implementar esse pattern agora na primeira versão.

**Conclusão:** Eu resolvi expor os métodos `BeginTransaction`, `Rollback` e `Commit` no repository. A ideia é, mesmo o service tendo que conhecer o conceito, ainda não ser necessário mudanças no Service independente do banco escolhido.

### Data race

Vi a vantagem de usar o Repository pattern. Por causa dele eu consegui identificar um problema de data race que está ocorrendo no código.

Eu identifiquei assim: Primeiro fiz um teste de carga no método de transferência.

Foi algo bem simples, criei uma conta com 100000 reais e abri 2 terminais para começar a fazer CURL transferindo para outra.

No final do processo eu chequei ambas contas. A origem estava com zero reais, correto. A destino, ora ficava com 100100, ora 90900. Após isso eu resolvi rodar com a flag de `race` ligada.  Percebi que estava ocorrendo problemas de dirty read, apliquei lock `sync.mutex` no método de transferência. O data race parou de ocorrer.

## Dia 6

### Banco de dados

Agora vou começar a implementação colocando um banco de verdade. Devido a natureza da solução vou escolher um banco de dados relacional com suporte a transações.

Vou testar o GORM (nunca utilizei), achei interessante a proposta dele e acho que vai ajudar a dar uma acelerada.

Achei interessante que ele é compatível com vários bancos de dados. Outra coisa que é interessante desse ORM é que ele já vem com um Migration embarcado.

 Vou começar desenvolvendo os repositories dele e fazer os primeiros testes com um SQLite mesmo.

Depois eu refatoro para conectar num MySQL e incremento o Docker Compose para subir um MySQL junto.

## Dia 7

Hoje é o ultimo dia, vou usar para revisar a aplicação e deixar ela o mais `production ready` o possível.

Eu tenho uma lista de melhorias que fui anotando conforme ia tendo as ideias. Vou pegar esse último dia para atacar essa lista.

### Evitando ponteiros desnecessários

O primeiro refactor que eu ataquei foi de revisar todos os métodos e parar de retornar ponteiros quando possivel. Isso vai ajudar o código a ficar mais estável, evitando mutações surpresa, nil dereference errors não tratados e etc (me inspirei nesse post: [https://medium.com/better-programming/why-you-should-avoid-pointers-in-go-36724365a2a7](https://medium.com/better-programming/why-you-should-avoid-pointers-in-go-36724365a2a7)).

Como a cobertura de teste está razoável, me senti bem tranquilo de refatorar.

### Erro muito descritivo não é seguro

A segunda mudança é uma correção de segurança, eu tinha desenvolvido de modo que, quando o usuário erra o CPF e senha, o erro estava descritivo demais, avisando qual campo estava errado. No trabalho atual aprendi que isso é uma má prática, visto que dá pistas para os má intencionados explorar.

### Validações na criação da conta

Implementei duas validações novas: CPF é unico na base e não é possivel cadastrar um CPF inválido.

Para deixar agnostico de banco de dados, resolvi deixar a regra de unicidade na camada de serviço, logo não utilizar somente o indice `unique` já presente nos bancos. Porém eu configurei o campo como `unique` só para garantir que não será criado duplicado por outro meio (double check).

### Paginação

Atualmente existem dois endpoints que podem retornar um volume algo: `GET accounts` e `GET Transfers`. A primeira versão deles implementei sem paginação. Porém entendo que isso não é production ready. Aproveitei o tempo que sobrou para implementar esse ajuste.

## O que mais dá pra fazer?

Ainda há mais algumas coisas que eu gostaria fazer caso tivesse mais tempo. Vou listar abaixo:

- **Testes end-to-end:** Eu acredito que os testes end-to-end são muito importantes pois iriam abranger partes que os testes unitários não cobrem, como por exemplo a integração correta com o banco de dados.
- **Testes de carga:** Gostaria de aplicar um teste de carga para entender as limitações da aplicação, por exemplo, qual o limite de operações. Dependendo desse limite, gostaria de refatorar, aplicando técnicas como caching, com o intuito de melhorar esse tempo de resposta.
- **Metricas:** A métricas talvez viriam antes dos testes de carga, pois acredito que uma coisa complementa a outra. Acompanhar e diagnosticar o teste de carga fica muito mais tranquilo com métricas. Neste caso eu iria utilizar Prometheus + Grafana.
- **CI:** gostaria de implementar um pipeline CI, que pelo menos rode os testes e o coverage (podendo então colocar o badges mostrando os perceituais)