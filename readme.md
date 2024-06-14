# Estudo de Goroutines em Go

Este repositório contém um exemplo simples que explora o uso de goroutines, uma das principais características de concorrência em Go.

## Objetivo

O objetivo deste código é demonstrar como as goroutines podem ser utilizadas para criar um modelo de processamento concorrente e como canais são usados para a comunicação entre essas goroutines.

## Funcionamento

1. **Goroutines como Trabalhadores**
   - O programa cria 10 goroutines, cada uma representando um "trabalhador" (`worker`).
   - Cada trabalhador aguarda por valores de um canal.

2. **Envio de Valores para o Canal**
   - Após a criação das goroutines, o programa envia 100 valores para o canal.

3. **Processamento e Saída**
   - Cada goroutine (trabalhador) recebe um valor do canal, imprime o valor recebido junto com seu identificador e aguarda 1 segundo antes de processar o próximo valor.
   - Isso simula um processamento fictício, como uma operação que leva algum tempo para ser concluída.

## Observações

- **Concorrência em Go**: Goroutines permitem que múltiplas funções sejam executadas simultaneamente, tornando Go eficiente para tarefas concorrentes e paralelas.
- **Canais**: São usados para transmitir dados entre goroutines de maneira segura, facilitando a sincronização e a comunicação.

## Como Executar

Para executar este código, certifique-se de ter o ambiente Go configurado corretamente:

1. Clone este repositório: 
   ```bash
   git clone https://github.com/BredexBR/Goroutines.git

2. Navegue até o diretório do projeto: 
   ```bash
   cd Goroutines

3. Execute o código Go:
   ```bash
   go run main.go