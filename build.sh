#!/bin/bash

# Ustawienia
OUTPUT_DIR="build"
OUTPUT_BIN="$OUTPUT_DIR/myapp"
SOURCE_DIR="./cmd/sysinfo"

# Sprawdź czy katalog build istnieje
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "Tworzenie katalogu $OUTPUT_DIR..."
    mkdir -p "$OUTPUT_DIR"
fi

# Budowanie aplikacji
echo "Budowanie aplikacji Go..."
go build -o "$OUTPUT_BIN" "$SOURCE_DIR"

# Sprawdzenie, czy build się powiódł
if [ $? -eq 0 ]; then
    echo "✅ Aplikacja została zbudowana jako $OUTPUT_BIN"
else
    echo "❌ Błąd podczas budowania aplikacji"
    exit 1
fi
#TODO: translate it to enshlish