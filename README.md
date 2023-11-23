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

4a) GET http://<url>/avisaagua/leitura?sensor1=ligado
    - retorna 400 bad request (porque faltaram as informacoes dos sensores 2 e 3, so tem do 1)
    
4b) GET http://<url>/avisaagua/leitura? sensor1=ligado &sensor2=ligado &sensor3=desligado
    - insere uma nova linha com data/hora em medicoes.txt e com os valores dos sensores
    - retorna 200 OK + "Leitura OK"

5) Calcular se deve ou nao ligar o motor de drenagem da agua baseado nos valores dos sensores 3  de nivel de agua-
  - cada sensor pode estar "ligado" ou desligado".
  - sensor ligado significa que ele esta debaixo dagua
  - sao 3 sensores colocados em alturas diferente
  - quando a caixa de maquinas esta sem agua os 3 senores estao desligados
  - o sensor 1 é o mais fundo. se ele estiver ligado é porque ja tem agua ali
  - o sensor 2 esta um pouco acima
  - o sensor 3 é o mais alto. se ele estiver ligado é porque a  caixa de maquinas ja esta critica, bastante alagada. 

  - calcular se quer ou nao ligar o motor para drenar a agua baseado nos valores dos 3 sensores
  - retornar "LIGAR MOTOR" ou "DESLIGAR MOTOR" para o comando GET http://<url>/avisaagua/leitura?sensor1=<valor>&sensor2=<valor>&sensor3=<valor>

