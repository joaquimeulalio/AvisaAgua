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
 
8) GET http://<url>/avisaagua/leitura?sensor1=ligado
    - retorna 400 bad request (porque faltaram as informacoes dos sensores 2 e 3, so tem do 1)
    
9) GET http://<url>/avisaagua/leitura? sensor1=ligado &sensor2=ligado &sensor3=desligado
    - insere uma nova linha com data/hora em medicoes.txt e com os valores dos sensores
    - retorna 200 OK + "Leitura OK"

10) Calcular se deve ou nao ligar o motor de drenagem da agua baseado nos valores dos sensores 3  de nivel de agua-
  - cada sensor pode estar "ligado" ou desligado".
  - sensor ligado significa que ele esta debaixo dagua
  - sao 3 sensores colocados em alturas diferente
  - quando a caixa de maquinas esta sem agua os 3 senores estao desligados
  - o sensor 1 é o mais fundo. se ele estiver ligado é porque ja tem agua ali
  - o sensor 2 esta um pouco acima
  - o sensor 3 é o mais alto. se ele estiver ligado é porque a  caixa de maquinas ja esta critica, bastante alagada. 

  - calcular se quer ou nao ligar o motor para drenar a agua baseado nos valores dos 3 sensores
  - retornar "LIGAR MOTOR" ou "DESLIGAR MOTOR" para o comando GET http://<url>/avisaagua/leitura?sensor1=<valor>&sensor2=<valor>&sensor3=<valor>

11) Publicar a URL
- usar "ngrok http 8080"
- testar com browser em outro computador

12) 
