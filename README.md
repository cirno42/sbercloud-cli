# sbercloud-cli

##  Установка

### Шаг 0:
Перед установкой необходимо установить Go версии 1.16 или выше.


### Шаг 1:

В директории с файлом main.go выполнить:

Для Windows:
> go build -o scli.exe

Для Linux:
> go build -o scli

### Шаг 3:

Сгенерировать Access Key и Security Key.

Подробнее о том, как создать Access Key и Security Key: https://support.hc.sbercloud.ru/en-us/usermanual/iam/iam_02_0003.html

В разделе My Credentials получить Project ID.

### Шаг 4:

Сконфигурировать приложение при помощи команды:

>./scli configure

