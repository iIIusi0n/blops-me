for file in $(find .next -type f -name '*.js'); do
  sed -i "s|BAKED_EXTERNAL_API_URL|${EXTERNAL_API_URL}|g" "$file"
done

npm start
