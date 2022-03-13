# Заметки с примерами на языке Go

## Используемые пакеты
    * go get golang.org/x/text/encoding/charmap
## Структура пакетов
### 1. gocore - база языка.
    * array - Массивы.
    * booleantype - Булевый тип.
    * conditionalgo - Условные операторы.
### 2. iodata - Работа с вводом выводом.
    * keyboard - Ввод с клавиатуры строки и Scanf.
    * scanner - Echo с помощью bufio.NewScanner.  
    * reader - Запись в слайс os.Stdin. 
    * stdoutoerr - Вывод в канал ошибок.
    * openfile - Чтение и запись в файл.
    * readfile - Считываение при помощи ридера и запись в переменную.
    * charset - Запись и чтение в разных кодировках.
    * readBinary - Чтение и запись двоичных данных.
    * multiwr - Запись данных в несколько файлов одновременно.