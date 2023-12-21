# AvisaAgua
monitora nivel e bombeia agua da casa de maquinas das piscina

# Requisitos
1) GET http://<url>/avisaagua/hello
    - retorna 200 OK + "Hello AvisaAgua <data><hora>"
    
2) GET http://<url>/avisaagua/teste
    - insere uma nova linha com data/hora em medicoes.txt
    - retorna 200 OK + "Teste AvisaAgua <data><hora>"

3) GET http://<url>/avisaagua/leitura
    - retorna 400 bad request (porque faltaram as informacoes dos 3 sensores)

4) Aceitar url com "query string" e informações dos sensores
    - https://blog.petehouston.com/parse-query-string-in-gin-web-application/
    - imprimir no console e incluir no arquivo medicoes.txt o valor do query string "sensor1" 

5) Acrescentar no medicoes.txt o valor dos sensores 2 e 3 alem do 1.
    - escrever "sensor2=True" ou "sensor2=False" em vez de ligado/desligado, user bool do golang

6) ...
 
7) GET http://<url>/avisaagua/leitura?sensor1=ligado
    - retorna 400 bad request (porque faltaram as informacoes dos sensores 2 e 3, so tem do 1)
    
8) GET http://<url>/avisaagua/leitura? sensor1=ligado &sensor2=ligado &sensor3=desligado
    - insere uma nova linha com data/hora em medicoes.txt e com os valores dos sensores
    - retorna 200 OK + "Leitura OK"

9) Calcular se deve ou nao ligar o motor de drenagem da agua baseado nos valores dos sensores 3  de nivel de agua-
  - cada sensor pode estar "ligado" ou desligado".
  - sensor ligado significa que ele esta debaixo dagua
  - sao 3 sensores colocados em alturas diferente
  - quando a caixa de maquinas esta sem agua os 3 senores estao desligados
  - o sensor 1 é o mais fundo. se ele estiver ligado é porque ja tem agua ali
  - o sensor 2 esta um pouco acima
  - o sensor 3 é o mais alto. se ele estiver ligado é porque a  caixa de maquinas ja esta critica, bastante alagada. 

  - calcular se quer ou nao ligar o motor para drenar a agua baseado nos valores dos 3 sensores
  - retornar "LIGAR MOTOR" ou "DESLIGAR MOTOR" para o comando GET http://<url>/avisaagua/leitura?sensor1=<valor>&sensor2=<valor>&sensor3=<valor>

10) Publicar a URL no linux
- usar "ngrok http 8080"
- testar com browser em outro computador

11) mover o codigo fonte par ao github
- mover o main.go para dentro do diretorio "webserver"
- o codigo do arduino vai ficar em "arduino"
  
12) acessar o github com chaves SSH
- gerar a chave publica e privada
- fazer upload da chave publica para o github.com
- ver se git clone funciona
- fazer uma mudanca minima em main.go, entao no linux a) git status b) git add main.go c) git commit "descricao do commit" d) git push, ver no browser se atualizou

13) vamos preparar o main.go para crescer.
- vamos ter duas funcionalidades 1) receber as leituras e 2) uma pagina web para monitorar as ultimas leituras
- em (1) no futuro a gente pode hum ... trocar de arquivo medicoes.txt para um database, e suportar mais de uma casa no mesmo webserver
- em (2) podemos ter uma pagina para mostrar a ultima medicao, as da ultima hora, ultimo dia etc. ultima vez que alagou. os ultimos 5 alagamentos, etc. tipo relatorios.
- então vamos dividir o fonte em pacotes golang, um pacote para (1) e um pacote para (2)

14) Escolher o nome dos 2 packages
- https://go.dev/blog/package-names
- escolher os dois nomes
- 
15) 
16)  
17) 
18) 
