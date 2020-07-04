# PlayGround для BTC (тестирование стратегий сборки транзакции)

## btc

Здесь хранятся стандартные настройки нод, всего их 3 (смотри: docker-compose.yml)

Для генерации адреса, надо указывать явный тип адреса (p2sh-segwit)

## generator

Здесь хранится скрипт генерации блоков и биткойнов на miner ноде, по сути какой-то bash скрипт, который внутри себя будет выпускать блоки.

При первом запуске должен:
1. Получить адрес для майнинга: `curl -d '{"jsonrpc":"2.0","id":"1","method":"getnewaddress", "params": ["mining"]}' -u user:password miner:18400`
2. Сгенерить 100 блоков: `curl -d '{"jsonrpc":"2.0","id":"1","method":"generatetoaddress", "params":[101, "_ADDRESS_"]}' -u user:password localhost:18400`
3. Запустить `watch -n _SECONDS_ 'curl -d \'{"jsonrpc":"2.0","id":"1","method":"generatetoaddress", "params":[1, "_ADDRESS_"]}\' -u user:password localhost:18400'`

При повторных запусках только
1. Запустить `watch -n _SECONDS_ 'curl -d \'{"jsonrpc":"2.0","id":"1","method":"generatetoaddress", "params":[1, "_ADDRESS_"]}\' -u user:password localhost:18400'`

Проверку можно делать так: `curl -d '{"jsonrpc":"2.0","id":"1","method":"getaddressesbylabel", "params": ["mining"]}' -u user:password miner:18400`

## tester
