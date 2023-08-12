# teletekst

telegazeta to nazwa handlowa TVP - do 2013

<!-- # Telewizja analogowa ? -->

# World System Teletext

1976

1989 w polsce, TVP

Bez kolorów:
red - powrót do pocz. magazynu
green - powrót do pocz. magazynu
yellow - poprz. strona
blue - nast strona

# Strony

czytaj co zostało z TODO

- 100 - **Start**
- 200 - **Prezentacja**
  - 201 - Poziom 1 1976 - WST (Ceefax 1974)
  - 202 - numeracja stron, magazyny (do 898)
  - 203 - podstrony (Natalia odpisz, zmienilem sie)
  - 204 - formatowanie tekstu
  - 205 - kolory
  - 206 - GSPS tekst
- 300 - **Dwa i pół** - poziom 2.5
- 400 - **Różności**
  - 401 - Nacu
- 500 - **Harmonogram**
- 600 - **Stunt GP**
- 700 - **Strony testowe**
  - 700 - tekst, miganie, kolory, G1, ukryty
  - 701 - boxing
  - 702 - 1.5 - polskie, G2, G3
  - 703 - 2.5 - palette, redefine CLUT 2, 3
- 747 - **Strony Testowe**
- 800 - labirynt / **Quiz**
  - 898 - Fastnav
  - 899 - Ostatnia strona
  - 89A - strony w hexach

201 - Poziom 1
202 - numeracja
898 - fastnav
899 - ostatnia
89A - 00 - FF
203 - podstrony
204 - Tekst, double height, migajacy, ukryty
205 - kolory
206 - LOGO GSPS tekst
mozaika
logo GSPS G1, "zebraliśmy już xxx zl" - L1
boxing
Zebralismy juz xxx zl - boxing

Info o nieskończonej pętli stron, zmiana na LCD

- Poziom 1.5
  logo GSPS G1, "zebraliśmy już xxx zł" - L1 - polskie znaki
  G2, G3
  logo GSPS G3, "zebraliśmy już xxx zł" - L1.5

- Poziom 2.5
  kolory
  podwójny tekst, szeroki tekst, migający
  migajacy na osobnej stronie, photosensitivity
  logo GSPS G3 + kolory, "zebraliśmy już xxx zł" - L2 - domyślne kolory
   <!-- Ale szybko poziom 2 został zatąpiony poziomem 2.5, który pozwalał na podmianę 16 kolorów na własne, z puli 4096 kolorów-->
  logo GSPS G3 + kolory, "zebraliśmy już xxx zł" - L2.5

Nie pokazuję 3.5, bo nie mam sprzętu

## dodaj kod QR

# Pominięte

polskie znaki - tylko wspomnij, bez info o diakrytykaahc i X/26

- 700
  - 1.0 - szybkie update'y (animacja?)
  - 2.5 - objects\*, kolory
    - Opcjonalne
      - custom znaczki (DRCS)

# TODO

jakiś prosty quiz?

| Element  |  kolor  | Kolor teletekstu |
| :------- | :-----: | :--------------: |
| zewnątrz | #0dcbff |       #0cf       |
| wewnątrz | #5a72ff |       #57f       |
| tekst    | #ffffff |       #fff       |
| kółko    | #224277 |       #247       |
| www      | #41228e |       #428       |

- bids BIDS_URL = `/search?type=allbids&event=22`;
- bids CURRENT_BIDS_URL = `/search?type=allbids&event=22&state=OPENED`;
- donations `/search?event=22&type=donation&feed=toread`,
- donations `/search?event=22&type=donation&commentstate=PENDING&transactionstate=COMPLETED`,
- donations `/search?event=22&type=bidtarget&state=PENDING`,
- donations `/search?event=22&type=donationbid`,
- prizes `/search?event=22&type=prize`,

dzieciom 2023 - event 21
