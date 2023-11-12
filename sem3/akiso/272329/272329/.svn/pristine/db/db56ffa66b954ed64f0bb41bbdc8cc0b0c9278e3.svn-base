for file in *; do
    if [ -e "$file" ]; then
        new_name=$(echo "$file" | tr '[:upper:]' '[:lower:]')
        if [ "$file" != "$new_name" ]; then
            mv -- "$file" "$new_name"
            echo "Zmieniono nazwę pliku: $file -> $new_name"
        fi
    fi
done
