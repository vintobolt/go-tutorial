# Заметки с примерами на языке Go

## Используемые пакеты
    * go get golang.org/x/text/encoding/charmap
## Структура пакетов
1. iodata - Работа с вводом выводом.
    * keyboard - С помощью scanf
    * scanner - echo с помощью bufio.NewScanner  
    * reader - Запись в слайс os.Stdin 
    * stdoutoerr - Вывод в канал ошибок.
    * openfile - Чтение и запись в файл.
    * readfile - Считываение при помощи ридера и запись в переменную.
    * charset - Запись и чтение в разных кодировках.
    * readBinary - Чтение и запись двоичных данных.