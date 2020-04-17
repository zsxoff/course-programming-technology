# Лабораторная работа 3. Формат JSON

Требуется разработать приложение или программный комплекс, обменивающийся данными по сети в формате JSON. Приложение должно позволять всем пользователям, физически находящимся за разными вычислительными устройствами, обмениваться текстовыми сообщениями по сети.

> Вариант 14

Сетевая игра «Звёздное ремесло» для двух игроков. Изначально у каждого игрока есть 0 кристаллов, 5 рабочих и 0 воинов. Игроки ходят одновременно, они не видят количество кристаллов, рабочих и воинов другого игрока. В свой ход игрок может либо потратить некоторое количество имеющихся кристаллов на строительство рабочих или воинов, либо напасть на другого игрока. Производство рабочего стоит 5 кристаллов, а воина – 10 кристаллов. В начале хода игрок получает столько кристаллов, сколько у него рабочих. Если один из игроков решил напасть на другого игрока, то он выигрывает, если у него строго больше воинов, чем у другого игрока, и проигрывает в противном случае.

## Запуск

### Имитация нескольких вычислительных устройств с помощью Docker

Создание подсети для контейнеров выполняется командой:

```bash
docker network create lab-net
```

Сборка Docker-контейнера выполняется с помощью команды:

```bash
docker build -t "lab-container-3" .
```

Запуск контейнеров для клиента и сервера выполняется с помощью запуска скриптов в контейнере:

> Порт по умолчанию: 7777

```bash
docker run --rm -it --net=lab-net lab-container-3 ./laboratory-3 -mode=server -port=7777
docker run --rm -it --net=lab-net lab-container-3 ./laboratory-3 -mode=client -port=7777 -ip=SERVERADDR
```

### Сборка и запуск программы на одном вычислительном устройстве

Сборка программы выполняется с помощью команды:

```bash
go build
```

Запуск сервера и клиента выполняется командами:

```bash
./laboratory-3 -mode=server -port=7777
./laboratory-3 -mode=client -port=7777 -ip=SERVERADDR
```