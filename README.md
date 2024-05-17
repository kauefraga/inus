# Inus

> Um sistema confiável de login/logout. Veja o cliente em [inuc](https://github.com/kauefraga/inuc).

Contextualizando, inus e [inuc](https://github.com/kauefraga/inuc) são as duas partes de um só projeto, o inu. A palavra "*inu*" vem do japonês e traduz para "cachorro", um animal confiável e fenomenal.

É isso que você pode esperar desse projeto: uma **interface fenomenal** e, por trás dessa interface, um **sistema confiável**.

- inu + **c**lient :: inuc
- inu + **s**erver :: inus

Resumindo o funcionamento do sistema é o seguinte: uma rota para **criação** de uma conta, outra para efetuar o **login** de uma conta e, por último, uma rota para **excluir** uma conta. Para mais detalhes, confira a seção [entendendo o sistema](#entendendo-o-sistema).

## Entendendo o sistema

###### Criação de uma conta

O servidor recebe uma requisição POST na rota `/users` contendo os dados a seguir no formato JSON.

```json
{
  "name": "nomedeusuario",
  "email": "nome.de.usuario@example.com",
  "password": "usu4ri0d3n0m3"
}
```

Ao receber a requisição, o servidor verifica se o campo "name" não está vazio e não tem menos do que 4 nem mais do que 50 caracteres. Também verifica se o campo "email" não está vazio.

Após a validação dos dados, é verificado através de uma consulta ao banco de dados se o usuário com tal nome já existe.

Depois das verificações mencionadas, é gerado um *hash* da senha usando o algoritmo **bcrypt**, que é recomendado pela segurança extra do mesmo. Concluído esse passo, o usuário é inserido no banco de dados.

Por último, é criado um **token JWT** que contém algumas informações (nome de usuário, está autorizado ou não e tempo de expiração do token) e será enviado como resposta em um *cookie HTTPOnly*.

###### Login de uma conta

Um pouco mais simples do que a criação de uma conta, o servidor também recebe uma requisição POST contendo um JSON (`/login`).

```json
{
  "name": "nomedeusuario",
  "password": "usu4ri0d3n0m3"
}
```

Depois da mesma validação do nome é feita uma consulta para pegar o *hash* da senha do usuário, guardado no banco de dados. Nesse momento também é possível saber se o usuário existe ou não.

Com o hash em mãos, a senha da requisição é comparada ao hash da mesma. Caso a senha corresponda, um token JWT é criado com os mesmos parâmetros da [criação de uma conta](#criação-de-uma-conta) e enviado como resposta.

###### Exclusão de uma conta

O servidor recebe uma requisição DELETE na rota `/users` e pega o nome de usuário que está no payload do token JWT. Caso não tenha um token, o usuário não está autenticado e, por conseguinte, não pode excluir tal conta.

Com o nome do usuário, um `DELETE` COM `WHERE` é executado. Apenas o código de status 204 (*no content*) é retornado no sucesso dessa operação de exclusão.
