# Учим GO
### 1 ДЗ. Точное время
Написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.

### 2 ДЗ. Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* "qwe\4\5" => "qwe45" (*)
* "qwe\45" => "qwe44444" (*)
* "qwe\\5" => "qwe\\\\\" (*)
### 3 ДЗ. Частотный анализ
Написать функцию, которая получает на вход текст и возвращает
10 самых часто встречающихся слов без учета словоформ
