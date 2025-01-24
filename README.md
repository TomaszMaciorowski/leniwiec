# leniwiec
Ten program służy do organizowania i przenoszenia plików graficznych z jednego katalogu do drugiego. Obsługuje pliki z wybranymi rozszerzeniami (np. .jpg, .png) i działa niezależnie od wielkości liter w rozszerzeniach. Dodatkowo, pliki są automatycznie nazywane na podstawie daty utworzenia (jeśli jest dostępna w metadanych EXIF) lub bieżącego czasu, a następnie zapisywane w katalogu docelowym.

Funkcjonalność programu:
Skanowanie katalogu źródłowego: Program przechodzi przez katalog źródłowy, wyszukując pliki z podanymi rozszerzeniami (np. .jpg, .png).
Odczyt daty z metadanych EXIF: Jeżeli plik graficzny zawiera dane EXIF (np. data zrobienia zdjęcia), program używa tej daty do generowania nazwy pliku.
Tworzenie nazw plików: Każdy plik otrzymuje nazwę w formacie YYYYMMDDHHMMSS_INDEX.jpg, gdzie INDEX to numer porządkowy pliku.
Przenoszenie plików: Pliki są przenoszone z katalogu źródłowego do docelowego z nowymi nazwami.
Wymagania:
Program wymaga podania trzech argumentów:
Ścieżka do katalogu źródłowego.
Ścieżka do katalogu docelowego.
Lista rozszerzeń plików oddzielonych przecinkami (np. jpg,png).

leniwiec.exe /ścieżka/do/źródła /ścieżka/do/docelowego jpg,png
